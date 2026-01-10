import type { ReactNode } from 'react';

interface BadgeProps {
  children: ReactNode;
  variant?: 'primary' | 'secondary' | 'success' | 'error' | 'warning';
  className?: string;
}

const variantClasses = {
  primary: 'bg-primary text-black',
  secondary: 'bg-secondary text-white',
  success: 'bg-accent-green text-black',
  error: 'bg-accent-red text-white',
  warning: 'bg-accent-yellow text-black',
};

export function Badge({ children, variant = 'primary', className = '' }: BadgeProps) {
  return (
    <span className={`inline-block px-3 py-1 font-display font-bold text-sm uppercase border-3 border-black ${variantClasses[variant]} ${className}`}>
      {children}
    </span>
  );
}

