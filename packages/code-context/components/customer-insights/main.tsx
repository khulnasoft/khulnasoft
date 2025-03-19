'use client';

import { useCallback, useState } from 'react';
import QueryInterface from './query';
import ReportTabs from './report';
import { CustomerInsightsQuery, CustomerInsightsReport } from '@/app/lib/actions/customer_insights/types';
import { fetchCustomerInsights } from '@/app/lib/actions/customer_insights/actions';

export default function Main() {
  const [report, setReport] = useState<CustomerInsightsReport | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchReport = useCallback(async (query: CustomerInsightsQuery) => {
    setIsLoading(true);
    setError(null);

    try {
      const insights = await fetchCustomerInsights(query);

      console.log(insights);
      setReport(insights);
    } catch (err) {
      console.error('Error fetching customer insights:', err);
      setError('Failed to load customer insights. Please try again later.');
    } finally {
      setIsLoading(false);
    }
  }, []);

  return (
    <div className="flex flex-1 overflow-auto">
      <div className="border-border flex w-1/4 flex-col border-r">
        <QueryInterface onSubmit={fetchReport} />
      </div>
      <div className="flex w-3/4 flex-col overflow-hidden p-4">
        <div className="mb-4 flex items-center justify-between">
          <h1 className="text-2xl font-bold">Customer Insights Report</h1>
        </div>
        <div className="flex-1 overflow-auto rounded-md border">
          {isLoading ? (
            <div>Loading...</div>
          ) : error ? (
            <div>{error}</div>
          ) : report ? (
            <ReportTabs results={report} />
          ) : null}
        </div>
      </div>
    </div>

    // <div className="p-6">
    //   <h1 className="text-3xl font-bold mb-4">Customer Insights</h1>
    //   <QueryInterface onSubmit={fetchReport} />
    //   {
    //     isLoading ? <div>Loading...</div> : error ? <div>{error}</div> : report ? <ReportTabs results={report} /> : null
    //   }
    // </div>
  );
}
