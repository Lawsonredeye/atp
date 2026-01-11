import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Layout } from '../components/layout';
import { Button, Input, Card, Alert } from '../components/ui';
import { useAuth } from '../context/AuthContext';
import { registerSchema, type RegisterFormData } from '../utils/validators';
import { UserPlus, Mail, Lock, User, ArrowRight, CheckCircle } from 'lucide-react';
import type { AxiosError } from 'axios';

const benefits = [
  'Access to 10,000+ JAMB past questions',
  'Track your progress and improvement',
  'Compete on the national leaderboard',
  'Completely free to use',
];

function Register() {
  const navigate = useNavigate();
  const { register: registerUser } = useAuth();
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = async (data: RegisterFormData) => {
    try {
      setError(null);
      const { confirmPassword: _, ...registerData } = data;
      await registerUser(registerData);
      navigate('/dashboard', { replace: true });
    } catch (err) {
      const axiosError = err as AxiosError<{ message: string }>;
      setError(axiosError.response?.data?.message || 'Registration failed. Please try again.');
    }
  };

  return (
    <Layout showFooter={false}>
      <div className="min-h-[calc(100vh-80px)] flex items-center justify-center px-4 py-12 sm:py-16">
        <div className="w-full max-w-5xl">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12 items-start">
            {/* Left Side - Benefits */}
            <div className="hidden lg:block">
              <span className="inline-block px-4 py-2 bg-accent-yellow border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-6">
                Join 50,000+ Students
              </span>
              <h2 className="font-display text-4xl font-bold mb-6">
                Start Your <span className="text-primary">JAMB</span> Prep Journey Today
              </h2>
              <p className="font-body text-lg text-gray-700 mb-8">
                Create your free account and get instant access to everything you need to ace your JAMB exams.
              </p>
              <ul className="space-y-4">
                {benefits.map((benefit, index) => (
                  <li key={index} className="flex items-center gap-3">
                    <div className="w-8 h-8 bg-accent-green border-3 border-black flex items-center justify-center flex-shrink-0">
                      <CheckCircle className="w-5 h-5" />
                    </div>
                    <span className="font-body text-lg">{benefit}</span>
                  </li>
                ))}
              </ul>

              {/* Decorative Card */}
              <Card className="mt-8 p-6 bg-primary">
                <p className="font-display font-bold text-xl mb-2">
                  "ScoreThatExam helped me improve my score by 40%!"
                </p>
                <p className="font-body">— Chioma A., UNILAG Aspirant</p>
              </Card>
            </div>

            {/* Right Side - Form */}
            <div>
              {/* Header - Mobile Only */}
              <div className="text-center mb-8 lg:hidden">
                <div className="inline-flex items-center justify-center w-16 h-16 bg-primary border-4 border-black shadow-brutal mb-4">
                  <UserPlus className="w-8 h-8" />
                </div>
                <h1 className="font-display text-3xl font-bold mb-2">Create Account</h1>
                <p className="font-body text-gray-600">
                  Start your JAMB prep journey for free
                </p>
              </div>

              {/* Registration Form Card */}
              <Card className="p-6 sm:p-8">
                <div className="hidden lg:block mb-6">
                  <div className="flex items-center gap-4 mb-4">
                    <div className="w-12 h-12 bg-primary border-3 border-black shadow-brutal-sm flex items-center justify-center">
                      <UserPlus className="w-6 h-6" />
                    </div>
                    <div>
                      <h1 className="font-display text-2xl font-bold">Create Account</h1>
                      <p className="font-body text-gray-600 text-sm">It's free and takes less than a minute</p>
                    </div>
                  </div>
                </div>

                {error && (
                  <Alert variant="error" className="mb-6">
                    {error}
                  </Alert>
                )}

                <form onSubmit={handleSubmit(onSubmit)} noValidate>
                  <div className="relative mb-4">
                    <User className="absolute left-4 top-[42px] w-5 h-5 text-gray-500" />
                    <Input
                      label="Full Name"
                      type="text"
                      placeholder="John Doe"
                      className="pl-12"
                      error={errors.full_name?.message}
                      {...register('full_name')}
                    />
                  </div>

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

                  <div className="relative mb-4">
                    <Lock className="absolute left-4 top-[42px] w-5 h-5 text-gray-500" />
                    <Input
                      label="Password"
                      type="password"
                      placeholder="Min. 6 characters"
                      className="pl-12"
                      error={errors.password?.message}
                      {...register('password')}
                    />
                  </div>

                  <div className="relative mb-6">
                    <Lock className="absolute left-4 top-[42px] w-5 h-5 text-gray-500" />
                    <Input
                      label="Confirm Password"
                      type="password"
                      placeholder="Confirm your password"
                      className="pl-12"
                      error={errors.confirmPassword?.message}
                      {...register('confirmPassword')}
                    />
                  </div>

                  <Button
                    type="submit"
                    isLoading={isSubmitting}
                    className="w-full"
                    size="lg"
                  >
                    {isSubmitting ? 'Creating Account...' : 'Create Free Account'}
                    {!isSubmitting && <ArrowRight className="inline ml-2 w-5 h-5" />}
                  </Button>
                </form>

                <div className="mt-6 pt-6 border-t-3 border-black text-center">
                  <p className="font-body text-gray-600">
                    Already have an account?{' '}
                    <Link
                      to="/login"
                      className="font-display font-bold text-primary hover:text-primary-dark underline underline-offset-4"
                    >
                      Sign in
                    </Link>
                  </p>
                </div>
              </Card>

              <p className="mt-4 text-center font-body text-sm text-gray-500">
                By creating an account, you agree to our{' '}
                <Link to="/terms" className="underline hover:text-primary">Terms of Service</Link>
                {' '}and{' '}
                <Link to="/privacy" className="underline hover:text-primary">Privacy Policy</Link>
              </p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Register;

