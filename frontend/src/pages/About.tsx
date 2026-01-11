import { Layout } from '../components/layout';
    import { Card } from '../components/ui';
    import { Target, Users, BookOpen, Award, Heart, Zap } from 'lucide-react';

    const values = [
      {
        icon: Target,
        color: 'bg-primary',
        description: 'We strive to provide the highest quality learning materials and practice questions.',
        title: 'Excellence',
      },
      {
        icon: Users,
        color: 'bg-accent-yellow',
        description: 'Education should be accessible to all students, regardless of their background.',
        title: 'Accessibility',
      },
      {
        icon: BookOpen,
        color: 'bg-accent-green',
        description: 'We continuously improve our platform with the latest educational technologies.',
        title: 'Innovation',
      },
      {
        icon: Award,
        color: 'bg-secondary',
        description: 'Your success is our success. We measure our impact by your achievements.',
        title: 'Success',
      },
    ];

    const stats = [
      { value: '50,000+', label: 'Active Students' },
      { value: '10,000+', label: 'Practice Questions' },
      { value: '15+', label: 'JAMB Subjects' },
      { value: '85%', label: 'Success Rate' },
    ];

    const team = [
      {
        name: 'Dr. Adaeze Okwu',
        role: 'Founder & CEO',
        description: 'Former JAMB examiner with 15+ years in education.',
      },
      {
        name: 'Chukwuemeka Obi',
        role: 'Head of Content',
        description: 'Curriculum specialist and educational content creator.',
      },
      {
        name: 'Fatima Bello',
        role: 'Lead Developer',
        description: 'Full-stack engineer passionate about EdTech.',
      },
    ];

    function About() {
      return (
        <Layout>
          {/* Hero Section */}
          <section className="bg-cream py-12 sm:py-16 lg:py-20 px-4 sm:px-6">
            <div className="max-w-7xl mx-auto">
              <div className="max-w-3xl">
                <span className="inline-block px-4 py-2 bg-primary border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-6">
                  About Us
                </span>
                <h1 className="font-display text-4xl sm:text-5xl lg:text-6xl font-bold mb-6">
                  Empowering Nigerian Students to{' '}
                  <span className="text-primary">Excel</span>
                </h1>
                <p className="font-body text-lg sm:text-xl text-gray-700">
                  ScoreThatExam is Nigeria's premier JAMB preparation platform, dedicated to helping
                  students achieve their university admission dreams through comprehensive practice
                  and personalized learning.
                </p>
              </div>
            </div>
          </section>

          {/* Mission Section */}
          <section className="py-12 sm:py-16 px-4 sm:px-6 bg-white">
            <div className="max-w-7xl mx-auto">
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12 items-center">
                <div>
                  <span className="inline-block px-3 py-1 bg-accent-yellow border-3 border-black font-display font-bold uppercase text-sm mb-4">
                    Our Mission
                  </span>
                  <h2 className="font-display text-3xl sm:text-4xl font-bold mb-6">
                    Making Quality JAMB Prep <span className="text-secondary">Accessible</span> to All
                  </h2>
                  <p className="font-body text-lg text-gray-700 mb-6">
                    We believe every Nigerian student deserves access to high-quality educational
                    resources. Our mission is to democratize JAMB preparation by providing free,
                    comprehensive, and engaging practice materials.
                  </p>
                  <p className="font-body text-lg text-gray-700">
                    Founded in 2024, ScoreThatExam has already helped thousands of students improve
                    their JAMB scores and secure admission to their dream universities.
                  </p>
                </div>
                <Card className="p-8 bg-secondary text-white">
                  <div className="flex items-center gap-4 mb-6">
                    <div className="w-16 h-16 bg-primary border-4 border-black flex items-center justify-center">
                      <Heart className="w-8 h-8 text-black" />
                    </div>
                    <div>
                      <h3 className="font-display text-2xl font-bold">Our Promise</h3>
                    </div>
                  </div>
                  <p className="font-body text-lg text-white/90">
                    "We promise to provide you with the most accurate, up-to-date JAMB practice
                    questions and the tools you need to track your progress. Your success is our
                    priority."
                  </p>
                </Card>
              </div>
            </div>
          </section>

          {/* Stats Section */}
          <section className="py-12 sm:py-16 px-4 sm:px-6 bg-black text-white">
            <div className="max-w-7xl mx-auto">
              <div className="grid grid-cols-2 lg:grid-cols-4 gap-6">
                {stats.map((stat, index) => (
                  <div key={index} className="text-center">
                    <div className="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-primary mb-2">
                      {stat.value}
                    </div>
                    <div className="font-body text-white/80">{stat.label}</div>
                  </div>
                ))}
              </div>
            </div>
          </section>

          {/* Values Section */}
          <section className="py-12 sm:py-16 px-4 sm:px-6 bg-cream">
            <div className="max-w-7xl mx-auto">
              <div className="text-center mb-12">
                <span className="inline-block px-4 py-2 bg-accent-green border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
                  Our Values
                </span>
                <h2 className="font-display text-3xl sm:text-4xl font-bold">
                  What Drives Us
                </h2>
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                {values.map((value, index) => (
                  <Card key={index} className="p-6 hover:shadow-brutal-lg transition-all duration-100 hover:-translate-y-1">
                    <div className={`w-14 h-14 ${value.color} border-4 border-black shadow-brutal-sm flex items-center justify-center mb-4`}>
                      <value.icon className="w-7 h-7" />
                    </div>
                    <h3 className="font-display text-xl font-bold mb-2">{value.title}</h3>
                    <p className="font-body text-gray-700">{value.description}</p>
                  </Card>
                ))}
              </div>
            </div>
          </section>

          {/* Team Section */}
          <section className="py-12 sm:py-16 px-4 sm:px-6 bg-white">
            <div className="max-w-7xl mx-auto">
              <div className="text-center mb-12">
                <span className="inline-block px-4 py-2 bg-primary border-4 border-black shadow-brutal-sm font-display font-bold uppercase text-sm mb-4">
                  Our Team
                </span>
                <h2 className="font-display text-3xl sm:text-4xl font-bold mb-4">
                  Meet the People Behind ScoreThatExam
                </h2>
                <p className="font-body text-lg text-gray-700 max-w-2xl mx-auto">
                  A passionate team of educators, developers, and designers committed to transforming
                  JAMB preparation in Nigeria.
                </p>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                {team.map((member, index) => (
                  <Card key={index} className="p-6 text-center">
                    <div className="w-24 h-24 bg-gray-200 border-4 border-black shadow-brutal mx-auto mb-4 flex items-center justify-center">
                      <Users className="w-12 h-12 text-gray-500" />
                    </div>
                    <h3 className="font-display text-xl font-bold mb-1">{member.name}</h3>
                    <p className="font-display text-primary font-bold text-sm uppercase mb-3">{member.role}</p>
                    <p className="font-body text-gray-700">{member.description}</p>
                  </Card>
                ))}
              </div>
            </div>
          </section>

          {/* CTA Section */}
          <section className="py-12 sm:py-16 px-4 sm:px-6 bg-primary">
            <div className="max-w-4xl mx-auto text-center">
              <Zap className="w-16 h-16 mx-auto mb-6" />
              <h2 className="font-display text-3xl sm:text-4xl font-bold mb-4">
                Ready to Start Your JAMB Prep?
              </h2>
              <p className="font-body text-lg mb-8 max-w-2xl mx-auto">
                Join over 50,000 students already using ScoreThatExam to prepare for their JAMB exams.
              </p>
              <a
                href="/register"
                className="inline-block px-8 py-4 bg-black text-white font-display font-bold uppercase border-4 border-black shadow-brutal hover:bg-secondary active:shadow-none active:translate-x-1 active:translate-y-1 transition-all duration-100"
              >
                Create Free Account
              </a>
            </div>
          </section>
        </Layout>
      );
    }

    export default About;