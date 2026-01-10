interface QuizProgressProps {
  total: number;
  answeredCount: number;
}

export function QuizProgress({ total, answeredCount }: QuizProgressProps) {
  const progressPercent = (answeredCount / total) * 100;
  return (
    <div className="mb-6">
      <div className="flex justify-between items-center mb-2">
        <span className="font-display font-bold text-sm uppercase">Progress: {answeredCount}/{total} answered</span>
        <span className="font-display font-bold text-sm">{Math.round(progressPercent)}%</span>
      </div>
      <div className="h-4 bg-white border-3 border-black">
        <div className="h-full bg-primary transition-all duration-300" style={{ width: `${progressPercent}%` }} />
      </div>
    </div>
  );
}

