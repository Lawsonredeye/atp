import type { ReactNode } from 'react';

type CardVariant = 'default' | 'yellow' | 'green' | 'red' | 'purple' | 'blue';

interface CardProps {
  variant?: CardVariant;
  className?: string;
  children: ReactNode;
  hoverable?: boolean;
  onClick?: () => void;
}

const variantClasses: Record<CardVariant, string> = {
  default: 'bg-white',
  yellow: 'bg-accent-yellow text-black',
  green: 'bg-accent-green text-black',
  red: 'bg-accent-red text-white',
  purple: 'bg-accent-purple text-white',
  blue: 'bg-secondary text-white',
};

export function Card({ variant = 'default', className = '', children, hoverable = false, onClick }: CardProps) {
  const Component = onClick ? 'button' : 'div';

  // Check if className contains a background color class - if so, skip the variant background
  const hasCustomBg = className.includes('bg-');
  const variantClass = hasCustomBg ? '' : variantClasses[variant];

  return (
    <Component
      className={`border-4 border-black shadow-brutal p-6 ${variantClass} ${hoverable ? 'hover:shadow-brutal-lg hover:-translate-y-1 transition-all duration-100 cursor-pointer' : ''} ${onClick ? 'text-left w-full' : ''} ${className}`}
      onClick={onClick}
    >
      {children}
    </Component>
  );
}

