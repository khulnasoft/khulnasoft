import React from 'react';
import { StrategicRecommendationReport } from '@/app/lib/actions/customer_insights/types';

export default function StrategicRecommendations({ data }: { data: StrategicRecommendationReport }) {
  return (
    <div>
      <h2 className="mb-4 text-2xl font-bold">Strategic Recommendations</h2>
      {data.trends.length === 0 ? (
        <p>No trends found in the dataset.</p>
      ) : (
        <div>
          <div>
            <h3 className="mb-2 text-xl font-semibold">Recommendations:</h3>
            <p>{data.recommendation}</p>
          </div>
          <table className="mb-6 w-full table-auto border-collapse border border-gray-300">
            <thead>
              <tr>
                <th className="border border-gray-300 px-4 py-2">Trend Name</th>
                <th className="border border-gray-300 px-4 py-2">Description</th>
                <th className="border border-gray-300 px-4 py-2">Frequency</th>
                <th className="border border-gray-300 px-4 py-2">Customer Impact</th>
              </tr>
            </thead>
            <tbody>
              {data.trends.map((trend, index) => (
                <tr key={index}>
                  <td className="border border-gray-300 px-4 py-2">{trend.name}</td>
                  <td className="border border-gray-300 px-4 py-2">{trend.description}</td>
                  <td className="border border-gray-300 px-4 py-2">{trend.metrics.frequency}</td>
                  <td className="border border-gray-300 px-4 py-2">{trend.metrics.customer_impact}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
