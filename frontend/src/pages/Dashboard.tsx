import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Layout } from '../components/layout';
import { Button, Card, LoadingSpinner, Alert, Select } from '../components/ui';
import { useAuth } from '../context/AuthContext';
import { subjectService } from '../services';
import type { Subject } from '../types';
import {
  BookOpen,
  Trophy,
  Target,
  TrendingUp,
  CheckCircle,
  XCircle,
  Play,
  Award,
  BarChart3,
} from 'lucide-react';
import { formatPercentage, formatScore } from '../utils/formatters';

function Dashboard() {
  const { state } = useAuth();
  const navigate = useNavigate();
  const [subjects, setSubjects] = useState<Subject[]>([]);
  const [selectedSubject, setSelectedSubject] = useState<string>('');
  const [numQuestions, setNumQuestions] = useState<string>('10');
  const [loadingSubjects, setLoadingSubjects] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const user = state.user;

  useEffect(() => {
    const fetchSubjects = async () => {
      try {
        const data = await subjectService.getAllSubjects();
        setSubjects(data);
      } catch (err) {
        setError('Failed to load subjects. Please refresh the page.');
      } finally {
        setLoadingSubjects(false);
      }
    };
    fetchSubjects();
  }, []);

  const handleStartQuiz = () => {
    if (selectedSubject) {
      navigate(`/quiz?subject=${selectedSubject}&questions=${numQuestions}`);
    }
  };

  if (!user) {
    return (
      <Layout>
        <div className="flex items-center justify-center min-h-[60vh]">
          <LoadingSpinner size="lg" />
        </div>
      </Layout>
    );
  }

  const accuracy = user.total_questions_answered > 0
    ? (user.total_correct_answers / user.total_questions_answered) * 100
    : 0;

  const stats = [
    {
      label: 'Total Quizzes',
      value: user.total_quizzes_taken,
      icon: BookOpen,
      color: 'bg-primary text-black',
    },
    {
      label: 'Questions Answered',
      value: formatScore(user.total_questions_answered),
      icon: Target,
      color: 'bg-secondary text-white',
    },
    {
      label: 'Correct Answers',
      value: formatScore(user.total_correct_answers),
      icon: CheckCircle,
      color: 'bg-accent-green text-black',
    },
    {
      label: 'Accuracy',
      value: formatPercentage(accuracy, 1),
      icon: TrendingUp,
      color: 'bg-accent-yellow text-black',
    },
  ];

  const questionOptions = [
    { value: '5', label: '5 Questions' },
    { value: '10', label: '10 Questions' },
    { value: '15', label: '15 Questions' },
    { value: '20', label: '20 Questions' },
    { value: '30', label: '30 Questions' },
    { value: '50', label: '50 Questions' },
  ];

  return (
    <Layout>
      <div className="px-4 sm:px-6 py-8 sm:py-12">
        <div className="max-w-7xl mx-auto">
          {/* Welcome Section */}
          <div className="mb-8">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
              <div>
                <h1 className="font-display text-3xl sm:text-4xl font-bold mb-2">
                  Welcome back, <span className="text-primary">{user.full_name.split(' ')[0]}!</span>
                </h1>
                <p className="font-body text-gray-600 text-lg">
                  Ready to continue your JAMB prep? Let's get started!
                </p>
              </div>
              <Link to="/leaderboard">
                <Button variant="secondary">
                  <Trophy className="inline mr-2 w-5 h-5" />
                  View Leaderboard
                </Button>
              </Link>
            </div>
          </div>

          {/* Stats Grid */}
          <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
            {stats.map((stat, index) => (
              <Card key={index} className={`${stat.color} p-4 sm:p-6`}>
                <div className="flex items-center gap-3 mb-2">
                  <stat.icon className="w-6 h-6" />
                  <span className="font-body text-sm uppercase font-medium opacity-80">
                    {stat.label}
                  </span>
                </div>
                <div className="font-display text-2xl sm:text-3xl font-bold">
                  {stat.value}
                </div>
              </Card>
            ))}
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
            {/* Start Quiz Section */}
            <div className="lg:col-span-2">
              <Card className="p-6 sm:p-8">
                <div className="flex items-center gap-4 mb-6">
                  <div className="w-14 h-14 bg-primary border-4 border-black shadow-brutal flex items-center justify-center">
                    <Play className="w-7 h-7" />
                  </div>
                  <div>
                    <h2 className="font-display text-2xl font-bold">Start a New Quiz</h2>
                    <p className="font-body text-gray-600">Choose a subject and number of questions</p>
                  </div>
                </div>

                {error && (
                  <Alert variant="error" className="mb-6">
                    {error}
                  </Alert>
                )}

                {loadingSubjects ? (
                  <LoadingSpinner />
                ) : (
                  <div className="space-y-4">
                    <Select
                      label="Select Subject"
                      options={subjects.map((s) => ({ value: s.id.toString(), label: s.name }))}
                      placeholder="Choose a subject"
                      value={selectedSubject}
                      onChange={(e) => setSelectedSubject(e.target.value)}
                    />

                    <Select
                      label="Number of Questions"
                      options={questionOptions}
                      value={numQuestions}
                      onChange={(e) => setNumQuestions(e.target.value)}
                    />

                    <Button
                      onClick={handleStartQuiz}
                      disabled={!selectedSubject}
                      size="lg"
                      className="w-full"
                    >
                      <Play className="inline mr-2 w-5 h-5" />
                      Start Quiz
                    </Button>
                  </div>
                )}
              </Card>
            </div>

            {/* Performance Summary */}
            <div>
              <Card className="p-6 bg-cream">
                <div className="flex items-center gap-3 mb-6">
                  <div className="w-12 h-12 bg-accent-yellow border-3 border-black shadow-brutal-sm flex items-center justify-center">
                    <BarChart3 className="w-6 h-6" />
                  </div>
                  <h2 className="font-display text-xl font-bold">Performance</h2>
                </div>

                <div className="space-y-4">
                  {/* Accuracy Bar */}
                  <div>
                    <div className="flex justify-between mb-2">
                      <span className="font-body font-medium">Overall Accuracy</span>
                      <span className="font-display font-bold text-primary">
                        {formatPercentage(accuracy, 1)}
                      </span>
                    </div>
                    <div className="h-4 bg-white border-3 border-black">
                      <div
                        className="h-full bg-accent-green transition-all duration-500"
                        style={{ width: `${Math.min(accuracy, 100)}%` }}
                      />
                    </div>
                  </div>

                  {/* Correct vs Incorrect */}
                  <div className="grid grid-cols-2 gap-3">
                    <div className="p-4 bg-accent-green border-3 border-black text-center">
                      <CheckCircle className="w-6 h-6 mx-auto mb-1" />
                      <div className="font-display font-bold text-xl">
                        {formatScore(user.total_correct_answers)}
                      </div>
                      <div className="font-body text-sm">Correct</div>
                    </div>
                    <div className="p-4 bg-accent-red text-white border-3 border-black text-center">
                      <XCircle className="w-6 h-6 mx-auto mb-1" />
                      <div className="font-display font-bold text-xl">
                        {formatScore(user.total_incorrect_answers)}
                      </div>
                      <div className="font-body text-sm">Incorrect</div>
                    </div>
                  </div>

                  {/* Quick Actions */}
                  <div className="pt-4 border-t-3 border-black">
                    <Link to="/leaderboard" className="block">
                      <div className="p-4 bg-secondary text-white border-3 border-black hover:shadow-brutal transition-all duration-100 flex items-center justify-between">
                        <div className="flex items-center gap-3">
                          <Award className="w-6 h-6 text-accent-yellow" />
                          <span className="font-display font-bold">View Rankings</span>
                        </div>
                        <span className="text-2xl">→</span>
                      </div>
                    </Link>
                  </div>
                </div>
              </Card>
            </div>
          </div>

          {/* Quick Subject Cards */}
          {subjects.length > 0 && (
            <div className="mt-8">
              <h2 className="font-display text-2xl font-bold mb-4">Quick Start</h2>
              <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
                {subjects.slice(0, 6).map((subject, index) => {
                  const colors = [
                    'bg-accent-yellow text-black',
                    'bg-accent-green text-black',
                    'bg-primary text-black',
                    'bg-secondary text-white',
                    'bg-accent-purple text-white',
                    'bg-accent-red text-white',
                  ];
                  const colorClass = colors[index % colors.length];

                  return (
                    <Card
                      key={subject.id}
                      hoverable
                      className={`text-center py-6 cursor-pointer ${colorClass}`}
                      onClick={() => {
                        setSelectedSubject(subject.id.toString());
                        window.scrollTo({ top: 0, behavior: 'smooth' });
                      }}
                    >
                      <h3 className="font-display font-bold text-lg">{subject.name}</h3>
                    </Card>
                  );
                })}
              </div>
            </div>
          )}
        </div>
      </div>
    </Layout>
  );
}

export default Dashboard;

