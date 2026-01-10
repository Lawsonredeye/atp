interface LoadingSpinnerProps {
  size?: 'sm' | 'md' | 'lg';
  className?: string;
}

const sizeClasses = { sm: 'w-6 h-6', md: 'w-12 h-12', lg: 'w-16 h-16' };

export function LoadingSpinner({ size = 'md', className = '' }: LoadingSpinnerProps) {
  return (
    <div className={`flex items-center justify-center p-8 ${className}`}>
      <div className={`${sizeClasses[size]} border-4 border-black border-t-primary rounded-full animate-spin`} role="status" aria-label="Loading" />
    </div>
  );
}

export function LoadingScreen() {
  return (
    <div className="min-h-screen bg-cream flex items-center justify-center">
      <div className="text-center">
        <LoadingSpinner size="lg" />
        <p className="mt-4 font-display font-bold text-xl uppercase">Loading...</p>
      </div>
    </div>
  );
}

