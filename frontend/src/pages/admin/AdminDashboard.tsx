import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Layout } from '../../components/layout';
import { Button, Card, Alert, Input, Select, LoadingSpinner } from '../../components/ui';
import { adminService, leaderboardService } from '../../services';
import { LeaderboardTable } from '../../components/leaderboard';
import type { Subject, QuestionData, LeaderboardEntry } from '../../types';
import {
  Shield,
  Upload,
  Plus,
  BookOpen,
  Trophy,
  FileJson,
  CheckCircle,
  Copy,
  LogOut,
} from 'lucide-react';

type ActiveTab = 'bulk' | 'single' | 'subjects' | 'leaderboard';

const SAMPLE_JSON: QuestionData[] = [
  {
    name: "What is the capital of Nigeria?",
    options: ["Lagos", "Abuja", "Kano", "Port Harcourt"],
    answer: "Abuja",
    explanation: "Abuja became the capital of Nigeria in 1991, replacing Lagos."
  },
  {
    name: "Which river is the longest in Nigeria?",
    options: ["River Benue", "River Niger", "River Kaduna", "River Ogun"],
    answer: "River Niger",
    explanation: "River Niger is the longest river in Nigeria, flowing through several states."
  }
];

function AdminDashboard() {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<ActiveTab>('bulk');
  const [subjects, setSubjects] = useState<Subject[]>([]);
  const [loadingSubjects, setLoadingSubjects] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  // Bulk upload state
  const [bulkSubjectId, setBulkSubjectId] = useState('');
  const [bulkJson, setBulkJson] = useState('');
  const [bulkLoading, setBulkLoading] = useState(false);

  // Single question state
  const [singleSubjectId, setSingleSubjectId] = useState('');
  const [singleQuestion, setSingleQuestion] = useState('');
  const [singleOptions, setSingleOptions] = useState(['', '', '', '']);
  const [singleAnswer, setSingleAnswer] = useState('');
  const [singleExplanation, setSingleExplanation] = useState('');
  const [singleLoading, setSingleLoading] = useState(false);

  // Subject creation state
  const [newSubjectName, setNewSubjectName] = useState('');
  const [subjectLoading, setSubjectLoading] = useState(false);

  // Leaderboard state
  const [leaderboardEntries, setLeaderboardEntries] = useState<LeaderboardEntry[]>([]);
  const [leaderboardLoading, setLeaderboardLoading] = useState(false);

  useEffect(() => {
    // Check if admin
    if (!adminService.isAdmin()) {
      navigate('/admin/login');
      return;
    }
    fetchSubjects();
  }, [navigate]);

  const fetchSubjects = async () => {
    try {
      setLoadingSubjects(true);
      const data = await adminService.getAllSubjects();
      setSubjects(data);
    } catch (err) {
      setError('Failed to load subjects');
    } finally {
      setLoadingSubjects(false);
    }
  };

  const fetchLeaderboard = async () => {
    try {
      setLeaderboardLoading(true);
      const data = await leaderboardService.getLeaderboard({ limit: 100 });
      setLeaderboardEntries(data.entries || []);
    } catch (err) {
      setError('Failed to load leaderboard');
    } finally {
      setLeaderboardLoading(false);
    }
  };

  useEffect(() => {
    if (activeTab === 'leaderboard') {
      fetchLeaderboard();
    }
  }, [activeTab]);

  const handleLogout = () => {
    adminService.logout();
    navigate('/admin/login');
  };

  const clearMessages = () => {
    setError(null);
    setSuccess(null);
  };

  // Bulk upload handler
  const handleBulkUpload = async () => {
    clearMessages();
    if (!bulkSubjectId) {
      setError('Please select a subject');
      return;
    }
    if (!bulkJson.trim()) {
      setError('Please enter JSON data');
      return;
    }

    try {
      const questions = JSON.parse(bulkJson) as QuestionData[];
      if (!Array.isArray(questions)) {
        setError('JSON must be an array of questions');
        return;
      }

      setBulkLoading(true);
      await adminService.createBulkQuestions(parseInt(bulkSubjectId), questions);
      setSuccess(`Successfully uploaded ${questions.length} questions!`);
      setBulkJson('');
    } catch (err) {
      if (err instanceof SyntaxError) {
        setError('Invalid JSON format. Please check your JSON syntax.');
      } else {
        setError('Failed to upload questions. Please try again.');
      }
    } finally {
      setBulkLoading(false);
    }
  };

  // Single question handler
  const handleSingleSubmit = async () => {
    clearMessages();
    if (!singleSubjectId) {
      setError('Please select a subject');
      return;
    }
    if (!singleQuestion.trim()) {
      setError('Please enter a question');
      return;
    }
    const validOptions = singleOptions.filter(o => o.trim());
    if (validOptions.length < 2) {
      setError('Please enter at least 2 options');
      return;
    }
    if (!singleAnswer.trim()) {
      setError('Please enter the correct answer');
      return;
    }
    if (!singleExplanation.trim()) {
      setError('Please enter an explanation');
      return;
    }

    try {
      setSingleLoading(true);
      await adminService.createSingleQuestion(parseInt(singleSubjectId), {
        name: singleQuestion,
        options: validOptions,
        answer: singleAnswer,
        explanation: singleExplanation,
      });
      setSuccess('Question created successfully!');
      // Reset form
      setSingleQuestion('');
      setSingleOptions(['', '', '', '']);
      setSingleAnswer('');
      setSingleExplanation('');
    } catch (err) {
      setError('Failed to create question. Please try again.');
    } finally {
      setSingleLoading(false);
    }
  };

  // Subject creation handler
  const handleCreateSubject = async () => {
    clearMessages();
    if (!newSubjectName.trim()) {
      setError('Please enter a subject name');
      return;
    }

    try {
      setSubjectLoading(true);
      await adminService.createSubject({ name: newSubjectName });
      setSuccess(`Subject "${newSubjectName}" created successfully!`);
      setNewSubjectName('');
      fetchSubjects();
    } catch (err) {
      setError('Failed to create subject. It may already exist.');
    } finally {
      setSubjectLoading(false);
    }
  };

  const copyJsonTemplate = () => {
    navigator.clipboard.writeText(JSON.stringify(SAMPLE_JSON, null, 2));
    setSuccess('JSON template copied to clipboard!');
    setTimeout(() => setSuccess(null), 3000);
  };

  const tabs = [
    { id: 'bulk' as ActiveTab, label: 'Bulk Upload', icon: Upload },
    { id: 'single' as ActiveTab, label: 'Single Question', icon: Plus },
    { id: 'subjects' as ActiveTab, label: 'Subjects', icon: BookOpen },
    { id: 'leaderboard' as ActiveTab, label: 'Leaderboard', icon: Trophy },
  ];

  return (
    <Layout showFooter={false}>
      <div className="px-4 sm:px-6 py-8">
        <div className="max-w-7xl mx-auto">
          {/* Header */}
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-8">
            <div className="flex items-center gap-4">
              <div className="w-14 h-14 bg-secondary border-4 border-black shadow-brutal flex items-center justify-center">
                <Shield className="w-7 h-7 text-white" />
              </div>
              <div>
                <h1 className="font-display text-3xl font-bold">Admin Dashboard</h1>
                <p className="font-body text-gray-600">Manage questions, subjects, and view leaderboards</p>
              </div>
            </div>
            <Button variant="outline" onClick={handleLogout}>
              <LogOut className="inline mr-2 w-4 h-4" />
              Logout
            </Button>
          </div>

          {/* Alerts */}
          {error && (
            <Alert variant="error" className="mb-6">
              {error}
            </Alert>
          )}
          {success && (
            <Alert variant="success" className="mb-6">
              {success}
            </Alert>
          )}

          {/* Tabs */}
          <div className="flex flex-wrap gap-2 mb-6">
            {tabs.map((tab) => (
              <button
                key={tab.id}
                onClick={() => { setActiveTab(tab.id); clearMessages(); }}
                className={`flex items-center gap-2 px-4 py-3 font-display font-bold uppercase border-4 border-black transition-all ${
                  activeTab === tab.id
                    ? 'bg-secondary text-white shadow-brutal'
                    : 'bg-white hover:bg-cream'
                }`}
              >
                <tab.icon className="w-5 h-5" />
                {tab.label}
              </button>
            ))}
          </div>

          {/* Tab Content */}
          {activeTab === 'bulk' && (
            <Card className="p-6 sm:p-8">
              <h2 className="font-display text-2xl font-bold mb-6 flex items-center gap-3">
                <FileJson className="w-7 h-7" />
                Bulk Question Upload
              </h2>

              {/* JSON Structure Guide */}
              <Card className="mb-6 p-4 bg-accent-yellow">
                <div className="flex items-start justify-between gap-4">
                  <div>
                    <h3 className="font-display font-bold mb-2">JSON Structure Required:</h3>
                    <pre className="font-mono text-sm bg-white p-4 border-3 border-black overflow-x-auto">
{JSON.stringify(SAMPLE_JSON, null, 2)}
                    </pre>
                  </div>
                  <Button variant="outline" size="sm" onClick={copyJsonTemplate}>
                    <Copy className="w-4 h-4" />
                  </Button>
                </div>
              </Card>

              {loadingSubjects ? (
                <LoadingSpinner />
              ) : (
                <>
                  <Select
                    label="Select Subject"
                    options={subjects.map(s => ({ value: s.id.toString(), label: s.name }))}
                    placeholder="Choose a subject"
                    value={bulkSubjectId}
                    onChange={(e) => setBulkSubjectId(e.target.value)}
                  />

                  <div className="mb-4">
                    <label className="block font-display font-bold text-sm uppercase mb-2">
                      Questions JSON
                    </label>
                    <textarea
                      className="w-full h-64 px-4 py-3 font-mono text-sm bg-white border-4 border-black shadow-brutal-sm focus:shadow-brutal focus:outline-none"
                      placeholder="Paste your JSON array of questions here..."
                      value={bulkJson}
                      onChange={(e) => setBulkJson(e.target.value)}
                    />
                  </div>

                  <Button
                    onClick={handleBulkUpload}
                    isLoading={bulkLoading}
                    disabled={!bulkSubjectId || !bulkJson.trim()}
                    size="lg"
                  >
                    <Upload className="inline mr-2 w-5 h-5" />
                    Upload Questions
                  </Button>
                </>
              )}
            </Card>
          )}

          {activeTab === 'single' && (
            <Card className="p-6 sm:p-8">
              <h2 className="font-display text-2xl font-bold mb-6 flex items-center gap-3">
                <Plus className="w-7 h-7" />
                Create Single Question
              </h2>

              {loadingSubjects ? (
                <LoadingSpinner />
              ) : (
                <div className="space-y-4">
                  <Select
                    label="Select Subject"
                    options={subjects.map(s => ({ value: s.id.toString(), label: s.name }))}
                    placeholder="Choose a subject"
                    value={singleSubjectId}
                    onChange={(e) => setSingleSubjectId(e.target.value)}
                  />

                  <div>
                    <label className="block font-display font-bold text-sm uppercase mb-2">
                      Question
                    </label>
                    <textarea
                      className="w-full h-24 px-4 py-3 font-body text-lg bg-white border-4 border-black shadow-brutal-sm focus:shadow-brutal focus:outline-none"
                      placeholder="Enter the question..."
                      value={singleQuestion}
                      onChange={(e) => setSingleQuestion(e.target.value)}
                    />
                  </div>

                  <div>
                    <label className="block font-display font-bold text-sm uppercase mb-2">
                      Options (at least 2 required)
                    </label>
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
                      {singleOptions.map((opt, idx) => (
                        <Input
                          key={idx}
                          placeholder={`Option ${String.fromCharCode(65 + idx)}`}
                          value={opt}
                          onChange={(e) => {
                            const newOptions = [...singleOptions];
                            newOptions[idx] = e.target.value;
                            setSingleOptions(newOptions);
                          }}
                        />
                      ))}
                    </div>
                    <Button
                      variant="outline"
                      size="sm"
                      className="mt-2"
                      onClick={() => setSingleOptions([...singleOptions, ''])}
                    >
                      + Add Option
                    </Button>
                  </div>

                  <Input
                    label="Correct Answer"
                    placeholder="Enter the correct answer (must match one of the options exactly)"
                    value={singleAnswer}
                    onChange={(e) => setSingleAnswer(e.target.value)}
                  />

                  <div>
                    <label className="block font-display font-bold text-sm uppercase mb-2">
                      Explanation
                    </label>
                    <textarea
                      className="w-full h-24 px-4 py-3 font-body text-lg bg-white border-4 border-black shadow-brutal-sm focus:shadow-brutal focus:outline-none"
                      placeholder="Explain why this is the correct answer..."
                      value={singleExplanation}
                      onChange={(e) => setSingleExplanation(e.target.value)}
                    />
                  </div>

                  <Button
                    onClick={handleSingleSubmit}
                    isLoading={singleLoading}
                    size="lg"
                  >
                    <CheckCircle className="inline mr-2 w-5 h-5" />
                    Create Question
                  </Button>
                </div>
              )}
            </Card>
          )}

          {activeTab === 'subjects' && (
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {/* Create Subject */}
              <Card className="p-6">
                <h2 className="font-display text-2xl font-bold mb-6 flex items-center gap-3">
                  <Plus className="w-7 h-7" />
                  Create Subject
                </h2>

                <div className="space-y-4">
                  <Input
                    label="Subject Name"
                    placeholder="e.g., Mathematics, Physics, Chemistry"
                    value={newSubjectName}
                    onChange={(e) => setNewSubjectName(e.target.value)}
                  />

                  <Button
                    onClick={handleCreateSubject}
                    isLoading={subjectLoading}
                    disabled={!newSubjectName.trim()}
                  >
                    <Plus className="inline mr-2 w-5 h-5" />
                    Create Subject
                  </Button>
                </div>
              </Card>

              {/* Subject List */}
              <Card className="p-6">
                <h2 className="font-display text-2xl font-bold mb-6 flex items-center gap-3">
                  <BookOpen className="w-7 h-7" />
                  Existing Subjects
                </h2>

                {loadingSubjects ? (
                  <LoadingSpinner />
                ) : subjects.length === 0 ? (
                  <p className="font-body text-gray-600">No subjects created yet.</p>
                ) : (
                  <div className="space-y-3">
                    {subjects.map((subject) => (
                      <div
                        key={subject.id}
                        className="flex items-center justify-between p-4 bg-cream border-3 border-black"
                      >
                        <div>
                          <span className="font-display font-bold">{subject.name}</span>
                          <span className="ml-2 text-sm text-gray-500">ID: {subject.id}</span>
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </Card>
            </div>
          )}

          {activeTab === 'leaderboard' && (
            <Card className="p-6 sm:p-8">
              <h2 className="font-display text-2xl font-bold mb-6 flex items-center gap-3">
                <Trophy className="w-7 h-7" />
                Leaderboard
              </h2>

              {leaderboardLoading ? (
                <LoadingSpinner />
              ) : (
                <LeaderboardTable
                  entries={leaderboardEntries}
                  isLoading={leaderboardLoading}
                />
              )}
            </Card>
          )}
        </div>
      </div>
    </Layout>
  );
}

export default AdminDashboard;

