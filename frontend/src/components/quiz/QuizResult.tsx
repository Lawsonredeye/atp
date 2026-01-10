import { Trophy, CheckCircle, XCircle } from 'lucide-react';
import { Card, Button } from '../ui';
import type { QuizSubmitResponse } from '../../types';
import { formatPercentage } from '../../utils/formatters';

interface QuizResultProps {
  result: QuizSubmitResponse;
  onReviewAnswers: () => void;
  onStartNewQuiz: () => void;
}

export function QuizResultCard({ result, onReviewAnswers, onStartNewQuiz }: QuizResultProps) {
  const accuracy = (result.correct_answers / result.total_questions) * 100;
  let performanceMessage = '';
  let performanceColor = '';

  if (accuracy >= 80) {
    performanceMessage = 'Excellent! Keep it up! 🎉';
    performanceColor = 'text-accent-green';
  } else if (accuracy >= 60) {
    performanceMessage = "Good job! You're improving! 💪";
    performanceColor = 'text-primary';
  } else if (accuracy >= 40) {
    performanceMessage = "Keep practicing! You'll get better! 📚";
    performanceColor = 'text-accent-yellow';
  } else {
    performanceMessage = "Don't give up! Review and try again! 🔄";
    performanceColor = 'text-accent-red';
  }

  return (
    <div className="max-w-2xl mx-auto">
      <Card className="bg-secondary text-white text-center p-8 mb-6">
        <Trophy className="w-16 h-16 mx-auto mb-4 text-accent-yellow" />
        <h1 className="font-display text-3xl sm:text-4xl font-bold uppercase mb-2">Quiz Complete!</h1>
        <p className={`font-body text-xl ${performanceColor}`}>{performanceMessage}</p>
      </Card>

      <div className="grid grid-cols-2 gap-4 mb-6">
        <Card className="bg-accent-green text-black text-center">
          <CheckCircle className="w-10 h-10 mx-auto mb-2" />
          <div className="text-3xl font-display font-bold">{result.correct_answers}</div>
          <div className="font-body font-medium">Correct</div>
        </Card>
        <Card className="bg-accent-red text-white text-center">
          <XCircle className="w-10 h-10 mx-auto mb-2" />
          <div className="text-3xl font-display font-bold">{result.incorrect_answers}</div>
          <div className="font-body font-medium">Incorrect</div>
        </Card>
      </div>

      <Card className="mb-6">
        <div className="grid grid-cols-2 gap-4">
          <div className="text-center p-4 border-r-3 border-black">
            <div className="text-4xl font-display font-bold text-primary">{result.score}</div>
            <div className="font-body font-medium">Total Score</div>
          </div>
          <div className="text-center p-4">
            <div className="text-4xl font-display font-bold text-secondary">{formatPercentage(accuracy, 0)}</div>
            <div className="font-body font-medium">Accuracy</div>
          </div>
        </div>
      </Card>

      <div className="flex flex-col sm:flex-row gap-4">
        <Button onClick={onReviewAnswers} variant="outline" className="flex-1">Review Answers</Button>
        <Button onClick={onStartNewQuiz} className="flex-1">New Quiz</Button>
      </div>
    </div>
  );
}

