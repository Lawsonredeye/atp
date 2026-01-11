import { useState } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Layout } from '../components/layout';
import { Button, Input, Card, Alert } from '../components/ui';
import { useAuth } from '../context/AuthContext';
import { loginSchema, type LoginFormData } from '../utils/validators';
import { LogIn, Mail, Lock, ArrowRight } from 'lucide-react';
import type { AxiosError } from 'axios';

function Login() {
  const navigate = useNavigate();
  const location = useLocation();
  const { login } = useAuth();
  const [error, setError] = useState<string | null>(null);

  const from = (location.state as { from?: { pathname: string } })?.from?.pathname || '/dashboard';

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginFormData) => {
    try {
      setError(null);
      await login(data);
      navigate(from, { replace: true });
    } catch (err) {
      const axiosError = err as AxiosError<{ message: string }>;
      setError(axiosError.response?.data?.message || 'Invalid email or password. Please try again.');
    }
  };

  return (
    <Layout showFooter={false}>
      <div className="min-h-[calc(100vh-80px)] flex items-center justify-center px-4 py-12 sm:py-16">
        <div className="w-full max-w-md">
          {/* Header */}
          <div className="text-center mb-8">
            <div className="inline-flex items-center justify-center w-16 h-16 bg-primary border-4 border-black shadow-brutal mb-4">
              <LogIn className="w-8 h-8" />
            </div>
            <h1 className="font-display text-3xl sm:text-4xl font-bold mb-2">Welcome Back!</h1>
            <p className="font-body text-gray-600">
              Sign in to continue your JAMB prep journey
            </p>
          </div>

          {/* Login Form Card */}
          <Card className="p-6 sm:p-8">
            {error && (
              <Alert variant="error" className="mb-6">
                {error}
              </Alert>
            )}

            <form onSubmit={handleSubmit(onSubmit)} noValidate>
              <div className="relative mb-4">
                <Mail className="absolute left-4 top-[42px] w-5 h-5 text-gray-500" />
                <Input
                  label="Email Address"
                  type="email"
                  placeholder="you@example.com"
                  className="pl-12"
                  error={errors.email?.message}
                  {...register('email')}
                />
              </div>

              <div className="relative mb-6">
                <Lock className="absolute left-4 top-[42px] w-5 h-5 text-gray-500" />
                <Input
                  label="Password"
                  type="password"
                  placeholder="Enter your password"
                  className="pl-12"
                  error={errors.password?.message}
                  {...register('password')}
                />
              </div>

              <Button
                type="submit"
                isLoading={isSubmitting}
                className="w-full"
                size="lg"
              >
                {isSubmitting ? 'Signing in...' : 'Sign In'}
                {!isSubmitting && <ArrowRight className="inline ml-2 w-5 h-5" />}
              </Button>
            </form>

            <div className="mt-6 pt-6 border-t-3 border-black text-center">
              <p className="font-body text-gray-600">
                Don't have an account?{' '}
                <Link
                  to="/register"
                  className="font-display font-bold text-primary hover:text-primary-dark underline underline-offset-4"
                >
                  Sign up for free
                </Link>
              </p>
            </div>
          </Card>

          {/* Admin Link */}
          <div className="mt-4 text-center">
            <Link
              to="/admin/login"
              className="font-body text-sm text-gray-500 hover:text-secondary underline"
            >
              Admin Login →
            </Link>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Login;

