import { useState, useEffect } from "react";
import { fetchQuizzes } from "../../services/quizService";
import type { QuizResponse } from "../../models/QuizResponse";
import type { NavigationProps } from "../../interfaces/navigation";

type QuizListProps = NavigationProps;

const QuizList: React.FC<QuizListProps> = ({ navigateTo }) => {
  const [quizzes, setQuizzes] = useState<QuizResponse[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const getQuizzes = async () => {
      try {
        const data = await fetchQuizzes();
        setQuizzes(data);
      } catch {
        setError("Failed to load quizzes. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    getQuizzes();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-48">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-blue-500"></div>
        <p className="ml-4 text-lg text-gray-700">Loading quizzes...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center p-6 bg-red-100 border border-red-400 text-red-700 rounded-lg mx-auto max-w-md">
        <p className="font-bold">Error:</p>
        <p>{error}</p>
        <button
          onClick={() => navigateTo("/")}
          className="mt-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors"
        >
          Go Home
        </button>
      </div>
    );
  }

  return (
    <div className="p-6 max-w-3xl mx-auto bg-white rounded-xl shadow-lg">
      <h2 className="text-3xl font-bold text-center text-gray-800 mb-8">
        Available Quizzes
      </h2>
      {quizzes.length === 0 ? (
        <p className="text-center text-gray-600 text-lg">No quizzes found.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {quizzes.map((quiz) => (
            <div
              key={quiz.id}
              className="bg-gray-50 p-6 rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300"
            >
              <h3 className="text-xl font-semibold text-gray-900 mb-2">
                {quiz.title}
              </h3>
              <p className="text-gray-700 text-base">{quiz.description}</p>
              <button
                onClick={() => console.log(`Starting quiz: ${quiz.title}`)}
                className="mt-4 w-full px-4 py-2 bg-green-500 text-white font-medium rounded-md hover:bg-green-600 transition-colors focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
              >
                Start Quiz
              </button>
            </div>
          ))}
        </div>
      )}
      <div className="text-center mt-8">
        <button
          onClick={() => navigateTo("/")}
          className="px-6 py-3 bg-indigo-600 text-white font-semibold rounded-lg shadow-md hover:bg-indigo-700 transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50"
        >
          Back to Home
        </button>
      </div>
    </div>
  );
};

export default QuizList;
