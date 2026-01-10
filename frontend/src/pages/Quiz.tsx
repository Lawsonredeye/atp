import { useEffect, useState, useCallback } from 'react';
import { useSearchParams, useNavigate, Link } from 'react-router-dom';
import { Layout } from '../components/layout';
import { Button, Card, LoadingSpinner, Alert } from '../components/ui';
import { QuestionCard, QuizProgress, QuizResultCard } from '../components/quiz';
import { quizService, subjectService } from '../services';
import { useAuth } from '../context/AuthContext';
import type { GeneratedQuiz, QuizSubmitResponse, SubmitQuizRequest, Subject, QuizResult } from '../types';
import { ArrowLeft, ArrowRight, Send, RotateCcw, Home, ChevronLeft, ChevronRight } from 'lucide-react';

type QuizState = 'loading' | 'setup' | 'active' | 'submitting' | 'result' | 'review' | 'error';

function Quiz() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { refreshUser } = useAuth();

  // Quiz state
  const [quizState, setQuizState] = useState<QuizState>('loading');
  const [quiz, setQuiz] = useState<GeneratedQuiz | null>(null);
  const [subject, setSubject] = useState<Subject | null>(null);
  const [result, setResult] = useState<QuizSubmitResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  // Current question tracking
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [answers, setAnswers] = useState<Map<number, number[]>>(new Map());

  // Get params from URL
  const subjectId = searchParams.get('subject');
  const numQuestions = parseInt(searchParams.get('questions') || '10', 10);

  // Load quiz on mount
  useEffect(() => {
    const initQuiz = async () => {
      if (!subjectId) {
        setQuizState('setup');
        return;
      }

      try {
        setQuizState('loading');
        setError(null);

        // Fetch subject info
        const subjectData = await subjectService.getSubjectById(parseInt(subjectId));
        setSubject(subjectData);

        // Generate quiz
        const quizData = await quizService.createQuiz({
          subject_id: parseInt(subjectId),
          num_of_questions: numQuestions,
        });
        setQuiz(quizData);
        setQuizState('active');
      } catch (err) {
        setError('Failed to load quiz. Please try again.');
        setQuizState('error');
      }
    };

    initQuiz();
  }, [subjectId, numQuestions]);

  // Handle option selection
  const handleSelectOption = useCallback((questionId: number, optionId: number, isMultipleChoice: boolean) => {
    setAnswers((prev) => {
      const newAnswers = new Map(prev);
      const currentOptions = newAnswers.get(questionId) || [];

      if (isMultipleChoice) {
        // Toggle for multiple choice
        if (currentOptions.includes(optionId)) {
          newAnswers.set(questionId, currentOptions.filter((id) => id !== optionId));
        } else {
          newAnswers.set(questionId, [...currentOptions, optionId]);
        }
      } else {
        // Single select
        newAnswers.set(questionId, [optionId]);
      }

      return newAnswers;
    });
  }, []);

  // Navigation
  const goToNextQuestion = () => {
    if (quiz && currentQuestionIndex < quiz.questions.length - 1) {
      setCurrentQuestionIndex((prev) => prev + 1);
    }
  };

  const goToPrevQuestion = () => {
    if (currentQuestionIndex > 0) {
      setCurrentQuestionIndex((prev) => prev - 1);
    }
  };

  const goToQuestion = (index: number) => {
    setCurrentQuestionIndex(index);
  };

  // Submit quiz
  const handleSubmitQuiz = async () => {
    if (!quiz) return;

    try {
      setQuizState('submitting');
      setError(null);

      const submitData: SubmitQuizRequest[] = quiz.questions.map((q) => ({
        question_id: q.question_id,
        is_multiple_choice: q.is_multiple_choice,
        option_ids: answers.get(q.question_id) || [],
      }));

      const resultData = await quizService.submitQuiz(submitData);
      setResult(resultData);
      setQuizState('result');

      // Refresh user stats
      await refreshUser();
    } catch (err) {
      setError('Failed to submit quiz. Please try again.');
      setQuizState('active');
    }
  };

  // Start new quiz
  const handleStartNewQuiz = () => {
    navigate('/dashboard');
  };

  // Review answers
  const handleReviewAnswers = () => {
    setCurrentQuestionIndex(0);
    setQuizState('review');
  };

  // Back to results
  const handleBackToResults = () => {
    setQuizState('result');
  };

  // Get result for current question in review mode
  const getQuestionResult = (questionId: number): QuizResult | undefined => {
    return result?.results.find((r) => r.question_id === questionId);
  };

  // Render based on state
  if (quizState === 'loading') {
    return (
      <Layout showFooter={false}>
        <div className="min-h-[60vh] flex flex-col items-center justify-center px-4">
          <LoadingSpinner size="lg" />
          <p className="mt-4 font-display font-bold text-xl uppercase">
            Generating your quiz...
          </p>
        </div>
      </Layout>
    );
  }

  if (quizState === 'setup' || !subjectId) {
    return (
      <Layout>
        <div className="min-h-[60vh] flex items-center justify-center px-4 py-12">
          <Card className="max-w-md w-full p-8 text-center">
            <h1 className="font-display text-2xl font-bold mb-4">No Subject Selected</h1>
            <p className="font-body text-gray-600 mb-6">
              Please select a subject from the dashboard to start a quiz.
            </p>
            <Link to="/dashboard">
              <Button className="w-full">
                <Home className="inline mr-2 w-5 h-5" />
                Go to Dashboard
              </Button>
            </Link>
          </Card>
        </div>
      </Layout>
    );
  }

  if (quizState === 'error') {
    return (
      <Layout>
        <div className="min-h-[60vh] flex items-center justify-center px-4 py-12">
          <Card className="max-w-md w-full p-8 text-center">
            <Alert variant="error" className="mb-6">
              {error}
            </Alert>
            <div className="flex flex-col gap-4">
              <Button onClick={() => window.location.reload()}>
                <RotateCcw className="inline mr-2 w-5 h-5" />
                Try Again
              </Button>
              <Link to="/dashboard">
                <Button variant="outline" className="w-full">
                  Back to Dashboard
                </Button>
              </Link>
            </div>
          </Card>
        </div>
      </Layout>
    );
  }

  if (quizState === 'result' && result) {
    return (
      <Layout>
        <div className="px-4 sm:px-6 py-8 sm:py-12">
          <QuizResultCard
            result={result}
            onReviewAnswers={handleReviewAnswers}
            onStartNewQuiz={handleStartNewQuiz}
          />
        </div>
      </Layout>
    );
  }

  if (!quiz) return null;

  const currentQuestion = quiz.questions[currentQuestionIndex];
  const answeredCount = answers.size;
  const isLastQuestion = currentQuestionIndex === quiz.questions.length - 1;
  const isFirstQuestion = currentQuestionIndex === 0;
  const canSubmit = answeredCount === quiz.questions.length;

  // Review mode
  if (quizState === 'review' && result) {
    const questionResult = getQuestionResult(currentQuestion.question_id);

    return (
      <Layout showFooter={false}>
        <div className="px-4 sm:px-6 py-6 sm:py-8">
          <div className="max-w-4xl mx-auto">
            {/* Header */}
            <div className="mb-6 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
              <div>
                <Button variant="outline" size="sm" onClick={handleBackToResults}>
                  <ArrowLeft className="inline mr-2 w-4 h-4" />
                  Back to Results
                </Button>
              </div>
              <div className="flex items-center gap-2">
                <span className={`px-3 py-1 border-3 border-black font-display font-bold text-sm ${questionResult?.is_correct ? 'bg-accent-green' : 'bg-accent-red text-white'}`}>
                  {questionResult?.is_correct ? '✓ Correct' : '✗ Incorrect'}
                </span>
              </div>
            </div>

            {/* Question Card */}
            <QuestionCard
              question={currentQuestion}
              questionNumber={currentQuestionIndex + 1}
              totalQuestions={quiz.questions.length}
              selectedOptionIds={answers.get(currentQuestion.question_id) || []}
              onSelectOption={() => {}}
              showResult={true}
              correctAnswer={questionResult?.correct_answer}
            />

            {/* Explanation */}
            {questionResult?.explanation && (
              <Card className="mt-4 p-6 bg-accent-yellow">
                <h3 className="font-display font-bold uppercase mb-2">Explanation</h3>
                <p className="font-body">{questionResult.explanation}</p>
              </Card>
            )}

            {/* Navigation */}
            <div className="mt-6 flex items-center justify-between">
              <Button
                variant="outline"
                onClick={goToPrevQuestion}
                disabled={isFirstQuestion}
              >
                <ChevronLeft className="inline mr-1 w-5 h-5" />
                Previous
              </Button>

              <span className="font-display font-bold">
                {currentQuestionIndex + 1} / {quiz.questions.length}
              </span>

              <Button
                variant="outline"
                onClick={goToNextQuestion}
                disabled={isLastQuestion}
              >
                Next
                <ChevronRight className="inline ml-1 w-5 h-5" />
              </Button>
            </div>

            {/* Question Navigator */}
            <Card className="mt-6 p-4">
              <p className="font-display font-bold text-sm uppercase mb-3">Jump to Question</p>
              <div className="flex flex-wrap gap-2">
                {quiz.questions.map((q, index) => {
                  const qResult = getQuestionResult(q.question_id);
                  return (
                    <button
                      key={q.question_id}
                      onClick={() => goToQuestion(index)}
                      className={`w-10 h-10 border-3 border-black font-display font-bold transition-all ${
                        index === currentQuestionIndex
                          ? 'bg-secondary text-white shadow-brutal-sm'
                          : qResult?.is_correct
                          ? 'bg-accent-green'
                          : 'bg-accent-red text-white'
                      }`}
                    >
                      {index + 1}
                    </button>
                  );
                })}
              </div>
            </Card>
          </div>
        </div>
      </Layout>
    );
  }

  // Active quiz mode
  return (
    <Layout showFooter={false}>
      <div className="px-4 sm:px-6 py-6 sm:py-8">
        <div className="max-w-4xl mx-auto">
          {/* Header */}
          <div className="mb-6 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div>
              <h1 className="font-display text-2xl font-bold">
                {subject?.name || 'Quiz'}
              </h1>
              <p className="font-body text-gray-600">
                Answer all questions to submit your quiz
              </p>
            </div>
            <Link to="/dashboard">
              <Button variant="outline" size="sm">
                <ArrowLeft className="inline mr-2 w-4 h-4" />
                Exit Quiz
              </Button>
            </Link>
          </div>

          {/* Progress */}
          <QuizProgress
            total={quiz.questions.length}
            answeredCount={answeredCount}
          />

          {/* Error Alert */}
          {error && (
            <Alert variant="error" className="mb-6">
              {error}
            </Alert>
          )}

          {/* Question Card */}
          <QuestionCard
            question={currentQuestion}
            questionNumber={currentQuestionIndex + 1}
            totalQuestions={quiz.questions.length}
            selectedOptionIds={answers.get(currentQuestion.question_id) || []}
            onSelectOption={(optionId) =>
              handleSelectOption(currentQuestion.question_id, optionId, currentQuestion.is_multiple_choice)
            }
          />

          {/* Navigation */}
          <div className="mt-6 flex flex-col sm:flex-row items-center justify-between gap-4">
            <Button
              variant="outline"
              onClick={goToPrevQuestion}
              disabled={isFirstQuestion}
              className="w-full sm:w-auto"
            >
              <ArrowLeft className="inline mr-2 w-5 h-5" />
              Previous
            </Button>

            <div className="flex gap-4 w-full sm:w-auto">
              {!isLastQuestion ? (
                <Button onClick={goToNextQuestion} className="flex-1 sm:flex-none">
                  Next
                  <ArrowRight className="inline ml-2 w-5 h-5" />
                </Button>
              ) : (
                <Button
                  onClick={handleSubmitQuiz}
                  disabled={!canSubmit}
                  isLoading={quizState === 'submitting'}
                  className="flex-1 sm:flex-none bg-accent-green text-black hover:bg-green-400"
                >
                  <Send className="inline mr-2 w-5 h-5" />
                  {quizState === 'submitting' ? 'Submitting...' : 'Submit Quiz'}
                </Button>
              )}
            </div>
          </div>

          {/* Question Navigator */}
          <Card className="mt-6 p-4 bg-cream">
            <p className="font-display font-bold text-sm uppercase mb-3">Question Navigator</p>
            <div className="flex flex-wrap gap-2">
              {quiz.questions.map((q, index) => {
                const isAnswered = answers.has(q.question_id);
                const isCurrent = index === currentQuestionIndex;

                return (
                  <button
                    key={q.question_id}
                    onClick={() => goToQuestion(index)}
                    className={`w-10 h-10 border-3 border-black font-display font-bold transition-all ${
                      isCurrent
                        ? 'bg-secondary text-white shadow-brutal-sm'
                        : isAnswered
                        ? 'bg-accent-green'
                        : 'bg-white hover:bg-gray-100'
                    }`}
                  >
                    {index + 1}
                  </button>
                );
              })}
            </div>
            {!canSubmit && (
              <p className="mt-3 font-body text-sm text-gray-600">
                Answer all {quiz.questions.length} questions to submit. ({answeredCount}/{quiz.questions.length} answered)
              </p>
            )}
          </Card>
        </div>
      </div>
    </Layout>
  );
}

export default Quiz;

