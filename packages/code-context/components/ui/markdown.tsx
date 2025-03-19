import React from 'react';
import ReactMarkdown from 'react-markdown';
import rehypeHighlight from 'rehype-highlight';
import remarkGfm from 'remark-gfm';
import remarkGemoji from 'remark-gemoji';
import 'highlight.js/styles/github-dark.css';

interface MarkdownProps {
  contents: string;
}

interface CodeProps extends React.HTMLAttributes<HTMLElement> {
  inline?: boolean;
  className?: string;
  children?: React.ReactNode;
}

const Markdown: React.FC<MarkdownProps> = ({ contents }) => {
  return (
    <ReactMarkdown
      remarkPlugins={[remarkGfm, remarkGemoji]}
      rehypePlugins={[rehypeHighlight]}
      components={{
        h1: ({ ...props }) => <h1 className="text-foreground mb-4 mt-8 text-3xl font-bold" {...props} />,
        h2: ({ ...props }) => <h2 className="text-foreground mb-3 mt-6 text-2xl font-semibold" {...props} />,
        h3: ({ ...props }) => <h3 className="text-foreground mb-2 mt-5 text-xl font-medium" {...props} />,
        p: ({ ...props }) => <p className="text-foreground mb-4 leading-relaxed" {...props} />,
        ul: ({ ...props }) => <ul className="text-foreground mb-4 list-disc space-y-2 pl-6" {...props} />,
        ol: ({ ...props }) => <ol className="text-foreground mb-4 list-decimal space-y-2 pl-6" {...props} />,
        li: ({ ...props }) => <li className="text-foreground mb-1" {...props} />,
        blockquote: ({ ...props }) => (
          <blockquote className="border-primary text-foreground mb-4 border-l-4 py-2 pl-4 italic" {...props} />
        ),
        a: ({ ...props }) => <a className="text-primary hover:underline" {...props} />,
        code: ({ inline, className, children, ...props }: CodeProps) => {
          const match = /language-(\w+)/.exec(className || '');
          if (!inline && match) {
            return (
              <pre className="bg-muted my-4 overflow-x-auto rounded-md p-4">
                <code className={className} {...props}>
                  {children}
                </code>
              </pre>
            );
          } else {
            return (
              <code className="bg-muted text-foreground rounded-md px-1 py-0.5 text-sm" {...props}>
                {children}
              </code>
            );
          }
        },
        table: ({ ...props }) => (
          <div className="mb-4 overflow-x-auto">
            <table className="divide-border min-w-full divide-y" {...props} />
          </div>
        ),
        th: ({ ...props }) => <th className="bg-muted text-foreground px-3 py-2 text-left font-semibold" {...props} />,
        td: ({ ...props }) => <td className="border-border text-foreground border-t px-3 py-2" {...props} />,
      }}
    >
      {contents}
    </ReactMarkdown>
  );
};

export default Markdown;
