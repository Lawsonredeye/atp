import { Link } from 'react-router-dom';
import { Layout } from '../components/layout';
import { Button, Card } from '../components/ui';
import { BookOpen, Trophy, Target, Users, CheckCircle, Zap, BarChart3, ArrowRight } from 'lucide-react';
import { useAuth } from '../context/AuthContext';

const features = [
  {
    icon: BookOpen,
    title: 'JAMB Past Questions',
    description: 'Practice with thousands of curated past questions from previous JAMB exams.',
    color: 'bg-accent-yellow',
  },
  {
    icon: Target,
    title: 'Subject-Based Quizzes',
    description: 'Focus on specific subjects to strengthen your weak areas.',
    color: 'bg-accent-green',
  },
  {
    icon: Trophy,
    title: 'Leaderboard',
    description: 'Compete with other students and track your ranking nationwide.',
    color: 'bg-primary',
  },
  {
    icon: BarChart3,
    title: 'Progress Tracking',
    description: 'Monitor your performance with detailed analytics and insights.',
    color: 'bg-accent-purple',
  },
];

const subjects = [
  { name: 'Mathematics', count: 500, color: 'bg-accent-yellow text-black' },
  { name: 'English', count: 450, color: 'bg-accent-green text-black' },
  { name: 'Physics', count: 400, color: 'bg-primary text-black' },
  { name: 'Chemistry', count: 380, color: 'bg-secondary text-white' },
  { name: 'Biology', count: 420, color: 'bg-accent-purple text-white' },
  { name: 'Government', count: 350, color: 'bg-accent-red text-white' },
];

function Home() {
  const { state } = useAuth();

  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-cream px-4 sm:px-6 py-12 sm:py-16 lg:py-24">
        <div className="max-w-7xl mx-auto">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12 items-center">
            <div>
              <span className="inline-block px-4 py-2 bg-accent-yellow border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-6">
                🎓 #1 JAMB Prep Platform
              </span>
              <h1 className="font-display text-4xl sm:text-5xl lg:text-6xl font-bold leading-tight mb-6">
                Ace Your <span className="text-primary">JAMB</span> Exams With{' '}
                <span className="relative">
                  Confidence
                  <svg className="absolute -bottom-2 left-0 w-full h-3" viewBox="0 0 200 12" preserveAspectRatio="none">
                    <path d="M0,6 Q50,0 100,6 T200,6" fill="none" stroke="#FF6B35" strokeWidth="4" />
                  </svg>
                </span>
              </h1>
              <p className="font-body text-lg sm:text-xl text-gray-700 mb-8 leading-relaxed">
                Practice with thousands of past questions, track your progress, and compete with students across Nigeria. Your success starts here!
              </p>
              <div className="flex flex-col sm:flex-row gap-4">
                {state.isAuthenticated ? (
                  <Link to="/dashboard">
                    <Button size="lg" className="w-full sm:w-auto">
                      Go to Dashboard <ArrowRight className="inline ml-2 w-5 h-5" />
                    </Button>
                  </Link>
                ) : (
                  <>
                    <Link to="/register">
                      <Button size="lg" className="w-full sm:w-auto">
                        Start Practicing Free <ArrowRight className="inline ml-2 w-5 h-5" />
                      </Button>
                    </Link>
                    <Link to="/login">
                      <Button variant="outline" size="lg" className="w-full sm:w-auto">
                        Login
                      </Button>
                    </Link>
                  </>
                )}
              </div>
              <div className="flex items-center gap-6 mt-8 pt-6 border-t-3 border-black">
                <div className="flex items-center gap-2">
                  <CheckCircle className="w-6 h-6 text-accent-green" />
                  <span className="font-body font-medium">10,000+ Questions</span>
                </div>
                <div className="flex items-center gap-2">
                  <Users className="w-6 h-6 text-secondary" />
                  <span className="font-body font-medium">50,000+ Students</span>
                </div>
              </div>
            </div>
            <div className="relative hidden lg:block">
              <div className="absolute -top-4 -left-4 w-full h-full bg-primary border-4 border-black" />
              <Card className="relative z-10 p-8 bg-white">
                <div className="text-center mb-6">
                  <div className="inline-flex items-center justify-center w-20 h-20 bg-accent-yellow border-4 border-black shadow-brutal mb-4">
                    <Zap className="w-10 h-10" />
                  </div>
                  <h3 className="font-display text-2xl font-bold">Quick Stats</h3>
                </div>
                <div className="space-y-4">
                  <div className="flex justify-between items-center p-4 bg-cream border-3 border-black">
                    <span className="font-body font-medium">Questions Answered</span>
                    <span className="font-display font-bold text-2xl text-primary">2.5M+</span>
                  </div>
                  <div className="flex justify-between items-center p-4 bg-cream border-3 border-black">
                    <span className="font-body font-medium">Quizzes Completed</span>
                    <span className="font-display font-bold text-2xl text-secondary">150K+</span>
                  </div>
                  <div className="flex justify-between items-center p-4 bg-cream border-3 border-black">
                    <span className="font-body font-medium">Avg. Improvement</span>
                    <span className="font-display font-bold text-2xl text-accent-green">+35%</span>
                  </div>
                </div>
              </Card>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="bg-white px-4 sm:px-6 py-12 sm:py-16 lg:py-20 border-y-4 border-black">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-12">
            <span className="inline-block px-4 py-2 bg-secondary text-white border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
              Why Choose Us
            </span>
            <h2 className="font-display text-3xl sm:text-4xl lg:text-5xl font-bold">
              Everything You Need to <span className="text-primary">Succeed</span>
            </h2>
          </div>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {features.map((feature, index) => (
              <Card
                key={index}
                hoverable
                className={`${feature.color} transition-all duration-100`}
              >
                <feature.icon className="w-12 h-12 mb-4 border-3 border-black p-2 bg-white" />
                <h3 className="font-display text-xl font-bold mb-2">{feature.title}</h3>
                <p className="font-body text-black/80">{feature.description}</p>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Subjects Section */}
      <section className="bg-cream px-4 sm:px-6 py-12 sm:py-16 lg:py-20">
        <div className="max-w-7xl mx-auto">
          <div className="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-12">
            <div>
              <span className="inline-block px-4 py-2 bg-primary border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
                Available Subjects
              </span>
              <h2 className="font-display text-3xl sm:text-4xl font-bold">
                Practice Any Subject
              </h2>
            </div>
            {state.isAuthenticated ? (
              <Link to="/quiz">
                <Button variant="secondary">Browse All Subjects</Button>
              </Link>
            ) : (
              <Link to="/register">
                <Button variant="secondary">Get Started</Button>
              </Link>
            )}
          </div>
          <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
            {subjects.map((subject, index) => (
              <Card
                key={index}
                hoverable
                className={`${subject.color} text-center py-8`}
              >
                <h3 className="font-display text-lg font-bold mb-1">{subject.name}</h3>
                <p className="font-body text-sm opacity-80">{subject.count}+ Questions</p>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="bg-secondary px-4 sm:px-6 py-12 sm:py-16 lg:py-20 border-t-4 border-black">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-white mb-6">
            Ready to Start Your <span className="text-primary">JAMB</span> Prep Journey?
          </h2>
          <p className="font-body text-lg sm:text-xl text-white/80 mb-8 max-w-2xl mx-auto">
            Join thousands of students who have improved their scores with ScoreThatExam. It's completely free to get started!
          </p>
          {state.isAuthenticated ? (
            <Link to="/quiz">
              <Button size="lg" className="bg-accent-yellow text-black border-white shadow-[4px_4px_0px_0px_#FFFFFF] hover:bg-white">
                Start a Quiz Now
              </Button>
            </Link>
          ) : (
            <Link to="/register">
              <Button size="lg" className="bg-accent-yellow text-black border-white shadow-[4px_4px_0px_0px_#FFFFFF] hover:bg-white">
                Create Free Account
              </Button>
            </Link>
          )}
        </div>
      </section>
    </Layout>
  );
}

export default Home;

