'use client';

import { useState } from 'react';
import { Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { FlashCard } from '@/app/lib/actions/learn/flashcards';
import Markdown from '../ui/markdown';

export default function Flashcards({ flashcards }: { flashcards: FlashCard[] }) {
  const [currentCard, setCurrentCard] = useState(0);
  const [showAnswer, setShowAnswer] = useState(false);

  const nextCard = () => {
    setCurrentCard((prev) => (prev + 1) % flashcards.length);
    setShowAnswer(false);
  };

  const prevCard = () => {
    setCurrentCard((prev) => (prev - 1 + flashcards.length) % flashcards.length);
    setShowAnswer(false);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Flashcards</CardTitle>
      </CardHeader>
      <CardContent>
        {!flashcards || flashcards.length === 0 ? (
          <div className="flex min-h-32 items-center justify-center">
            <Loader2 className="h-8 w-8 animate-spin text-gray-400" />
          </div>
        ) : (
          <div>
            <h3 className="mb-4 text-xl font-semibold">
              {showAnswer ? 'Answer' : 'Question'} ({currentCard + 1}/{flashcards.length})
            </h3>
            <div className="mb-4 min-h-24">
              {showAnswer ? <Markdown contents={flashcards[currentCard].answer} /> : flashcards[currentCard].question}
            </div>
            <div className="mb-4 flex justify-center space-x-4">
              <Button onClick={prevCard}>Previous</Button>
              <Button onClick={() => setShowAnswer(!showAnswer)}>{showAnswer ? 'Show Question' : 'Show Answer'}</Button>
              <Button onClick={nextCard}>Next</Button>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
