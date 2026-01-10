import { Card, Badge } from '../ui';
import { OptionButton } from './OptionButton';
import type { QuizQuestion as QuizQuestionType } from '../../types';

interface QuestionCardProps {
  question: QuizQuestionType;
  questionNumber: number;
  totalQuestions: number;
  selectedOptionIds: number[];
  onSelectOption: (optionId: number) => void;
  showResult?: boolean;
  correctAnswer?: string;
}

export function QuestionCard({ question, questionNumber, totalQuestions, selectedOptionIds, onSelectOption, showResult = false, correctAnswer }: QuestionCardProps) {
  const handleOptionClick = (optionId: number) => {
    if (showResult) return;
    onSelectOption(optionId);
  };

  return (
    <Card className="bg-cream p-6 sm:p-8">
      <div className="flex flex-wrap items-center gap-3 mb-4">
        <Badge variant="primary">Question {questionNumber} of {totalQuestions}</Badge>
        {question.is_multiple_choice && <Badge variant="secondary">Multiple Choice</Badge>}
      </div>
      <h2 className="font-display text-xl sm:text-2xl font-bold mb-6 leading-relaxed">{question.question}</h2>
      <div className="space-y-3">
        {question.options.map((option, index) => {
          const isSelected = selectedOptionIds.includes(option.id);
          const isCorrectOption = showResult && correctAnswer === option.option;
          const isWrong = showResult && isSelected && !isCorrectOption;
          return (
            <OptionButton
              key={option.id}
              option={option}
              index={index}
              isSelected={isSelected}
              isCorrect={isCorrectOption}
              isWrong={isWrong}
              showResult={showResult}
              disabled={showResult}
              onClick={() => handleOptionClick(option.id)}
            />
          );
        })}
      </div>
    </Card>
  );
}

