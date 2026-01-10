import { format } from 'date-fns';

export function formatDate(date: string | Date): string {
  return format(new Date(date), 'MMM dd, yyyy');
}

export function formatScore(score: number | undefined | null): string {
  if (score === undefined || score === null) return '0';
  return score.toLocaleString();
}

export function formatPercentage(value: number | undefined | null, decimals = 0): string {
  if (value === undefined || value === null) return '0%';
  return `${value.toFixed(decimals)}%`;
}

export function formatAccuracy(correct: number, total: number): string {
  if (total === 0) return '0%';
  return formatPercentage((correct / total) * 100, 1);
}
