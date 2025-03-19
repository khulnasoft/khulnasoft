import React, { useState } from 'react';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { BookOpen, Link, MessageSquare, ChevronDown, ChevronUp, LockKeyhole } from 'lucide-react';
import type { Issue } from '@/app/lib/actions/common/entities/issue';
import Markdown from '@/components/ui/markdown';

const IssueAnalysis = ({ issue }: { issue: Issue }) => {
  const analysis = issue.analysis;
  const [expandedIssues, setExpandedIssues] = useState<{ [key: number]: boolean }>({});
  const [expandedMRs, setExpandedMRs] = useState<{ [key: number]: boolean }>({});

  const toggleIssueExpansion = (index: number) => {
    setExpandedIssues((prev) => ({
      ...prev,
      [index]: !prev[index],
    }));
  };

  const toggleMRExpansion = (index: number) => {
    setExpandedMRs((prev) => ({
      ...prev,
      [index]: !prev[index],
    }));
  };

  return (
    <Tabs defaultValue="understanding" className="w-full">
      <TabsList className="grid w-full grid-cols-4">
        <TabsTrigger value="understanding">
          <BookOpen className="mr-2 h-4 w-4" />
          Understanding
        </TabsTrigger>
        <TabsTrigger value="related">
          <Link className="mr-2 h-4 w-4" />
          Related Items
        </TabsTrigger>
        <TabsTrigger value="comments">
          <MessageSquare className="mr-2 h-4 w-4" />
          Comments
        </TabsTrigger>
        <TabsTrigger value="security">
          <LockKeyhole className="mr-2 h-4 w-4" />
          Security
        </TabsTrigger>
      </TabsList>

      <TabsContent value="understanding">
        <ScrollArea className="h-[calc(100vh-400px)] w-full rounded-md border p-4">
          <div className="space-y-6">
            <section>
              <h3 className="mb-2 text-lg font-semibold">Main Problem & Outcome</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.understanding.mainProblem}
              </p>
            </section>

            <section>
              <h3 className="mb-2 text-lg font-semibold">Requirements & Details</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.understanding.requirements}
              </p>
            </section>

            <section>
              <h3 className="mb-2 text-lg font-semibold">Use Case Analysis</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.understanding.useCase}
              </p>
            </section>

            <section>
              <h3 className="mb-2 text-lg font-semibold">Unfamiliar Terms</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.understanding.unfamiliarTerms}
              </p>
            </section>

            <section>
              <h3 className="mb-2 text-lg font-semibold">Key Terms & Concepts</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.understanding.keyTerms}
              </p>
            </section>
          </div>
        </ScrollArea>
      </TabsContent>

      <TabsContent value="related">
        <ScrollArea className="h-[calc(100vh-400px)] w-full rounded-md border p-4">
          <div className="space-y-6">
            <section>
              <h3 className="mb-4 text-lg font-semibold">Related Issues</h3>
              {issue.linkedIssues.map((issue, index) => (
                <div key={index} className="mt-2 space-y-2 rounded-md border p-4">
                  <div className="flex items-center justify-between">
                    <div className="font-medium">
                      <span className="font-semibold">Title: </span>
                      {issue.title}
                    </div>
                    <button
                      className="text-muted-foreground hover:text-foreground transition-colors"
                      onClick={() => toggleIssueExpansion(index)}
                    >
                      {expandedIssues[index] ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
                    </button>
                  </div>
                  {expandedIssues[index] && <Markdown contents={issue.summary} />}
                </div>
              ))}
            </section>

            <section>
              <h3 className="mb-4 text-lg font-semibold">Related Merge Requests</h3>
              {issue.mergeRequests.map((mr, index) => (
                <div key={index} className="mt-2 space-y-2 rounded-md border p-4">
                  <div className="flex items-center justify-between">
                    <div className="font-medium">
                      <span className="font-semibold">Title: </span>
                      {mr.title}
                    </div>
                    <button
                      className="text-muted-foreground hover:text-foreground transition-colors"
                      onClick={() => toggleMRExpansion(index)}
                    >
                      {expandedMRs[index] ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
                    </button>
                  </div>
                  {expandedMRs[index] && <Markdown contents={mr.summary} />}
                </div>
              ))}
            </section>
          </div>
        </ScrollArea>
      </TabsContent>

      <TabsContent value="comments">
        <ScrollArea className="h-[calc(100vh-400px)] w-full rounded-md border p-4">
          <div className="space-y-6">
            <section>
              <h3 className="mb-2 text-lg font-semibold">Key Insights</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.comments.insights}
              </p>
            </section>

            <section>
              <h3 className="mb-2 text-lg font-semibold">Concerns & Suggestions</h3>
              <p className="text-muted-foreground text-sm" style={{ whiteSpace: 'pre-line' }}>
                {analysis.comments.concerns}
              </p>
            </section>
          </div>
        </ScrollArea>
      </TabsContent>

      <TabsContent value="security">
        <ScrollArea className="h-[calc(100vh-400px)] w-full rounded-md border p-4">
          <Markdown contents={issue.securityRecommentations} />
        </ScrollArea>
      </TabsContent>
    </Tabs>
  );
};

export default IssueAnalysis;
