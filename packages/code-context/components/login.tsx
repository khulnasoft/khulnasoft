'use client';

import { Code } from 'lucide-react';
import { signIn } from 'next-auth/react';
import { TypeAnimation } from 'react-type-animation';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter } from '@/components/ui/card';

export default function Login() {
  return (
    <div className="bg-background flex min-h-screen items-center justify-center">
      <Card className="w-full max-w-md">
        <CardContent className="flex flex-col items-center space-y-6 pt-6">
          <div className="flex items-center space-x-2">
            <Code className="text-primary h-10 w-10" />
            <TypeAnimation
              sequence={['Insights', 1000]}
              wrapper="h2"
              cursor={true}
              repeat={Infinity}
              className="text-primary text-2xl font-bold"
            />
          </div>
          <p className="text-muted-foreground text-center">
            AI powered 2nd pair of eyes for your code reviews and issue analysis.
          </p>
        </CardContent>
        <CardFooter>
          <Button className="w-full" size="lg" onClick={() => signIn('gitlab')}>
            Login with GitLab
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}
