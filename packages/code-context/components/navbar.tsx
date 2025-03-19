'use client';

import {
  Code,
  LogOut,
  MessageSquare,
  Moon,
  Settings,
  Sun,
  Gitlab,
  BugPlay,
  Trophy,
  Home,
  Book,
  PencilLine,
  BookA,
  Video,
  Presentation,
} from 'lucide-react';
import { Button } from './ui/button';
import { signOut } from 'next-auth/react';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import FeatureRequestDialog from './feature_request_dialog';

interface NavbarProps {
  showSettings?: boolean;
  onSettingsClick?: () => void;
  onThemeUpdate?: (isDarkMode: boolean) => void;
}

export default function Navbar({ showSettings = true, onSettingsClick, onThemeUpdate }: NavbarProps) {
  const [darkMode, setDarkMode] = useState(false);
  const [isFeatureRequestOpen, setIsFeatureRequestOpen] = useState(false);
  const pathname = usePathname();

  useEffect(() => {
    // Check for user's preference in localStorage
    const isDarkMode = localStorage.getItem('darkMode') === 'true';
    setDarkMode(isDarkMode);
    if (onThemeUpdate) {
      onThemeUpdate(isDarkMode);
    }
    if (isDarkMode) {
      document.documentElement.classList.add('dark');
    }
  }, [onThemeUpdate]);

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
    if (darkMode) {
      document.documentElement.classList.remove('dark');
      localStorage.setItem('darkMode', 'false');
    } else {
      document.documentElement.classList.add('dark');
      localStorage.setItem('darkMode', 'true');
    }
    if (onThemeUpdate) {
      onThemeUpdate(!darkMode);
    }
  };

  return (
    <nav className="bg-background sticky top-0 z-10 border-b">
      <div className="mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 justify-between">
          <div className="flex items-center space-x-4">
            <div className="flex flex-shrink-0 items-center">
              <Code className="text-primary h-8 w-8" />
              <span className="text-primary ml-2 text-2xl font-bold">Insights</span>
            </div>
            <div className="flex space-x-2">
              <Link href="/" passHref>
                <Button variant={pathname === '/' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <Home className="h-4 w-4" />
                  Home
                </Button>
              </Link>
              <Link href="/achievement-log" passHref>
                <Button variant={pathname === '/achievement-log' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <Trophy className="h-4 w-4" />
                  Achievements
                </Button>
              </Link>
              <Link href="/runbooks-docs" passHref>
                <Button variant={pathname === '/runbooks-docs' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <Book className="h-4 w-4" />
                  Runbooks docs
                </Button>
              </Link>
              <Link href="/composer" passHref>
                <Button variant={pathname === '/composer' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <PencilLine className="h-4 w-4" />
                  Composer
                </Button>
              </Link>
              <Link href="/learn" passHref>
                <Button variant={pathname === '/learn' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <BookA className="h-4 w-4" />
                  Learn
                </Button>
              </Link>
              <Link href="/customer-insights" passHref>
                <Button variant={pathname === '/customer-insights' ? 'default' : 'ghost'} size="sm" className="gap-2">
                  <BookA className="h-4 w-4" />
                  Customer Insights
                </Button>
              </Link>
            </div>
          </div>
          <div className="hidden sm:ml-6 sm:flex sm:items-center">
            {showSettings && onSettingsClick && (
              <Button variant="ghost" size="sm" className="gap-2" onClick={onSettingsClick}>
                <Settings className="h-4 w-4" />
                Settings
              </Button>
            )}
            <Link
              href="https://www.youtube.com/playlist?list=PL05JrBw4t0KrZcLI5g7nHVGb7shCcGonl"
              target="_blank"
              rel="noopener noreferrer"
              className="hover:text-foreground mr-4"
            >
              <Button variant="ghost" size="sm" className="gap-2">
                <Video className="h-4 w-4" />
                Videos
              </Button>
            </Link>
            <Link
              href="https://docs.google.com/presentation/d/1-HAKr6hA8hL6IcnzSAAUPcmQEg1jwNRkhgwMobORMcM/edit#slide=id.p"
              target="_blank"
              rel="noopener noreferrer"
              className="hover:text-foreground mr-4"
            >
              <Button variant="ghost" size="sm" className="gap-2">
                <Presentation className="h-4 w-4" />
                Slides
              </Button>
            </Link>
            <Link
              href="https://gitlab.fra1.qualtrics.com/jfe/form/SV_cYjJCFvBTu0xyqG"
              target="_blank"
              rel="noopener noreferrer"
              className="hover:text-foreground mr-4"
            >
              <Button variant="ghost" size="sm" className="gap-2">
                <MessageSquare className="h-4 w-4" />
                Feedback
              </Button>
            </Link>
            <Button variant="ghost" size="sm" className="gap-2" onClick={() => setIsFeatureRequestOpen(true)}>
              <BugPlay className="h-4 w-4" />
              Request Feature/Bug
            </Button>
            <Button variant="ghost" size="icon" onClick={toggleDarkMode} aria-label="Toggle dark mode">
              {darkMode ? <Sun className="h-5 w-5" /> : <Moon className="h-5 w-5" />}
            </Button>
            <Link
              href="https://gitlab.com/gitlab-org/ai-powered/ai-framework/code-context"
              target="_blank"
              rel="noopener noreferrer"
              className="hover:text-foreground inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium"
            >
              <Button variant="ghost" size="icon">
                <Gitlab className="h-5 w-5" />
              </Button>
            </Link>
            <Button variant="ghost" onClick={() => signOut()} className="hover:text-foreground">
              <LogOut className="mr-2 h-4 w-4" />
              Logout
            </Button>
          </div>
        </div>
      </div>
      <FeatureRequestDialog open={isFeatureRequestOpen} onOpenChange={setIsFeatureRequestOpen} />
    </nav>
  );
}
