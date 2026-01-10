import { Link } from 'react-router-dom';

export function Footer() {
  return (
    <footer className="bg-secondary text-white border-t-4 border-black py-8 px-6">
      <div className="max-w-7xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <h3 className="font-display text-2xl font-bold uppercase mb-4">
              Score<span className="text-primary">That</span>Exam
            </h3>
            <p className="font-body text-white/80">Ace your JAMB exams with confidence.</p>
          </div>
          <div>
            <h4 className="font-display font-bold uppercase mb-4 text-primary">Quick Links</h4>
            <ul className="space-y-2 font-body">
              <li><Link to="/dashboard" className="hover:text-primary transition-colors">Dashboard</Link></li>
              <li><Link to="/quiz" className="hover:text-primary transition-colors">Start Quiz</Link></li>
              <li><Link to="/leaderboard" className="hover:text-primary transition-colors">Leaderboard</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-display font-bold uppercase mb-4 text-primary">Subjects</h4>
            <ul className="space-y-2 font-body">
              <li><Link to="/quiz?subject=1" className="hover:text-primary transition-colors">Mathematics</Link></li>
              <li><Link to="/quiz?subject=2" className="hover:text-primary transition-colors">English</Link></li>
              <li><Link to="/quiz?subject=3" className="hover:text-primary transition-colors">Physics</Link></li>
            </ul>
          </div>
        </div>
        <div className="mt-8 pt-6 border-t-2 border-white/20 text-center font-display font-bold uppercase text-sm">
          © {new Date().getFullYear()} ScoreThatExam. All rights reserved.
        </div>
      </div>
    </footer>
  );
}

