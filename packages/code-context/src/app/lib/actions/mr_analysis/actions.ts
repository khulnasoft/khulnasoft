'use server';

import { getServerSession } from 'next-auth';
import { authOptions } from '../../../api/auth/[...nextauth]/options';
import { trackAction, trackRun } from '@/app/lib/telemetry';
import { ReviewType } from '../../../../components/reviewer_prompts';
import { redirect } from 'next/navigation';
import { extractProjectInfo } from '@/app/lib/utils';
import { generateMRSummary } from './mr_summary';
import { generateMRAnalysis } from './mr_analysis';
import { getIssueSummaries } from '../common/issue_summary';
import { getChangeSummaries } from './change_summary';
import { formatDiscussions } from '../common/format_discussions';
import { getSecurityReview } from './security_review';
import { Discussion } from '../common/entities/discussions';
import { analyseDiscussions } from './analyse_discussions';
import { GITLAB_BASE_URL } from '../common/constants';
import { findReasonsForFailure, getFailingJobsForMR } from './pipeline';
import { MergeRequest } from '../common/entities/merge_request';

export async function fetchMRDetails(
  url: string,
  reviewType: ReviewType,
  customPromptCodeComments: string | null
): Promise<MergeRequest | null> {
  const mrURL = extractProjectInfo(url);
  const session = await getServerSession(authOptions);
  if (!session) {
    return null;
  }

  const { accessToken, user } = session;

  trackRun(user?.name, user?.email, url, 'merge_request').catch((e) => console.error('Could not track run:', e));

  if (customPromptCodeComments) {
    trackAction(session?.user?.name, 'custom_prompt_code_comments').catch((e) =>
      console.error('Could not track action:', e)
    );
  }

  try {
    const headers = {
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    };

    const mrDetailsURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}`;
    const mrCommitsURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}/commits`;
    const mrIssuesURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}/related_issues`;
    const mrChangesURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}/diffs`;
    const mrDiscussionsURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}/discussions`;

    const [
      mrDetailsResponse,
      mrCommitsResponse,
      mrIssuesResponse,
      mrChangesResponse,
      mrDiscussionsResponse,
      failedJobs,
    ] = await Promise.all([
      fetch(mrDetailsURL, { headers }),
      fetch(mrCommitsURL, { headers }),
      fetch(mrIssuesURL, { headers }),
      fetch(mrChangesURL, { headers }),
      fetch(mrDiscussionsURL, { headers }),
      getFailingJobsForMR(mrURL.projectPath, mrURL.mrIid, headers),
    ]);

    if (
      !mrDetailsResponse.ok ||
      !mrCommitsResponse.ok ||
      !mrIssuesResponse.ok ||
      !mrChangesResponse.ok ||
      !mrDiscussionsResponse.ok
    ) {
      throw new Error('GitLab API error in one or more requests');
    }

    const [mrDetails, mrCommits, mrIssues, mrChanges, mrDiscussions] = await Promise.all([
      mrDetailsResponse.json(),
      mrCommitsResponse.json(),
      mrIssuesResponse.json(),
      mrChangesResponse.json(),
      mrDiscussionsResponse.json(),
    ]);

    const relatedMRs = extractRelatedMRs(mrDiscussions);

    // Start parallel async calls
    const [relatedIssues, discussions, codeChanges, securityReview, failedJobsWithReason] = await Promise.all([
      getIssueSummaries(mrIssues, headers),
      formatDiscussions(mrDiscussions, mrDetails.web_url),
      getChangeSummaries(mrChanges, mrDetails, reviewType, customPromptCodeComments),
      getSecurityReview(mrChanges),
      findReasonsForFailure(failedJobs, mrDetails, mrChanges),
    ]);

    // Assign results to `mrDetails`
    mrDetails.commits = mrCommits;
    mrDetails.relatedIssues = relatedIssues;
    mrDetails.discussions = discussions;
    mrDetails.codeChanges = codeChanges;
    mrDetails.relatedMRs = relatedMRs;
    mrDetails.securityReview = securityReview;
    mrDetails.failingJobs = failedJobsWithReason;
    [mrDetails.summary, mrDetails.analysis, mrDetails.discussionsAnalysis] = await Promise.all([
      generateMRSummary(mrDetails, reviewType),
      generateMRAnalysis(mrDetails, reviewType),
      analyseDiscussions(mrDiscussions, mrDetails, user?.name || 'Unknown'),
    ]);

    return mrDetails;
  } catch (error) {
    console.error('GitLab API error:', error);
    if ((error as Error).message?.includes('GitLab API error')) {
      redirect('/api/auth/signout');
    }
    throw error;
  }
}

function extractRelatedMRs(discussions: Discussion[]) {
  const relatedMRs: { relatedMergeRequest: string; link: string }[] = [];
  discussions.forEach((discussion) => {
    discussion.notes?.forEach((note) => {
      const regex = /!([0-9]+)/g;
      let match;
      while ((match = regex.exec(note.body)) !== null) {
        relatedMRs.push({
          relatedMergeRequest: match[1],
          link: note.web_url,
        });
      }
    });
  });
  return relatedMRs;
}
