import type { ReactNode } from 'react';
import { CheckCircle, XCircle, AlertTriangle, Info } from 'lucide-react';

type AlertVariant = 'success' | 'error' | 'warning' | 'info';

interface AlertProps {
  variant: AlertVariant;
  children: ReactNode;
  className?: string;
}

const variantConfig: Record<AlertVariant, { bg: string; Icon: typeof CheckCircle }> = {
  success: { bg: 'bg-accent-green text-black', Icon: CheckCircle },
  error: { bg: 'bg-accent-red text-white', Icon: XCircle },
  warning: { bg: 'bg-accent-yellow text-black', Icon: AlertTriangle },
  info: { bg: 'bg-secondary text-white', Icon: Info },
};

export function Alert({ variant, children, className = '' }: AlertProps) {
  const { bg, Icon } = variantConfig[variant];
  return (
    <div className={`${bg} border-4 border-black shadow-brutal p-4 flex items-center gap-3 ${className}`} role="alert">
      <Icon className="w-6 h-6 flex-shrink-0" />
      <span className="font-body font-bold">{children}</span>
    </div>
  );
}

