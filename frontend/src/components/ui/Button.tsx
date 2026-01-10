import type { ButtonHTMLAttributes, ReactNode } from 'react';

type ButtonVariant = 'primary' | 'secondary' | 'outline' | 'danger';
type ButtonSize = 'sm' | 'md' | 'lg';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  isLoading?: boolean;
  children: ReactNode;
}

const variantClasses: Record<ButtonVariant, string> = {
  primary: 'bg-primary text-black hover:bg-accent-yellow hover:shadow-brutal-lg',
  secondary: 'bg-secondary text-white hover:bg-secondary-light',
  outline: 'bg-white text-black hover:bg-black hover:text-white',
  danger: 'bg-accent-red text-white hover:bg-red-700',
};

const sizeClasses: Record<ButtonSize, string> = {
  sm: 'px-4 py-2 text-sm',
  md: 'px-6 py-3 text-lg',
  lg: 'px-8 py-4 text-xl',
};

export function Button({ variant = 'primary', size = 'md', isLoading = false, disabled, children, className = '', ...props }: ButtonProps) {
  const isDisabled = disabled || isLoading;
  return (
    <button
      className={`font-display font-bold uppercase border-4 border-black shadow-brutal active:shadow-none active:translate-x-1 active:translate-y-1 transition-all duration-100 focus:outline-none focus:ring-4 focus:ring-black focus:ring-offset-2 ${variantClasses[variant]} ${sizeClasses[size]} ${isDisabled ? 'opacity-50 cursor-not-allowed shadow-none' : 'cursor-pointer'} ${className}`}
      disabled={isDisabled}
      {...props}
    >
      {isLoading ? (
        <span className="flex items-center justify-center gap-2">
          <span className="w-5 h-5 border-3 border-current border-t-transparent rounded-full animate-spin" />
          Loading...
        </span>
      ) : children}
    </button>
  );
}
