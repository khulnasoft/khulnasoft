import React from 'react';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ClipboardList, GitBranch, TestTube, TestTube2, University, HelpCircle, LucideIcon } from 'lucide-react';
import Markdown from '@/components/ui/markdown';
import { MergeRequest } from '@/app/lib/actions/common/entities/merge_request';

interface AnalysisSectionProps {
  icon: LucideIcon;
  title: string;
  content: string;
}

const AnalysisSection = ({ icon: Icon, title, content }: AnalysisSectionProps) => (
  <Card className="mb-4">
    <CardHeader className="pb-3">
      <div className="flex items-center gap-2">
        <Icon className="text-primary h-5 w-5" />
        <CardTitle className="text-lg">{title}</CardTitle>
      </div>
    </CardHeader>
    <CardContent>
      <Markdown contents={content} />
    </CardContent>
  </Card>
);

const AnalysisTab = ({ mrData }: { mrData: MergeRequest }) => {
  if (!mrData?.analysis) return <div className="text-muted-foreground p-4 text-center">Analysis not available</div>;

  return (
    <ScrollArea className="h-[600px] w-full rounded-md border p-4">
      <div className="space-y-4">
        <AnalysisSection icon={ClipboardList} title="Review Approach" content={mrData.analysis.reviewApproach} />

        {mrData.analysis.breakdown && (
          <AnalysisSection icon={GitBranch} title="Proposed Breakdown" content={mrData.analysis.breakdown} />
        )}

        <AnalysisSection
          icon={University}
          title="Architectural Components"
          content={mrData.analysis.architecturalComponents || ''}
        />

        <AnalysisSection icon={TestTube2} title="Testing Strategy" content={mrData.analysis.testingStrategy} />

        <AnalysisSection
          icon={TestTube}
          title="Testing Documentation"
          content={mrData.analysis.testingDocumentation || ''}
        />

        <AnalysisSection icon={HelpCircle} title="Suggested Questions" content={mrData.analysis.suggestedQuestions} />
      </div>
    </ScrollArea>
  );
};

export default AnalysisTab;
