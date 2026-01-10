import { useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { Menu, X } from 'lucide-react';
import { useAuth } from '../../context/AuthContext';

export function Header() {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const { state, logout } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  const isActive = (path: string) => location.pathname === path;

  const navLinks = state.isAuthenticated
    ? [
        { path: '/dashboard', label: 'Dashboard' },
        { path: '/quiz', label: 'Quiz' },
        { path: '/leaderboard', label: 'Leaderboard' },
      ]
    : [];

  return (
    <header className="bg-black text-white border-b-4 border-primary py-4 px-6">
      <div className="max-w-7xl mx-auto flex items-center justify-between">
        <Link to="/" className="font-display text-2xl md:text-3xl font-bold uppercase tracking-tight hover:text-primary transition-colors">
          Score<span className="text-primary">That</span>Exam
        </Link>

        {/* Desktop Nav */}
        <nav className="hidden md:flex items-center gap-6">
          {navLinks.map((link) => (
            <Link key={link.path} to={link.path} className={`font-display font-bold uppercase transition-colors ${isActive(link.path) ? 'text-primary' : 'hover:text-primary'}`}>
              {link.label}
            </Link>
          ))}
          {state.isAuthenticated ? (
            <button onClick={handleLogout} className="px-4 py-2 bg-primary text-black font-display font-bold uppercase border-3 border-white shadow-[3px_3px_0px_0px_#FFFFFF] hover:bg-accent-yellow active:shadow-none active:translate-x-0.5 active:translate-y-0.5 transition-all duration-100 cursor-pointer">
              Logout
            </button>
          ) : (
            <div className="flex gap-4">
              <Link to="/login" className="px-4 py-2 bg-transparent text-white font-display font-bold uppercase border-3 border-white hover:bg-white hover:text-black transition-all duration-100">Login</Link>
              <Link to="/register" className="px-4 py-2 bg-primary text-black font-display font-bold uppercase border-3 border-white shadow-[3px_3px_0px_0px_#FFFFFF] hover:bg-accent-yellow transition-all duration-100">Sign Up</Link>
            </div>
          )}
        </nav>

        {/* Mobile Menu Button */}
        <button className="md:hidden p-2 border-3 border-white cursor-pointer" onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)} aria-label={isMobileMenuOpen ? 'Close menu' : 'Open menu'}>
          {isMobileMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
        </button>
      </div>

      {/* Mobile Nav */}
      {isMobileMenuOpen && (
        <nav className="md:hidden mt-4 pt-4 border-t-2 border-white/20">
          <div className="flex flex-col gap-4">
            {navLinks.map((link) => (
              <Link key={link.path} to={link.path} className={`font-display font-bold uppercase py-2 ${isActive(link.path) ? 'text-primary' : 'hover:text-primary'}`} onClick={() => setIsMobileMenuOpen(false)}>
                {link.label}
              </Link>
            ))}
            {state.isAuthenticated ? (
              <button onClick={() => { handleLogout(); setIsMobileMenuOpen(false); }} className="py-2 mt-2 bg-primary text-black font-display font-bold uppercase border-3 border-white cursor-pointer">Logout</button>
            ) : (
              <div className="flex flex-col gap-4 mt-2">
                <Link to="/login" className="py-2 text-center bg-transparent text-white font-display font-bold uppercase border-3 border-white" onClick={() => setIsMobileMenuOpen(false)}>Login</Link>
                <Link to="/register" className="py-2 text-center bg-primary text-black font-display font-bold uppercase border-3 border-white" onClick={() => setIsMobileMenuOpen(false)}>Sign Up</Link>
              </div>
            )}
          </div>
        </nav>
      )}
    </header>
  );
}
