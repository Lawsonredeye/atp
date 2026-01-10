import type { SelectHTMLAttributes } from 'react';
import { forwardRef } from 'react';

interface SelectOption {
  value: string | number;
  label: string;
}

interface SelectProps extends SelectHTMLAttributes<HTMLSelectElement> {
  label?: string;
  error?: string;
  options: SelectOption[];
  placeholder?: string;
}

export const Select = forwardRef<HTMLSelectElement, SelectProps>(
  ({ label, error, options, placeholder = 'Select an option', className = '', id, ...props }, ref) => {
    const selectId = id || label?.toLowerCase().replace(/\s+/g, '-');
    return (
      <div className="mb-4">
        {label && (
          <label htmlFor={selectId} className="block font-display font-bold text-sm uppercase mb-2">
            {label}
          </label>
        )}
        <select
          ref={ref}
          id={selectId}
          className={`w-full px-4 py-3 bg-white font-body text-lg appearance-none cursor-pointer focus:outline-none ${error ? 'border-4 border-accent-red' : 'border-4 border-black shadow-brutal-sm focus:shadow-brutal'} ${className}`}
          {...props}
        >
          <option value="">{placeholder}</option>
          {options.map((option) => (
            <option key={option.value} value={option.value}>{option.label}</option>
          ))}
        </select>
        {error && <p className="mt-2 font-body text-sm text-accent-red font-bold" role="alert">{error}</p>}
      </div>
    );
  }
);

Select.displayName = 'Select';

