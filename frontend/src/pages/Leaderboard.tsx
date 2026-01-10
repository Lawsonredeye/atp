import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Layout } from '../components/layout';
import { Button, Card, Alert } from '../components/ui';
import { LeaderboardTable } from '../components/leaderboard';
import { leaderboardService } from '../services';
import { useAuth } from '../context/AuthContext';
import type { LeaderboardEntry, UserRank } from '../types';
import { Trophy, RefreshCw, ArrowLeft, Target, TrendingUp, Users } from 'lucide-react';
import { formatScore, formatPercentage } from '../utils/formatters';

function Leaderboard() {
  const { state } = useAuth();
  const [entries, setEntries] = useState<LeaderboardEntry[]>([]);
  const [userRank, setUserRank] = useState<UserRank | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchLeaderboard = async () => {
    try {
      setIsLoading(true);
      setError(null);

      const [leaderboardData, userRankData] = await Promise.all([
        leaderboardService.getLeaderboard(50, 0),
        leaderboardService.getMyRank().catch(() => null),
      ]);

      setEntries(leaderboardData);
      setUserRank(userRankData);
    } catch (err) {
      setError('Failed to load leaderboard. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchLeaderboard();
  }, []);

  return (
    <Layout>
      <div className="px-4 sm:px-6 py-8 sm:py-12">
        <div className="max-w-5xl mx-auto">
          {/* Header */}
          <div className="mb-8">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
              <div className="flex items-center gap-4">
                <div className="w-14 h-14 bg-accent-yellow border-4 border-black shadow-brutal flex items-center justify-center">
                  <Trophy className="w-7 h-7" />
                </div>
                <div>
                  <h1 className="font-display text-3xl sm:text-4xl font-bold">Leaderboard</h1>
                  <p className="font-body text-gray-600">Top performers nationwide</p>
                </div>
              </div>
              <div className="flex gap-3">
                <Button variant="outline" onClick={fetchLeaderboard} disabled={isLoading}>
                  <RefreshCw className={`inline mr-2 w-4 h-4 ${isLoading ? 'animate-spin' : ''}`} />
                  Refresh
                </Button>
                <Link to="/dashboard">
                  <Button variant="secondary">
                    <ArrowLeft className="inline mr-2 w-4 h-4" />
                    Dashboard
                  </Button>
                </Link>
              </div>
            </div>

            {/* User Stats Card */}
            {userRank && (
              <Card className="p-4 sm:p-6 bg-accent-yellow mb-6">
                <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
                  <div className="text-center">
                    <div className="flex items-center justify-center gap-2 mb-1">
                      <Trophy className="w-5 h-5" />
                      <span className="font-body text-sm uppercase font-medium">Your Rank</span>
                    </div>
                    <div className="font-display text-3xl font-bold">#{userRank.rank}</div>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center gap-2 mb-1">
                      <Target className="w-5 h-5" />
                      <span className="font-body text-sm uppercase font-medium">Total Score</span>
                    </div>
                    <div className="font-display text-3xl font-bold">{formatScore(userRank.total_score)}</div>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center gap-2 mb-1">
                      <TrendingUp className="w-5 h-5" />
                      <span className="font-body text-sm uppercase font-medium">Accuracy</span>
                    </div>
                    <div className="font-display text-3xl font-bold text-accent-green">{formatPercentage(userRank.accuracy_percent, 0)}</div>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center gap-2 mb-1">
                      <Users className="w-5 h-5" />
                      <span className="font-body text-sm uppercase font-medium">Quizzes</span>
                    </div>
                    <div className="font-display text-3xl font-bold">{userRank.total_quizzes}</div>
                  </div>
                </div>
              </Card>
            )}
          </div>

          {/* Error Alert */}
          {error && (
            <Alert variant="error" className="mb-6">
              {error}
            </Alert>
          )}

          {/* Leaderboard Table */}
          <LeaderboardTable
            entries={entries}
            currentUserId={state.user?.id}
            userRank={userRank}
            isLoading={isLoading}
          />

          {/* Legend */}
          {!isLoading && entries.length > 0 && (
            <Card className="mt-6 p-4 bg-cream">
              <h3 className="font-display font-bold text-sm uppercase mb-3">How Rankings Work</h3>
              <div className="grid grid-cols-1 sm:grid-cols-3 gap-4 font-body text-sm">
                <div className="flex items-center gap-2">
                  <div className="w-4 h-4 bg-accent-green border-2 border-black" />
                  <span>80%+ Accuracy = Excellent</span>
                </div>
                <div className="flex items-center gap-2">
                  <div className="w-4 h-4 bg-primary border-2 border-black" />
                  <span>50-79% Accuracy = Good</span>
                </div>
                <div className="flex items-center gap-2">
                  <div className="w-4 h-4 bg-accent-red border-2 border-black" />
                  <span>Below 50% = Needs Work</span>
                </div>
              </div>
              <p className="mt-3 text-gray-600">
                Rankings are based on total score. Keep practicing to climb the leaderboard!
              </p>
            </Card>
          )}

          {/* CTA for non-ranked users */}
          {!isLoading && !userRank && (
            <Card className="mt-6 p-6 bg-secondary text-white text-center">
              <Trophy className="w-12 h-12 mx-auto mb-4 text-accent-yellow" />
              <h3 className="font-display text-2xl font-bold mb-2">Get on the Leaderboard!</h3>
              <p className="font-body mb-4 text-white/80">
                Complete your first quiz to appear on the leaderboard and compete with other students.
              </p>
              <Link to="/dashboard">
                <Button className="bg-accent-yellow text-black border-white shadow-[4px_4px_0px_0px_#FFFFFF] hover:bg-white">
                  Start a Quiz Now
                </Button>
              </Link>
            </Card>
          )}
        </div>
      </div>
    </Layout>
  );
}

export default Leaderboard;

