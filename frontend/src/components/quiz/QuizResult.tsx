import { Trophy, CheckCircle, XCircle, Users } from 'lucide-react';
import { Card, Button } from '../ui';
import type { QuizSubmitResponse } from '../../types';
import { formatPercentage } from '../../utils/formatters';

const FacebookIcon = ({ className = "w-6 h-6 mr-3" }: { className?: string }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    fill="currentColor"
    className={className}
  >
    <path d="M12 2.04C6.5 2.04 2 6.53 2 12.06C2 17.06 5.66 21.21 10.44 21.96V14.96H7.9V12.06H10.44V9.85C10.44 7.34 11.93 5.96 14.22 5.96C15.31 5.96 16.45 6.15 16.45 6.15V8.62H15.19C13.95 8.62 13.56 9.39 13.56 10.18V12.06H16.34L15.89 14.96H13.56V21.96A10 10 0 0 0 22 12.06C22 6.53 17.5 2.04 12 2.04Z" />
  </svg>
);

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
            <div className="text-4xl font-display font-bold text-primary">{Math.round(accuracy)}</div>
            <div className="font-body font-medium">Total Score</div>
          </div>
          <div className="text-center p-4">
            <div className="text-4xl font-display font-bold text-secondary">{formatPercentage(accuracy, 0)}</div>
            <div className="font-body font-medium">Accuracy</div>
          </div>
        </div>
      </Card>

      <div className="flex flex-col sm:flex-row gap-4 mb-6">
        <Button onClick={onReviewAnswers} variant="outline" className="flex-1">Review Answers</Button>
        <Button onClick={onStartNewQuiz} className="flex-1">New Quiz</Button>
      </div>

      {/* Viral Loop / Call to Actions */}
      <div className="flex flex-col gap-4 mt-8 pt-6 border-t-3 border-black">
        <h3 className="font-display font-bold text-center uppercase">Next Steps</h3>

        {accuracy >= 50 ? (
          <a
            href={`https://www.facebook.com/sharer/sharer.php?u=https://acethatpaper.com&quote=I just scored ${Math.round(accuracy)}% on a JAMB Mock Exam on AceThatPaper! 🏆 Can you beat my score?`}
            target="_blank"
            rel="noopener noreferrer"
            className="w-full flex items-center justify-center p-4 bg-[#1877F2] text-white border-4 border-black shadow-brutal hover:shadow-none hover:translate-x-1 hover:translate-y-1 transition-all font-display font-bold text-lg"
          >
            <FacebookIcon />
            Share Your Score on Facebook
          </a>
        ) : null}

        <a
          href="https://web.facebook.com/share/g/19zEFhz12A/"
          target="_blank"
          rel="noopener noreferrer"
          className="w-full flex items-center justify-center p-4 bg-white text-black border-4 border-black shadow-brutal hover:bg-cream hover:shadow-none hover:translate-x-1 hover:translate-y-1 transition-all font-display font-bold text-lg"
        >
          <Users className="w-6 h-6 mr-3 text-[#1877F2]" />
          Join the Facebook Study Group
        </a>
      </div>
    </div>
  );
}

