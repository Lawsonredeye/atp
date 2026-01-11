import { useState } from 'react';
import { Layout } from '../components/layout';
import { Card, Button, Input, Alert } from '../components/ui';
import { Mail, Phone, MapPin, Send, MessageSquare, Clock, CheckCircle } from 'lucide-react';

const contactMethods = [
  {
    icon: Mail,
    title: 'Email Us',
    value: 'support@scorethatexam.com',
    description: 'Send us an email anytime',
    color: 'bg-primary',
  },
  {
    icon: Phone,
    title: 'Call Us',
    value: '+234 800 123 4567',
    description: 'Mon-Fri, 9am-5pm WAT',
    color: 'bg-accent-green',
  },
  {
    icon: MapPin,
    title: 'Visit Us',
    value: 'Lagos, Nigeria',
    description: 'Our headquarters',
    color: 'bg-accent-yellow',
  },
];

const faqs = [
  {
    question: 'How do I reset my password?',
    answer: 'Click on "Forgot Password" on the login page and enter your email. You\'ll receive a reset link within minutes.',
  },
  {
    question: 'Are the practice questions free?',
    answer: 'Yes! ScoreThatExam is completely free to use. All subjects and questions are available at no cost.',
  },
  {
    question: 'How accurate are the JAMB questions?',
    answer: 'Our questions are curated from JAMB syllabus and past examinations. While we strive for accuracy, they are practice questions and may not be identical to actual exams.',
  },
  {
    question: 'Can I use ScoreThatExam on mobile?',
    answer: 'Yes! Our platform is fully responsive and works great on all devices - phones, tablets, and computers.',
  },
];

function Contact() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    subject: '',
    message: '',
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSubmitted, setIsSubmitted] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    setError(null);

    // Simulate form submission
    try {
      await new Promise((resolve) => setTimeout(resolve, 1500));
      setIsSubmitted(true);
      setFormData({ name: '', email: '', subject: '', message: '' });
    } catch {
      setError('Failed to send message. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-primary py-12 sm:py-16 lg:py-20 px-4 sm:px-6">
        <div className="max-w-4xl mx-auto text-center">
          <div className="inline-flex items-center justify-center w-20 h-20 bg-black border-4 border-black shadow-brutal mb-6">
            <MessageSquare className="w-10 h-10 text-white" />
          </div>
          <h1 className="font-display text-4xl sm:text-5xl font-bold mb-6">
            Get In Touch
          </h1>
          <p className="font-body text-lg sm:text-xl">
            Have questions, feedback, or need help? We're here for you.
            Reach out and we'll respond as soon as we can.
          </p>
        </div>
      </section>

      {/* Contact Methods */}
      <section className="py-12 px-4 sm:px-6 bg-cream">
        <div className="max-w-6xl mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {contactMethods.map((method, index) => (
              <Card key={index} className="p-6 text-center hover:shadow-brutal-lg transition-all duration-100 hover:-translate-y-1">
                <div className={`w-16 h-16 ${method.color} border-4 border-black shadow-brutal-sm mx-auto mb-4 flex items-center justify-center`}>
                  <method.icon className="w-8 h-8" />
                </div>
                <h3 className="font-display text-xl font-bold mb-2">{method.title}</h3>
                <p className="font-display font-bold text-primary mb-1">{method.value}</p>
                <p className="font-body text-gray-600 text-sm">{method.description}</p>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Contact Form & Info */}
      <section className="py-12 sm:py-16 px-4 sm:px-6 bg-white">
        <div className="max-w-6xl mx-auto">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
            {/* Contact Form */}
            <div>
              <span className="inline-block px-4 py-2 bg-accent-yellow border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
                Send a Message
              </span>
              <h2 className="font-display text-3xl font-bold mb-6">
                We'd Love to Hear From You
              </h2>

              {isSubmitted ? (
                <Card className="p-8 bg-accent-green text-center">
                  <CheckCircle className="w-16 h-16 mx-auto mb-4" />
                  <h3 className="font-display text-2xl font-bold mb-2">Message Sent!</h3>
                  <p className="font-body text-lg mb-4">
                    Thank you for reaching out. We'll get back to you within 24-48 hours.
                  </p>
                  <Button onClick={() => setIsSubmitted(false)}>
                    Send Another Message
                  </Button>
                </Card>
              ) : (
                <Card className="p-6 sm:p-8">
                  {error && (
                    <Alert variant="error" className="mb-6">
                      {error}
                    </Alert>
                  )}

                  <form onSubmit={handleSubmit}>
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
                      <Input
                        label="Your Name"
                        name="name"
                        placeholder="John Doe"
                        value={formData.name}
                        onChange={handleChange}
                        required
                      />
                      <Input
                        label="Email Address"
                        name="email"
                        type="email"
                        placeholder="you@example.com"
                        value={formData.email}
                        onChange={handleChange}
                        required
                      />
                    </div>

                    <div className="mb-4">
                      <label className="block font-display font-bold text-sm uppercase mb-2">
                        Subject
                      </label>
                      <select
                        name="subject"
                        value={formData.subject}
                        onChange={handleChange}
                        required
                        className="w-full px-4 py-3 bg-white border-4 border-black font-body text-lg shadow-brutal-sm focus:shadow-brutal focus:outline-none cursor-pointer"
                      >
                        <option value="">Select a subject</option>
                        <option value="general">General Inquiry</option>
                        <option value="technical">Technical Support</option>
                        <option value="feedback">Feedback & Suggestions</option>
                        <option value="bug">Report a Bug</option>
                        <option value="partnership">Partnership Inquiry</option>
                        <option value="other">Other</option>
                      </select>
                    </div>

                    <div className="mb-6">
                      <label className="block font-display font-bold text-sm uppercase mb-2">
                        Message
                      </label>
                      <textarea
                        name="message"
                        rows={5}
                        placeholder="How can we help you?"
                        value={formData.message}
                        onChange={handleChange}
                        required
                        className="w-full px-4 py-3 bg-white border-4 border-black font-body text-lg shadow-brutal-sm focus:shadow-brutal focus:outline-none resize-none"
                      />
                    </div>

                    <Button
                      type="submit"
                      isLoading={isSubmitting}
                      className="w-full"
                      size="lg"
                    >
                      {isSubmitting ? 'Sending...' : 'Send Message'}
                      {!isSubmitting && <Send className="inline ml-2 w-5 h-5" />}
                    </Button>
                  </form>
                </Card>
              )}
            </div>

            {/* Response Time & FAQ */}
            <div>
              {/* Response Time Card */}
              <Card className="p-6 bg-secondary text-white mb-8">
                <div className="flex items-center gap-4 mb-4">
                  <div className="w-14 h-14 bg-primary border-4 border-black flex items-center justify-center">
                    <Clock className="w-7 h-7 text-black" />
                  </div>
                  <div>
                    <h3 className="font-display text-xl font-bold">Response Time</h3>
                    <p className="font-body text-white/80">We typically respond within 24-48 hours</p>
                  </div>
                </div>
                <p className="font-body text-white/90">
                  Our support team is available Monday through Friday, 9am to 5pm (WAT).
                  For urgent issues, please include "URGENT" in your subject line.
                </p>
              </Card>

              {/* FAQ Section */}
              <div>
                <span className="inline-block px-4 py-2 bg-primary border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
                  FAQ
                </span>
                <h3 className="font-display text-2xl font-bold mb-6">
                  Frequently Asked Questions
                </h3>
                <div className="space-y-4">
                  {faqs.map((faq, index) => (
                    <Card key={index} className="p-4">
                      <h4 className="font-display font-bold mb-2">{faq.question}</h4>
                      <p className="font-body text-gray-700 text-sm">{faq.answer}</p>
                    </Card>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Map/Location Section */}
      <section className="py-12 px-4 sm:px-6 bg-black text-white">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="font-display text-3xl font-bold mb-4">
            Based in <span className="text-primary">Lagos, Nigeria</span>
          </h2>
          <p className="font-body text-lg text-white/80 mb-6">
            ScoreThatExam is proudly Nigerian-built, serving students across the nation
            and beyond.
          </p>
          <div className="inline-flex items-center gap-4">
            <div className="w-12 h-12 bg-accent-green border-4 border-white flex items-center justify-center">
              <span className="text-2xl">🇳🇬</span>
            </div>
            <span className="font-display font-bold uppercase">Made with ❤️ in Nigeria</span>
          </div>
        </div>
      </section>
    </Layout>
  );
}

export default Contact;

