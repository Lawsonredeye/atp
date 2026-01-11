import { Layout } from '../components/layout';
import { Card } from '../components/ui';
import { Shield, Lock, Eye, Database, UserCheck, Mail } from 'lucide-react';

const sections = [
  {
    icon: Database,
    title: 'Information We Collect',
    content: [
      {
        subtitle: 'Account Information',
        text: 'When you create an account, we collect your name, email address, and password (encrypted). This information is necessary to provide you with access to our services.',
      },
      {
        subtitle: 'Quiz Performance Data',
        text: 'We collect data about your quiz attempts, scores, and progress. This helps us provide personalized recommendations and track your improvement over time.',
      },
      {
        subtitle: 'Usage Information',
        text: 'We automatically collect information about how you use our platform, including pages visited, time spent on quizzes, and device information.',
      },
    ],
  },
  {
    icon: Eye,
    title: 'How We Use Your Information',
    content: [
      {
        subtitle: 'Providing Services',
        text: 'We use your information to create and manage your account, deliver quizzes, track your progress, and display your ranking on leaderboards.',
      },
      {
        subtitle: 'Improving Our Platform',
        text: 'We analyze usage patterns to improve our question quality, user interface, and overall learning experience.',
      },
      {
        subtitle: 'Communication',
        text: 'We may send you service-related emails, such as password resets, account updates, or important announcements about our platform.',
      },
    ],
  },
  {
    icon: Lock,
    title: 'Data Security',
    content: [
      {
        subtitle: 'Encryption',
        text: 'All passwords are encrypted using industry-standard hashing algorithms. We never store plain-text passwords.',
      },
      {
        subtitle: 'Secure Transmission',
        text: 'All data transmitted between your browser and our servers is encrypted using HTTPS/TLS protocols.',
      },
      {
        subtitle: 'Access Controls',
        text: 'We implement strict access controls to ensure only authorized personnel can access user data, and only when necessary.',
      },
    ],
  },
  {
    icon: UserCheck,
    title: 'Your Rights',
    content: [
      {
        subtitle: 'Access Your Data',
        text: 'You have the right to request a copy of the personal data we hold about you at any time.',
      },
      {
        subtitle: 'Update Your Information',
        text: 'You can update your account information at any time through your profile settings.',
      },
      {
        subtitle: 'Delete Your Account',
        text: 'You can request deletion of your account and associated data by contacting our support team.',
      },
    ],
  },
];

function Privacy() {
  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-secondary text-white py-12 sm:py-16 lg:py-20 px-4 sm:px-6">
        <div className="max-w-4xl mx-auto text-center">
          <div className="inline-flex items-center justify-center w-20 h-20 bg-white border-4 border-black shadow-brutal mb-6">
            <Shield className="w-10 h-10 text-secondary" />
          </div>
          <h1 className="font-display text-4xl sm:text-5xl font-bold mb-6">
            Privacy Policy
          </h1>
          <p className="font-body text-lg sm:text-xl text-white/90">
            Your privacy matters to us. This policy explains how we collect, use,
            and protect your personal information.
          </p>
          <p className="font-body text-sm text-white/70 mt-4">
            Last updated: January 11, 2026
          </p>
        </div>
      </section>

      {/* Introduction */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8 bg-white">
            <h2 className="font-display text-2xl font-bold mb-4">Introduction</h2>
            <p className="font-body text-gray-700 mb-4">
              ScoreThatExam ("we," "our," or "us") is committed to protecting your privacy.
              This Privacy Policy explains how we collect, use, disclose, and safeguard your
              information when you use our website and services.
            </p>
            <p className="font-body text-gray-700">
              By using ScoreThatExam, you agree to the collection and use of information in
              accordance with this policy. If you do not agree with our policies and practices,
              please do not use our services.
            </p>
          </Card>
        </div>
      </section>

      {/* Policy Sections */}
      <section className="py-12 px-4 sm:px-6 bg-white">
        <div className="max-w-4xl mx-auto space-y-8">
          {sections.map((section, index) => (
            <Card key={index} className="p-6 sm:p-8">
              <div className="flex items-center gap-4 mb-6">
                <div className="w-14 h-14 bg-primary border-4 border-black shadow-brutal-sm flex items-center justify-center flex-shrink-0">
                  <section.icon className="w-7 h-7" />
                </div>
                <h2 className="font-display text-2xl font-bold">{section.title}</h2>
              </div>
              <div className="space-y-6">
                {section.content.map((item, itemIndex) => (
                  <div key={itemIndex}>
                    <h3 className="font-display font-bold text-lg mb-2">{item.subtitle}</h3>
                    <p className="font-body text-gray-700">{item.text}</p>
                  </div>
                ))}
              </div>
            </Card>
          ))}
        </div>
      </section>

      {/* Cookies Section */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8">
            <h2 className="font-display text-2xl font-bold mb-4">Cookies</h2>
            <p className="font-body text-gray-700 mb-4">
              We use cookies and similar tracking technologies to track activity on our
              platform and store certain information. Cookies are files with a small amount
              of data that may include an anonymous unique identifier.
            </p>
            <p className="font-body text-gray-700 mb-4">
              We use the following types of cookies:
            </p>
            <ul className="list-disc list-inside font-body text-gray-700 space-y-2 ml-4">
              <li><strong>Essential Cookies:</strong> Required for the platform to function properly</li>
              <li><strong>Authentication Cookies:</strong> Used to keep you logged in</li>
              <li><strong>Preference Cookies:</strong> Remember your settings and preferences</li>
              <li><strong>Analytics Cookies:</strong> Help us understand how users interact with our platform</li>
            </ul>
          </Card>
        </div>
      </section>

      {/* Third-Party Services */}
      <section className="py-12 px-4 sm:px-6 bg-white">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8">
            <h2 className="font-display text-2xl font-bold mb-4">Third-Party Services</h2>
            <p className="font-body text-gray-700 mb-4">
              We may employ third-party companies and individuals to facilitate our service,
              provide the service on our behalf, perform service-related tasks, or assist us
              in analyzing how our service is used.
            </p>
            <p className="font-body text-gray-700">
              These third parties have access to your personal information only to perform
              these tasks on our behalf and are obligated not to disclose or use it for any
              other purpose.
            </p>
          </Card>
        </div>
      </section>

      {/* Children's Privacy */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8 bg-accent-yellow">
            <h2 className="font-display text-2xl font-bold mb-4">Children's Privacy</h2>
            <p className="font-body text-gray-800 mb-4">
              Our service is intended for users who are at least 13 years old. We do not
              knowingly collect personally identifiable information from children under 13.
              If you are a parent or guardian and you are aware that your child has provided
              us with personal information, please contact us.
            </p>
            <p className="font-body text-gray-800">
              If we become aware that we have collected personal information from children
              under 13 without verification of parental consent, we take steps to remove
              that information from our servers.
            </p>
          </Card>
        </div>
      </section>

      {/* Contact Section */}
      <section className="py-12 px-4 sm:px-6 bg-white">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8 bg-secondary text-white">
            <div className="flex items-center gap-4 mb-4">
              <div className="w-14 h-14 bg-primary border-4 border-black flex items-center justify-center">
                <Mail className="w-7 h-7 text-black" />
              </div>
              <div>
                <h2 className="font-display text-2xl font-bold">Questions About This Policy?</h2>
              </div>
            </div>
            <p className="font-body text-white/90 mb-4">
              If you have any questions or concerns about this Privacy Policy, please don't
              hesitate to contact us.
            </p>
            <a
              href="/contact"
              className="inline-block px-6 py-3 bg-primary text-black font-display font-bold uppercase border-4 border-black shadow-brutal hover:bg-accent-yellow active:shadow-none active:translate-x-1 active:translate-y-1 transition-all duration-100"
            >
              Contact Us
            </a>
          </Card>
        </div>
      </section>

      {/* Policy Changes */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8">
            <h2 className="font-display text-2xl font-bold mb-4">Changes to This Policy</h2>
            <p className="font-body text-gray-700 mb-4">
              We may update our Privacy Policy from time to time. We will notify you of any
              changes by posting the new Privacy Policy on this page and updating the
              "Last updated" date.
            </p>
            <p className="font-body text-gray-700">
              You are advised to review this Privacy Policy periodically for any changes.
              Changes to this Privacy Policy are effective when they are posted on this page.
            </p>
          </Card>
        </div>
      </section>
    </Layout>
  );
}

export default Privacy;

