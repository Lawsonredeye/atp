import type { QuizOption } from '../../types';

interface OptionButtonProps {
  option: QuizOption;
  index: number;
  isSelected: boolean;
  isCorrect?: boolean;
  isWrong?: boolean;
  showResult?: boolean;
  disabled?: boolean;
  onClick: () => void;
}

const optionLabels = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'];

export function OptionButton({ option, index, isSelected, isCorrect = false, isWrong = false, showResult = false, disabled = false, onClick }: OptionButtonProps) {
  let className = 'w-full p-4 mb-3 text-left border-4 border-black font-body text-lg transition-all duration-100';

  if (showResult) {
    if (isCorrect) {
      className += ' bg-accent-green shadow-brutal font-bold';
    } else if (isWrong) {
      className += ' bg-accent-red text-white shadow-brutal font-bold';
    } else {
      className += ' bg-white shadow-brutal-sm opacity-60';
    }
  } else if (isSelected) {
    className += ' bg-primary shadow-brutal font-bold';
  } else {
    className += ' bg-white shadow-brutal-sm hover:bg-cream hover:shadow-brutal cursor-pointer';
  }

  if (disabled) {
    className += ' cursor-not-allowed';
  }

  return (
    <button type="button" className={className} onClick={onClick} disabled={disabled} aria-pressed={isSelected}>
      <span className="font-display font-bold mr-3">{optionLabels[index]}.</span>
      {option.option}
      {showResult && isCorrect && <span className="ml-2">✓</span>}
      {showResult && isWrong && <span className="ml-2">✗</span>}
    </button>
  );
}
