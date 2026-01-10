import { Trophy, Medal, Award } from 'lucide-react';

interface RankBadgeProps {
  rank: number;
  size?: 'sm' | 'md' | 'lg';
}

export function RankBadge({ rank, size = 'md' }: RankBadgeProps) {
  const sizeClasses = {
    sm: 'w-8 h-8 text-sm',
    md: 'w-12 h-12 text-lg',
    lg: 'w-16 h-16 text-2xl',
  };

  const iconSizes = {
    sm: 'w-4 h-4',
    md: 'w-6 h-6',
    lg: 'w-8 h-8',
  };

  if (rank === 1) {
    return (
      <div className={`${sizeClasses[size]} bg-accent-yellow border-4 border-black shadow-brutal flex items-center justify-center`}>
        <Trophy className={`${iconSizes[size]} text-black`} />
      </div>
    );
  }

  if (rank === 2) {
    return (
      <div className={`${sizeClasses[size]} bg-gray-200 border-4 border-black shadow-brutal flex items-center justify-center`}>
        <Medal className={`${iconSizes[size]} text-gray-600`} />
      </div>
    );
  }

  if (rank === 3) {
    return (
      <div className={`${sizeClasses[size]} bg-amber-500 border-4 border-black shadow-brutal flex items-center justify-center`}>
        <Award className={`${iconSizes[size]} text-white`} />
      </div>
    );
  }

  return (
    <div className={`${sizeClasses[size]} bg-white border-4 border-black shadow-brutal flex items-center justify-center font-display font-bold`}>
      #{rank}
    </div>
  );
}

