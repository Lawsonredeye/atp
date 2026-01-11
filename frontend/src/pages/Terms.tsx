import { Layout } from '../components/layout';
import { Card } from '../components/ui';
import { FileText, CheckCircle, XCircle, AlertTriangle, Scale, Mail } from 'lucide-react';

const sections = [
  {
    title: 'Acceptance of Terms',
    content: `By accessing or using AceThatPaper, you agree to be bound by these Terms of Service and all applicable laws and regulations. If you do not agree with any of these terms, you are prohibited from using or accessing this site.

The materials contained in this website are protected by applicable copyright and trademark law. These terms of service may be updated from time to time, and your continued use of the platform constitutes acceptance of any changes.`,
  },
  {
    title: 'User Accounts',
    content: `To access certain features of AceThatPaper, you must register for an account. When you register, you agree to:

• Provide accurate, current, and complete information during the registration process
• Maintain the security of your password and accept all risks of unauthorized access to your account
• Promptly notify us if you discover or suspect any security breaches related to your account
• Take responsibility for all activities that occur under your account

We reserve the right to suspend or terminate accounts that violate these terms or engage in suspicious activity.`,
  },
  {
    title: 'Acceptable Use',
    content: `You agree to use AceThatPaper only for lawful purposes and in accordance with these Terms. You agree NOT to:

• Use the service for any illegal or unauthorized purpose
• Attempt to gain unauthorized access to any portion of the platform
• Use any automated system, including "robots" or "spiders," to access the service
• Share your account credentials with others
• Copy, redistribute, or sell any content from our platform
• Harass, abuse, or harm other users
• Submit false information or impersonate others
• Attempt to manipulate leaderboard rankings through fraudulent means`,
  },
  {
    title: 'Intellectual Property',
    content: `All content on AceThatPaper, including but not limited to questions, explanations, graphics, logos, and software, is the property of AceThatPaper or its content suppliers and is protected by Nigerian and international copyright laws.

You may not reproduce, distribute, modify, create derivative works of, publicly display, publicly perform, republish, download, store, or transmit any of the material on our website without prior written consent.

The questions provided are based on JAMB syllabus and past examinations. While we strive for accuracy, we do not guarantee that our questions are identical to actual JAMB examinations.`,
  },
  {
    title: 'User-Generated Content',
    content: `Any content you submit to AceThatPaper, including feedback, suggestions, or quiz performance data, may be used by us to improve our services.

You retain ownership of any intellectual property rights that you hold in the content you submit. However, by submitting content, you grant us a worldwide, royalty-free license to use, store, and process this content for the purpose of providing and improving our services.`,
  },
  {
    title: 'Disclaimer of Warranties',
    content: `AceThatPaper is provided on an "AS IS" and "AS AVAILABLE" basis. We make no warranties, expressed or implied, regarding:

• The accuracy or completeness of any content
• The availability or reliability of the service
• That the service will meet your specific requirements
• That any errors or defects will be corrected

While we strive to maintain accurate and up-to-date content, we cannot guarantee that our practice questions will match actual JAMB examination questions. Success on our platform does not guarantee success on the actual JAMB examination.`,
  },
  {
    title: 'Limitation of Liability',
    content: `To the fullest extent permitted by law, AceThatPaper shall not be liable for any indirect, incidental, special, consequential, or punitive damages, including but not limited to:

• Loss of profits or revenue
• Loss of data or information
• Loss of goodwill
• Any other intangible losses

resulting from your use or inability to use the service, even if we have been advised of the possibility of such damages.`,
  },
  {
    title: 'Termination',
    content: `We may terminate or suspend your account and access to AceThatPaper immediately, without prior notice or liability, for any reason, including:

• Breach of these Terms of Service
• Suspected fraudulent, abusive, or illegal activity
• Upon your request to delete your account

Upon termination, your right to use the service will immediately cease. All provisions of the Terms which should survive termination shall survive, including ownership provisions, warranty disclaimers, and limitations of liability.`,
  },
  {
    title: 'Governing Law',
    content: `These Terms shall be governed by and construed in accordance with the laws of the Federal Republic of Nigeria, without regard to its conflict of law provisions.

Any disputes arising from these terms or your use of AceThatPaper shall be resolved through the courts of Nigeria. You agree to submit to the personal and exclusive jurisdiction of the courts located in Lagos, Nigeria.`,
  },
];

function Terms() {
  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-black text-white py-12 sm:py-16 lg:py-20 px-4 sm:px-6">
        <div className="max-w-4xl mx-auto text-center">
          <div className="inline-flex items-center justify-center w-20 h-20 bg-primary border-4 border-white shadow-brutal mb-6">
            <FileText className="w-10 h-10 text-black" />
          </div>
          <h1 className="font-display text-4xl sm:text-5xl font-bold mb-6">
            Terms of Service
          </h1>
          <p className="font-body text-lg sm:text-xl text-white/90">
            Please read these terms carefully before using AceThatPaper.
          </p>
          <p className="font-body text-sm text-white/70 mt-4">
            Effective Date: January 11, 2026
          </p>
        </div>
      </section>

      {/* Quick Summary */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8">
            <h2 className="font-display text-2xl font-bold mb-6">Quick Summary</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div className="flex items-start gap-3">
                <div className="w-10 h-10 bg-accent-green border-3 border-black flex items-center justify-center flex-shrink-0">
                  <CheckCircle className="w-5 h-5" />
                </div>
                <div>
                  <h3 className="font-display font-bold mb-1">You CAN</h3>
                  <ul className="font-body text-gray-700 text-sm space-y-1">
                    <li>• Use our platform for personal study</li>
                    <li>• Track your progress and scores</li>
                    <li>• Compete on leaderboards</li>
                    <li>• Access all free features</li>
                  </ul>
                </div>
              </div>
              <div className="flex items-start gap-3">
                <div className="w-10 h-10 bg-accent-red border-3 border-black flex items-center justify-center flex-shrink-0">
                  <XCircle className="w-5 h-5 text-white" />
                </div>
                <div>
                  <h3 className="font-display font-bold mb-1">You CANNOT</h3>
                  <ul className="font-body text-gray-700 text-sm space-y-1">
                    <li>• Share your account with others</li>
                    <li>• Copy or redistribute our content</li>
                    <li>• Use bots or automated tools</li>
                    <li>• Cheat or manipulate rankings</li>
                  </ul>
                </div>
              </div>
            </div>
          </Card>
        </div>
      </section>

      {/* Important Notice */}
      <section className="py-8 px-4 sm:px-6 bg-accent-yellow">
        <div className="max-w-4xl mx-auto">
          <div className="flex items-center gap-4">
            <AlertTriangle className="w-8 h-8 flex-shrink-0" />
            <p className="font-body font-bold">
              By creating an account or using AceThatPaper, you acknowledge that you have read,
              understood, and agree to be bound by these Terms of Service.
            </p>
          </div>
        </div>
      </section>

      {/* Terms Sections */}
      <section className="py-12 px-4 sm:px-6 bg-white">
        <div className="max-w-4xl mx-auto space-y-8">
          {sections.map((section, index) => (
            <Card key={index} className="p-6 sm:p-8">
              <div className="flex items-center gap-4 mb-4">
                <span className="inline-flex items-center justify-center w-10 h-10 bg-secondary text-white font-display font-bold border-3 border-black">
                  {index + 1}
                </span>
                <h2 className="font-display text-xl sm:text-2xl font-bold">{section.title}</h2>
              </div>
              <div className="font-body text-gray-700 whitespace-pre-line">
                {section.content}
              </div>
            </Card>
          ))}
        </div>
      </section>

      {/* Contact Section */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-4xl mx-auto">
          <Card className="p-6 sm:p-8 bg-secondary text-white">
            <div className="flex flex-col sm:flex-row items-center gap-6">
              <div className="w-16 h-16 bg-primary border-4 border-black flex items-center justify-center flex-shrink-0">
                <Scale className="w-8 h-8 text-black" />
              </div>
              <div className="text-center sm:text-left flex-1">
                <h2 className="font-display text-2xl font-bold mb-2">Questions About These Terms?</h2>
                <p className="font-body text-white/90">
                  If you have any questions about these Terms of Service, please contact our
                  support team.
                </p>
              </div>
              <a
                href="/contact"
                className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-black font-display font-bold uppercase border-4 border-black shadow-brutal hover:bg-accent-yellow active:shadow-none active:translate-x-1 active:translate-y-1 transition-all duration-100"
              >
                <Mail className="w-5 h-5" />
                Contact Us
              </a>
            </div>
          </Card>
        </div>
      </section>
    </Layout>
  );
}

export default Terms;

