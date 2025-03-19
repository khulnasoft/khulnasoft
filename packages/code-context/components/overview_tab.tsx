import Markdown from './ui/markdown';
import { Badge } from './ui/badge';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from './ui/accordion';
import { ExternalLink, GitCommit } from 'lucide-react';
import { calculateDaysSince } from '@/app/lib/utils';
import { FailedJobsComponent } from './failed_jobs';
import { Commit, MergeRequest } from '@/app/lib/actions/common/entities/merge_request';

export const OverviewSection = ({ mrData }: { mrData: MergeRequest }) => (
  <div className="space-y-4">
    <div className="bg-muted/50 rounded-lg p-4">
      <h3 className="mb-2 text-lg font-semibold">Summary</h3>
      <div className="space-y-2">
        <div className="flex items-center gap-2">
          <Badge variant="outline">Changes: {mrData.codeChanges?.length || 0} files</Badge>
          <Badge variant="outline">Commits: {mrData.commits?.length || 0}</Badge>
          <Badge variant="outline">Comments: {mrData.discussions?.length || 0}</Badge>
          <Badge variant={calculateDaysSince(mrData.created_at) > 7 ? 'destructive' : 'outline'}>
            MR Age: {calculateDaysSince(mrData.created_at) || 0}
          </Badge>
        </div>
        {mrData.summary && (
          <div className="mt-3">
            <Markdown contents={mrData.summary.summary} />
          </div>
        )}
        <div className="mt-3">
          <h4 className="mb-4 mt-4 text-sm font-medium font-semibold">Key Code Changes</h4>
          <ul className="list-inside list-disc space-y-1 text-sm">
            <Markdown contents={mrData.summary.keyChanges} />
          </ul>
        </div>
      </div>
    </div>
    <Accordion type="single" collapsible className="w-full">
      <AccordionItem value="description">
        <AccordionTrigger className="text-lg font-semibold">Original MR Description</AccordionTrigger>
        <AccordionContent>
          <Markdown contents={mrData.description} />
        </AccordionContent>
      </AccordionItem>
    </Accordion>
    <div>
      <h3 className="mb-2 text-lg font-semibold">Status</h3>
      <div className="flex items-center gap-2">
        <Badge variant={mrData.state === 'open' ? 'default' : 'secondary'}>{mrData.state}</Badge>
        <Badge variant={mrData.pipeline?.status === 'success' ? 'default' : 'destructive'}>
          Pipeline: {mrData.pipeline?.status || 'unknown'}
        </Badge>
      </div>
    </div>
    {mrData.pipeline?.status === 'failed' && mrData.failingJobs && (
      <FailedJobsComponent failedJobs={mrData.failingJobs} />
    )}
    <div>
      <h3 className="mb-2 text-lg font-semibold">Commits ({mrData.commits?.length || 0})</h3>
      <div className="space-y-2">
        {mrData.commits?.map((commit: Commit, index) => (
          <a
            key={index}
            href={commit.web_url}
            target="_blank"
            rel="noopener noreferrer"
            className="hover:bg-muted group flex items-center gap-2 rounded-md p-2 transition-colors"
          >
            <GitCommit className="text-muted-foreground h-4 w-4" />
            <span className="flex-grow text-sm">{commit.message}</span>
            <ExternalLink className="h-4 w-4 opacity-0 transition-opacity group-hover:opacity-100" />
          </a>
        ))}
      </div>
    </div>
  </div>
);
