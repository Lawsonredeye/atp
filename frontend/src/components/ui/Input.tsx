import type { InputHTMLAttributes } from 'react';
import { forwardRef } from 'react';

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, className = '', id, ...props }, ref) => {
    const inputId = id || label?.toLowerCase().replace(/\s+/g, '-');
    return (
      <div className="mb-4">
        {label && (
          <label htmlFor={inputId} className="block font-display font-bold text-sm uppercase mb-2">
            {label}
          </label>
        )}
        <input
          ref={ref}
          id={inputId}
          className={`w-full px-4 py-3 bg-white font-body text-lg focus:outline-none placeholder:text-gray ${error ? 'border-4 border-accent-red shadow-[4px_4px_0px_0px_#FF3366]' : 'border-4 border-black shadow-brutal-sm focus:shadow-brutal'} ${className}`}
          aria-invalid={error ? 'true' : 'false'}
          {...props}
        />
        {error && <p className="mt-2 font-body text-sm text-accent-red font-bold" role="alert">{error}</p>}
      </div>
    );
  }
);

Input.displayName = 'Input';

