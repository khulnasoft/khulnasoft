'use client';

import { useState } from 'react';
import { Sidebar } from '@/components/my_achievements/sidebar';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { MergeRequestTable } from '@/components/my_achievements/merge_request_list';
import { ProjectList } from '@/components/my_achievements/project_list';
import { SummaryDisplay } from '@/components/my_achievements/summary_display';
import Navbar from '../navbar';
import { getAllMRsAndProjects } from '@/app/lib/actions/common/get_all_mrs';
import { MergeRequest } from '@/app/lib/actions/common/entities/merge_request';
import { Project } from '@/app/lib/actions/common/entities/project';
import { useSession } from 'next-auth/react';
import Login from '../login';
import { summarizeAchievements } from '@/app/lib/actions/achievements/actions';
import { Loader2, CheckCircle2 } from 'lucide-react';

export default function GitLabAchievementSummary() {
  const [mergeRequests, setMergeRequests] = useState<MergeRequest[]>([]);
  const [projects, setProjects] = useState<Project[]>([]);
  const [summary, setSummary] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [loadingStatus, setLoadingStatus] = useState<string>('');
  const [isDataGenerated, setIsDataGenerated] = useState<boolean>(false);

  const { data: session } = useSession();

  if (!session) {
    return <Login />;
  }

  const handleGenerate = async (
    dateRange: { from: Date; to: Date } | undefined,
    temperature: number,
    fetchDiffs: boolean
  ) => {
    try {
      setIsLoading(true);
      setLoadingStatus('');
      setLoadingStatus('Fetching Merge Requests & Projects...');
      const { mergeRequests, projects } = await getAllMRsAndProjects(dateRange, fetchDiffs);

      setLoadingStatus('Merge Requests & Projects fetched');
      setMergeRequests(mergeRequests);
      setProjects(projects);

      setLoadingStatus('Generating AI Summary...');
      const summary = await summarizeAchievements(mergeRequests, temperature);
      setSummary(summary);

      setLoadingStatus('Summary Generated');
      setIsDataGenerated(true);
    } catch (error) {
      console.error('Error generating summary:', error);
      setLoadingStatus('Error generating summary');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <Navbar showSettings={false} />
      <div className="flex h-screen">
        <Sidebar onGenerate={handleGenerate} />
        <main className="flex-1 overflow-auto p-6">
          <h1 className="mb-6 text-3xl font-bold">My Achievements</h1>

          <div className="mb-6 rounded-lg border border-blue-200 bg-blue-50 p-4 dark:border-blue-800 dark:bg-blue-900/30">
            <h3 className="mb-2 text-lg font-semibold text-blue-800 dark:text-blue-200">AI-Powered Achievement Log</h3>
            <p className="mb-2 text-blue-700 dark:text-blue-100">
              This log, powered by AI, does not encompass all your MRs, given its 1500-word limitation. It is a concise
              selection and does not cover confidential issues or private repositories.
            </p>
            <p className="text-blue-700 dark:text-blue-100">
              For those seeking to use this as a data point for talent assessments, it is imperative to click on the
              links associated with the summaries rigorously to verify their accuracy.
            </p>
          </div>

          {isLoading ? (
            <div className="flex h-64 flex-col items-center justify-center space-y-4">
              <Loader2 className="text-primary h-12 w-12 animate-spin" />
              <div className="flex items-center space-x-2">
                {loadingStatus && (
                  <>
                    {loadingStatus.includes('Error') ? (
                      <span className="text-red-500">{loadingStatus}</span>
                    ) : (
                      <>
                        <CheckCircle2 className="h-5 w-5 text-green-500" />
                        <span className="text-muted-foreground">{loadingStatus}</span>
                      </>
                    )}
                  </>
                )}
              </div>
            </div>
          ) : !isDataGenerated ? (
            <div className="text-muted-foreground flex h-64 items-center justify-center text-xl">
              No data has been selected yet. Use the sidebar to generate your achievement summary.
            </div>
          ) : (
            <Tabs defaultValue="merge-requests">
              <TabsList>
                <TabsTrigger value="merge-requests">Merge Requests</TabsTrigger>
                <TabsTrigger value="projects">Projects</TabsTrigger>
                <TabsTrigger value="summary">Summary</TabsTrigger>
              </TabsList>
              <TabsContent value="merge-requests">
                <MergeRequestTable mergeRequests={mergeRequests} />
              </TabsContent>
              <TabsContent value="projects">
                <ProjectList projects={projects} />
              </TabsContent>
              <TabsContent value="summary">
                <SummaryDisplay summary={summary} />
              </TabsContent>
            </Tabs>
          )}
        </main>
      </div>
    </div>
  );
}
