'use client';

import React, { useState, FormEvent, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { ReviewType } from './reviewer_prompts';
import { ArrowRight, User, Clock, Loader2 } from 'lucide-react';
import { getUserMRs } from '@/app/lib/actions/mr_analysis/user_mrs';
import { calculateDaysSince } from '@/app/lib/utils';
import { getCurrentUser } from '@/app/lib/actions/common/fetch_user';
import { MergeRequest } from '@/app/lib/actions/common/entities/merge_request';

const SearchLayout = ({
  onSubmit,
  hasResults,
  reviewType,
  setReviewType,
  defaultUrl,
}: {
  onSubmit: (url: string) => void;
  hasResults: boolean;
  reviewType: ReviewType;
  setReviewType: (type: ReviewType) => void;
  defaultUrl: string;
}) => {
  const [url, setURL] = useState('');
  const [assignedMRs, setAssignedMRs] = useState<MergeRequest[]>([]);
  const [loadingMRs, setLoadingMRs] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [username, setUsername] = useState<string | null>(null);
  const [loadingUser, setLoadingUser] = useState(true);

  useEffect(() => {
    const fetchUserInfo = async () => {
      setLoadingUser(true);
      try {
        // Replace this with your actual user info fetch call
        const response = await getCurrentUser();
        setUsername(response.username);
      } catch (error) {
        console.error('Failed to fetch user info:', error);
      } finally {
        setLoadingUser(false);
      }
    };

    const fetchAssignedMRs = async () => {
      setLoadingMRs(true);
      try {
        const mrs = await getUserMRs(3);
        setAssignedMRs(mrs);
      } catch (error) {
        console.error('Failed to fetch assigned MRs:', error);
      } finally {
        setLoadingMRs(false);
      }
    };

    fetchUserInfo();
    fetchAssignedMRs();
  }, []);

  useEffect(() => {
    if (defaultUrl) {
      setURL(defaultUrl);
    }
  }, [defaultUrl]);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setSubmitting(true);
    await onSubmit(url);
    setSubmitting(false);
  };

  const selectMR = async (mrUrl: string) => {
    setURL(mrUrl);
    setSubmitting(true);
    await onSubmit(mrUrl);
    setSubmitting(false);
  };

  const showMRs = !loadingMRs && assignedMRs.length > 0 && !hasResults && !submitting;

  if (loadingUser || loadingMRs) {
    return (
      <div className="flex min-h-screen flex-col items-center justify-center">
        <Loader2 className="text-primary mb-4 h-8 w-8 animate-spin" />
        <p className="text-muted-foreground">Loading...</p>
      </div>
    );
  }

  return (
    <div
      className={`w-full transition-all duration-300 ease-in-out ${
        hasResults ? 'pt-2' : '-mt-16 flex min-h-screen flex-col justify-center'
      }`}
    >
      <div
        className={`mx-auto max-w-6xl px-4 transition-all duration-300 ${hasResults ? '-translate-y-4 transform' : ''}`}
      >
        {!submitting && !hasResults && (
          <div className="mb-6 text-center">
            <h1 className="text-2xl font-semibold">Hello, @{username}</h1>
          </div>
        )}

        <form onSubmit={handleSubmit} className="mb-6 space-y-4">
          <div className="flex flex-col gap-4 sm:flex-row">
            <Input
              type="text"
              placeholder="Enter Merge Request, Issue, Epic or Blob URL"
              value={url}
              onChange={(e) => setURL(e.target.value)}
              className="min-w-[600px] flex-grow px-4"
              disabled={submitting}
            />
            {url.indexOf('merge_requests') > -1 && (
              <Select
                onValueChange={(value: string) => setReviewType(value as ReviewType)}
                value={reviewType}
                disabled={submitting}
              >
                <SelectTrigger className="w-full sm:w-[200px]">
                  <SelectValue placeholder="Select review type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="General">General</SelectItem>
                  <SelectItem value="Frontend">Frontend</SelectItem>
                  <SelectItem value="Backend">Backend</SelectItem>
                  <SelectItem value="Database">Database</SelectItem>
                  <SelectItem value="Tech Writing">Tech Writing</SelectItem>
                </SelectContent>
              </Select>
            )}
            <Button type="submit" disabled={submitting} className="w-full px-6 sm:w-auto">
              {submitting ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Analyzing...
                </>
              ) : (
                'Analyze'
              )}
            </Button>
          </div>
        </form>

        {showMRs && (
          <Card className="mt-8">
            <CardHeader>
              <CardTitle>Your Assigned Merge Requests</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                {assignedMRs.map((mr) => (
                  <Button
                    key={mr.id}
                    variant="ghost"
                    className="h-16 w-full justify-between text-left text-base hover:bg-slate-100"
                    onClick={() => selectMR(mr.web_url)}
                    disabled={submitting}
                  >
                    <div className="flex flex-col items-start">
                      <span className="truncate font-medium">{mr.title}</span>
                      <div className="text-muted-foreground mt-1 flex items-center text-sm">
                        <User className="mr-1 h-3 w-3" />
                        <span className="mr-3">{mr.author.name}</span>
                        <Clock className="mr-1 h-3 w-3" />
                        <span>
                          {calculateDaysSince(mr.created_at)} day{calculateDaysSince(mr.created_at) !== 1 ? 's' : ''}{' '}
                          ago
                        </span>
                      </div>
                    </div>
                    <ArrowRight className="ml-2 h-4 w-4 flex-shrink-0" />
                  </Button>
                ))}
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
};

export default SearchLayout;
