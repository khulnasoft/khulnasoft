'use client';

import Navbar from '../navbar';
import { useSession } from 'next-auth/react';
import Login from '../login';
import FullScreenChat from '../chat_fullscreen';

export default function GitLabRunbooksDocs() {
  const { data: session } = useSession();
  if (!session) {
    return <Login />;
  }
  return (
    <div className="flex h-screen flex-col">
      <Navbar showSettings={false} />
      <div className="relative flex-grow">
        <FullScreenChat />
      </div>
    </div>
  );
}
