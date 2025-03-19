import { ArrowUpRight, FileCode } from 'lucide-react';
import Markdown from './ui/markdown';
import { Badge } from './ui/badge';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from './ui/accordion';
import { MergeRequest } from '@/app/lib/actions/common/entities/merge_request';

const getImpactColor = (impact: string) => {
  switch (impact.toLowerCase()) {
    case 'high':
      return 'destructive';
    case 'medium':
      return 'secondary';
    default:
      return 'default';
  }
};

const navigateToCodeChange = (codeChangeUrl: string) => {
  window.open(codeChangeUrl, '_blank');
};

export const CodeChangesSection = ({ mrData }: { mrData: MergeRequest }) => (
  <div className="space-y-4">
    <Accordion type="multiple" className="w-full space-y-4">
      {mrData.codeChanges?.map((change, index) => (
        <AccordionItem value={`change-${index}`} key={index} className="rounded-lg border">
          <AccordionTrigger className="px-4">
            <div className="flex w-full items-center gap-3">
              <FileCode className="text-primary h-4 w-4" />
              <span className="text-sm font-medium">{change.new_path}</span>
              <Badge variant={getImpactColor(change.impact)} className="ml-auto">
                Impact: {change.impact}
              </Badge>
              <button
                className="hover:bg-muted ml-2 rounded-md p-1 transition-colors"
                onClick={() => navigateToCodeChange(change.web_url)}
              >
                <ArrowUpRight className="text-muted-foreground h-4 w-4" />
              </button>
            </div>
          </AccordionTrigger>
          <AccordionContent className="px-4 pt-2">
            <div className="space-y-4">
              <div className="bg-muted/50 rounded-lg p-4">
                <h4 className="mb-2 text-sm font-semibold">Changes Explained</h4>
                <Markdown contents={change.summary}></Markdown>
              </div>
              <div className="bg-muted/50 rounded-lg p-4">
                <h4 className="mb-2 text-sm font-semibold">Review Notes</h4>
                <Markdown contents={change.review}></Markdown>
              </div>
            </div>
          </AccordionContent>
        </AccordionItem>
      ))}
    </Accordion>
  </div>
);
