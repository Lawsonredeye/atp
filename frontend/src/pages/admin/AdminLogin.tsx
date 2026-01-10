import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { Layout } from '../../components/layout';
import { Button, Input, Card, Alert } from '../../components/ui';
import { adminService } from '../../services';
import { Shield, Mail, Lock, ArrowRight } from 'lucide-react';
import type { AxiosError } from 'axios';

const adminLoginSchema = z.object({
  email: z.string().email('Please enter a valid email address'),
  password: z.string().min(1, 'Password is required'),
});

type AdminLoginFormData = z.infer<typeof adminLoginSchema>;

function AdminLogin() {
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<AdminLoginFormData>({
    resolver: zodResolver(adminLoginSchema),
  });

  const onSubmit = async (data: AdminLoginFormData) => {
    try {
      setError(null);
      await adminService.login(data);
      navigate('/admin/dashboard', { replace: true });
    } catch (err) {
      const axiosError = err as AxiosError<{ message: string }>;
      setError(axiosError.response?.data?.message || 'Invalid credentials. Please try again.');
    }
  };

  return (
    <Layout showFooter={false}>
      <div className="min-h-[calc(100vh-80px)] flex items-center justify-center px-4 py-12 sm:py-16">
        <div className="w-full max-w-md">
          {/* Header */}
          <div className="text-center mb-8">
            <div className="inline-flex items-center justify-center w-16 h-16 bg-secondary border-4 border-black shadow-brutal mb-4">
              <Shield className="w-8 h-8 text-white" />
            </div>
            <h1 className="font-display text-3xl sm:text-4xl font-bold mb-2">Admin Login</h1>
            <p className="font-body text-gray-600">
              Access the admin dashboard
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
                  placeholder="admin@example.com"
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
                className="w-full bg-secondary hover:bg-secondary-light"
                size="lg"
              >
                {isSubmitting ? 'Signing in...' : 'Admin Sign In'}
                {!isSubmitting && <ArrowRight className="inline ml-2 w-5 h-5" />}
              </Button>
            </form>

            <div className="mt-6 pt-6 border-t-3 border-black text-center">
              <p className="font-body text-gray-600">
                Not an admin?{' '}
                <Link
                  to="/login"
                  className="font-display font-bold text-primary hover:text-primary-dark underline underline-offset-4"
                >
                  User login
                </Link>
              </p>
            </div>
          </Card>
        </div>
      </div>
    </Layout>
  );
}

export default AdminLogin;

