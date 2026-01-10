import type { LeaderboardEntry, UserRank } from '../../types';
import { Trophy, Medal, Award } from 'lucide-react';
import { formatScore, formatPercentage } from '../../utils/formatters';

interface LeaderboardTableProps {
  entries: LeaderboardEntry[];
  currentUserId?: number;
  userRank?: UserRank | null;
  isLoading?: boolean;
}

function RankIcon({ rank }: { rank: number }) {
  if (rank === 1) {
    return <Trophy className="w-6 h-6 text-accent-yellow" />;
  }
  if (rank === 2) {
    return <Medal className="w-6 h-6 text-gray-400" />;
  }
  if (rank === 3) {
    return <Award className="w-6 h-6 text-amber-600" />;
  }
  return <span className="font-display font-bold text-xl">#{rank}</span>;
}

function SkeletonRow() {
  return (
    <div className="bg-white grid grid-cols-12 gap-4 p-4 border-t-3 border-black animate-pulse">
      <div className="col-span-2 sm:col-span-1"><div className="h-6 bg-gray-200 rounded w-8" /></div>
      <div className="col-span-5 sm:col-span-5"><div className="h-6 bg-gray-200 rounded w-3/4" /></div>
      <div className="col-span-2 sm:col-span-3"><div className="h-6 bg-gray-200 rounded w-1/2" /></div>
      <div className="col-span-3 sm:col-span-3"><div className="h-6 bg-gray-200 rounded w-1/2" /></div>
    </div>
  );
}

export function LeaderboardTable({ entries, currentUserId, userRank, isLoading }: LeaderboardTableProps) {
  if (isLoading) {
    return (
      <div className="border-4 border-black shadow-brutal-lg overflow-hidden">
        <div className="bg-secondary text-white grid grid-cols-12 gap-4 p-4 font-display font-bold uppercase text-sm sm:text-base">
          <div className="col-span-2 sm:col-span-1">Rank</div>
          <div className="col-span-5 sm:col-span-5">Name</div>
          <div className="col-span-2 sm:col-span-3">Score</div>
          <div className="col-span-3 sm:col-span-3">Accuracy</div>
        </div>
        {[...Array(5)].map((_, i) => <SkeletonRow key={i} />)}
      </div>
    );
  }

  if (entries.length === 0) {
    return (
      <div className="border-4 border-black shadow-brutal p-8 text-center bg-white">
        <Trophy className="w-16 h-16 mx-auto mb-4 text-gray-300" />
        <h3 className="font-display text-xl font-bold mb-2">No Rankings Yet</h3>
        <p className="font-body text-gray-600">Be the first to complete a quiz and appear on the leaderboard!</p>
      </div>
    );
  }

  // Check if current user is in top entries, if not add them
  const showUserRankSeparately = userRank && !entries.some(e => e.user_id === currentUserId);

  return (
    <div className="border-4 border-black shadow-brutal-lg overflow-hidden">
      {/* Table Header */}
      <div className="bg-secondary text-white grid grid-cols-12 gap-2 sm:gap-4 p-3 sm:p-4 font-display font-bold uppercase text-xs sm:text-base">
        <div className="col-span-2 sm:col-span-1">Rank</div>
        <div className="col-span-4 sm:col-span-5">Name</div>
        <div className="col-span-3 sm:col-span-3">Score</div>
        <div className="col-span-3 sm:col-span-3">Accuracy</div>
      </div>

      {/* User's Rank (if not in top entries) */}
      {showUserRankSeparately && (
        <>
          <div className="bg-accent-yellow grid grid-cols-12 gap-2 sm:gap-4 p-3 sm:p-4 border-t-4 border-black">
            <div className="col-span-2 sm:col-span-1 flex items-center">
              <span className="font-display font-bold text-lg sm:text-xl">#{userRank.rank}</span>
            </div>
            <div className="col-span-4 sm:col-span-5 font-body font-bold flex items-center gap-2">
              <span className="truncate">{userRank.user_name}</span>
              <span className="text-xs bg-black text-white px-2 py-0.5 uppercase">You</span>
            </div>
            <div className="col-span-3 sm:col-span-3 font-display font-bold">
              {formatScore(userRank.total_score)}
            </div>
            <div className="col-span-3 sm:col-span-3 font-display font-bold text-accent-green">
              {formatPercentage(userRank.accuracy_percent, 0)}
            </div>
          </div>
          <div className="bg-gray-100 text-center py-2 border-t-2 border-dashed border-black">
            <span className="font-body text-sm text-gray-500">• • •</span>
          </div>
        </>
      )}

      {/* Leaderboard Entries */}
      {entries.map((entry, index) => {
        const isCurrentUser = entry.user_id === currentUserId;
        const isTopThree = entry.rank <= 3;

        let rowClass = 'bg-white';
        if (isCurrentUser) {
          rowClass = 'bg-accent-yellow';
        } else if (entry.rank === 1) {
          rowClass = 'bg-gradient-to-r from-yellow-100 to-white';
        } else if (entry.rank === 2) {
          rowClass = 'bg-gradient-to-r from-gray-100 to-white';
        } else if (entry.rank === 3) {
          rowClass = 'bg-gradient-to-r from-amber-100 to-white';
        }

        return (
          <div
            key={entry.user_id}
            className={`${rowClass} grid grid-cols-12 gap-2 sm:gap-4 p-3 sm:p-4 ${
              index === 0 && !showUserRankSeparately ? 'border-t-4' : 'border-t-3'
            } border-black transition-colors hover:bg-cream`}
          >
            <div className="col-span-2 sm:col-span-1 flex items-center justify-center sm:justify-start">
              {isTopThree ? (
                <div className="w-8 h-8 flex items-center justify-center">
                  <RankIcon rank={entry.rank} />
                </div>
              ) : (
                <span className="font-display font-bold text-lg">#{entry.rank}</span>
              )}
            </div>
            <div className="col-span-4 sm:col-span-5 font-body flex items-center gap-2">
              <span className={`truncate ${isCurrentUser ? 'font-bold' : ''}`}>
                {entry.user_name}
              </span>
              {isCurrentUser && (
                <span className="text-xs bg-black text-white px-2 py-0.5 uppercase flex-shrink-0">You</span>
              )}
              {entry.rank === 1 && <span className="flex-shrink-0">🏆</span>}
            </div>
            <div className="col-span-3 sm:col-span-3 font-display font-bold">
              {formatScore(entry.total_score)}
            </div>
            <div className={`col-span-3 sm:col-span-3 font-display font-bold ${
              entry.accuracy_percent >= 80 ? 'text-accent-green' : 
              entry.accuracy_percent >= 50 ? 'text-primary' : 'text-accent-red'
            }`}>
              {formatPercentage(entry.accuracy_percent, 0)}
            </div>
          </div>
        );
      })}
    </div>
  );
}

export default LeaderboardTable;

