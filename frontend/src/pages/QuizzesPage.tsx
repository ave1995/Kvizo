import QuizList from "../components/QuizList/QuizList";
import type { NavigationProps } from "../interfaces/navigation";

type QuizPageProps = NavigationProps;

const QuizzesPage: React.FC<QuizPageProps> = ({ navigateTo }) => {
  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center py-10">
      <QuizList navigateTo={navigateTo} />
    </div>
  );
};

export default QuizzesPage;
