'use client';

import { useState, useRef, useEffect } from 'react';
import { Send, Bot, User, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { ChatMessage, sendChatMessage } from '@/app/lib/actions/chat';
import { getRunbookQuery } from '@/app/lib/actions/runbook_query';
import Markdown from './ui/markdown';

export default function FullScreenChat() {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [newMessage, setNewMessage] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const messagesEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({ behavior: 'smooth' });
    }
  }, [messages]);

  const sendMessage = async (e: React.FormEvent) => {
    setSubmitting(true);
    e.preventDefault();
    if (!newMessage.trim()) return;
    setNewMessage('');
    let initialMessages: ChatMessage[] = [];
    if (messages.length === 0) {
      const runbookquery = await getRunbookQuery(newMessage).catch((error) => {
        console.error('Error in getting Runbooks Query:', error);
        throw error;
      });
      let contexts: string = '';
      for (const context of runbookquery.contexts) {
        contexts += `<context><conten>${context.content}</content><sourceLink>${context.sourceLink.replace('/raw/', '/tree/')}</sourceLink></context>`;
      }
      initialMessages = [
        {
          sender: 'System',
          content: `
            You are an assistant tasked with answering the user's question based on the information provided and your broader knowledge. Consider the following list contexts as background information:
            <contexts>${contexts}</contexts>

            Now, answer the user's question accurately, using both the given context and any additional knowledge you possess:
            <question>${runbookquery.query}</question>

            Focus on being concise, relevant, and clear while prioritizing context when applicable.
            If your answer is based on the given contexts, include a Sources section in bold at the bottom of the response. Use markdown format for the source links.
            `,
          timestamp: new Date().toISOString(),
        },
      ];
    }

    const userMessage = {
      sender: 'Human',
      content: newMessage,
      timestamp: new Date().toISOString(),
    };

    if (newMessage === '/reset' || newMessage === '/clear') {
      setMessages([]);
      setNewMessage('');
      setSubmitting(false);
      return;
    }

    const newMessages = [...messages, ...initialMessages, userMessage];
    setMessages(newMessages);

    const stream = await sendChatMessage(newMessages, null, null, null, true);

    let aiMessage = {
      sender: 'AI',
      content: '',
      timestamp: new Date().toISOString(),
    };
    setSubmitting(false);

    for await (const chunk of stream) {
      aiMessage = {
        ...aiMessage,
        content: aiMessage.content + chunk,
        timestamp: new Date().toISOString(),
      };
      setMessages([...newMessages, aiMessage]);
    }
  };

  return (
    <div className="bg-background flex h-full flex-col">
      {submitting && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
          <div className="flex items-center gap-2 text-lg font-semibold text-white">
            <Loader2 className="h-6 w-6 animate-spin" />
            <span>Gathering Context... The First query takes some time.</span>
          </div>
        </div>
      )}
      <div className="flex-grow overflow-y-auto">
        <ScrollArea className="h-full p-4">
          <div className="space-y-4">
            {messages
              .filter((message) => message.sender !== 'System') // Skip messages from 'System'
              .map((message, index) => (
                <div key={index} className="flex gap-3">
                  <Avatar className={`h-8 w-8 ${message.sender === 'AI' ? 'bg-blue-100' : 'bg-green-100'}`}>
                    <AvatarFallback>
                      {message.sender === 'AI' ? (
                        <Bot className="h-4 w-4 text-blue-500" />
                      ) : (
                        <User className="h-4 w-4 text-green-500" />
                      )}
                    </AvatarFallback>
                  </Avatar>
                  <div className="grid max-w-full gap-1">
                    <div className="flex items-center gap-2">
                      <span className="text-sm font-medium">{message.sender}</span>
                      <span className="text-muted-foreground text-xs">
                        {new Date(message.timestamp).toLocaleTimeString()}
                      </span>
                    </div>
                    <p className="break-words text-sm">
                      <Markdown contents={message.content} />
                    </p>
                  </div>
                </div>
              ))}
            <div ref={messagesEndRef} />
          </div>
        </ScrollArea>
      </div>
      <form onSubmit={sendMessage} className="mt-24 border-t p-3">
        <div className="flex gap-2">
          <Input
            placeholder="Type a message..."
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            className="flex-1 px-4 py-2"
            disabled={submitting}
          />
          <Button type="submit" size="icon">
            <Send className="h-4 w-4" />
            <span className="sr-only">Send</span>
          </Button>
        </div>
      </form>
    </div>
  );
}
