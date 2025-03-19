'use client';

import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import { useState, useRef, useEffect } from 'react';
import { ComposerChatMessage } from '@/app/lib/actions/composer/chat_message';
import Markdown from '../ui/markdown';
import { Loader2 } from 'lucide-react';

interface ChatProps {
  onChatMessage: (content: string) => void;
  messages: ComposerChatMessage[];
}

export default function Chat({ onChatMessage, messages }: ChatProps) {
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!input) return;
    setInput('');
    setLoading(true);
    try {
      await onChatMessage(input);
    } finally {
      setLoading(false);
    }
  }

  // Capture Ctrl+Enter or Cmd+Enter
  function handleKeyDown(e: React.KeyboardEvent<HTMLTextAreaElement>) {
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
      handleSubmit(e as unknown as React.FormEvent);
    }
  }

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const isAssistantTyping = loading && messages.length > 0 && messages[messages.length - 1].sender === 'user';

  return (
    <div className="flex h-full flex-col">
      <ScrollArea className="flex-1 p-4">
        {messages.map((message, i) => {
          const isUser = message.sender === 'user';
          const bubbleStyles = isUser
            ? 'ml-auto bg-blue-200 dark:bg-blue-500 text-black dark:text-white'
            : 'bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-gray-50';

          const content = message.content.trim();
          if (!content) return null;

          return (
            <div key={i} className={`mb-4 max-w-[80%] rounded-lg p-3 ${bubbleStyles}`}>
              <Markdown contents={message.content} />
            </div>
          );
        })}
        {isAssistantTyping && (
          <div className="bg-accent text-accent-foreground mb-4 flex max-w-[80%] items-center space-x-1 rounded-lg p-3">
            <span className="inline-block h-2 w-2 animate-bounce rounded-full bg-current" />
            <span
              className="inline-block h-2 w-2 animate-bounce rounded-full bg-current"
              style={{ animationDelay: '200ms' }}
            />
            <span
              className="inline-block h-2 w-2 animate-bounce rounded-full bg-current"
              style={{ animationDelay: '400ms' }}
            />
          </div>
        )}
        <div ref={messagesEndRef} />
      </ScrollArea>

      <div className="border-border border-t p-4">
        <form onSubmit={handleSubmit} className="flex flex-col space-y-2">
          <div className="flex items-stretch space-x-2">
            <textarea
              value={input}
              onChange={(e) => setInput(e.target.value)}
              placeholder="Type your message..."
              className="border-border h-24 flex-grow resize-none rounded-md border p-2 focus:border-blue-500 focus:outline-none dark:bg-gray-800 dark:text-gray-50"
              onKeyDown={handleKeyDown}
              disabled={loading}
            />
            <Button type="submit" disabled={loading} className="flex h-24 items-center justify-center">
              {loading ? <Loader2 className="animate-spin" /> : 'Send'}
            </Button>
          </div>
          {/* Small grey note */}
          <div className="text-xs text-gray-500">Press Ctrl or ⌘ + Enter to send</div>
        </form>
      </div>
    </div>
  );
}
