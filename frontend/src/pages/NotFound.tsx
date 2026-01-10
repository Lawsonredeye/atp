import { Link } from 'react-router-dom';
import { Layout } from '../components/layout';
import { Button, Card } from '../components/ui';
import { Home, Search, ArrowLeft } from 'lucide-react';

function NotFound() {
  return (
    <Layout showFooter={false}>
      <div className="min-h-[calc(100vh-80px)] flex items-center justify-center px-4 py-12">
        <div className="w-full max-w-lg text-center">
          {/* 404 Display */}
          <div className="relative mb-8">
            <div className="font-display text-[150px] sm:text-[200px] font-bold leading-none text-primary opacity-20">
              404
            </div>
            <div className="absolute inset-0 flex items-center justify-center">
              <Card className="bg-white px-8 py-4 shadow-brutal-lg">
                <Search className="w-12 h-12 mx-auto mb-2" />
                <p className="font-display font-bold uppercase">Page Not Found</p>
              </Card>
            </div>
          </div>

          {/* Message */}
          <h1 className="font-display text-3xl sm:text-4xl font-bold mb-4">
            Oops! Wrong Turn
          </h1>
          <p className="font-body text-lg text-gray-600 mb-8">
            The page you're looking for doesn't exist or has been moved.
            Let's get you back on track!
          </p>

          {/* Actions */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link to="/">
              <Button size="lg" className="w-full sm:w-auto">
                <Home className="inline mr-2 w-5 h-5" />
                Go Home
              </Button>
            </Link>
            <Button
              variant="outline"
              size="lg"
              onClick={() => window.history.back()}
              className="w-full sm:w-auto"
            >
              <ArrowLeft className="inline mr-2 w-5 h-5" />
              Go Back
            </Button>
          </div>

          {/* Fun Message */}
          <Card className="mt-12 p-6 bg-accent-yellow">
            <p className="font-display font-bold mb-2">While you're here...</p>
            <p className="font-body">
              Did you know that practicing just 10 questions a day can improve your JAMB score by up to 30%? 📚
            </p>
            <Link to="/dashboard" className="inline-block mt-4">
              <Button variant="secondary" size="sm">
                Start Practicing
              </Button>
            </Link>
          </Card>
        </div>
      </div>
    </Layout>
  );
}

export default NotFound;

