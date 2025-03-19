import { ExternalLink } from 'lucide-react';
import { Badge } from './ui/badge';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from './ui/accordion';
import { Button } from './ui/button';
import { useState } from 'react';
import { Job } from '@/app/lib/actions/common/entities/pipelines';

export const FailedJobsComponent = ({ failedJobs }: { failedJobs: Job[] }) => {
  const [expandedJob, setExpandedJob] = useState<number | null>(null);

  return (
    <div className="space-y-4">
      <h3 className="mb-2 text-lg font-semibold">Failed Jobs</h3>
      {failedJobs.map((job) => (
        <div key={job.id} className="space-y-2 rounded-lg border p-4">
          <div className="flex items-center justify-between">
            <h4 className="text-md font-medium">{job.name}</h4>
            <a
              href={job.web_url}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center gap-1 text-sm text-blue-500 hover:underline"
            >
              View Job <ExternalLink className="h-4 w-4" />
            </a>
          </div>
          <Badge variant="destructive">Failed</Badge>
          <Accordion type="single" collapsible className="w-full">
            <AccordionItem value={`analysis-job-${job.id}`}>
              <AccordionTrigger>Analysis</AccordionTrigger>
              <AccordionContent>
                <pre className="bg-muted max-h-40 overflow-y-auto rounded-md p-2 text-xs">{job.reason}</pre>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
          <Accordion type="single" collapsible className="w-full">
            <AccordionItem value={`job-${job.id}`}>
              <AccordionTrigger>Job Log</AccordionTrigger>
              <AccordionContent>
                <pre className="bg-muted max-h-40 overflow-y-auto rounded-md p-2 text-xs">{job.log}</pre>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
          <Button variant="outline" size="sm" onClick={() => setExpandedJob(expandedJob === job.id ? null : job.id)}>
            {expandedJob === job.id ? 'Hide Details' : 'Show Details'}
          </Button>
          {expandedJob === job.id && (
            <div className="mt-2 space-y-2">
              <p className="text-sm">
                <strong>Stage:</strong> {job.stage}
              </p>
              <p className="text-sm">
                <strong>Duration:</strong> {job.duration} seconds
              </p>
              <p className="text-sm">
                <strong>Started At:</strong> {new Date(job.started_at).toLocaleString()}
              </p>
              <p className="text-sm">
                <strong>Finished At:</strong> {new Date(job.finished_at).toLocaleString()}
              </p>
            </div>
          )}
        </div>
      ))}
    </div>
  );
};
