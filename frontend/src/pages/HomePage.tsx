import React from "react";
import type { NavigationProps } from "../interfaces/navigation";

type HomePageProps = NavigationProps;

const HomePage: React.FC<HomePageProps> = ({ navigateTo }) => {
  return (
    <main className="text-center p-8 bg-gray-100 min-h-screen flex flex-col items-center justify-center font-sans">
      <h1 className="text-5xl font-extrabold text-gray-900 mb-6">
        Welcome to the Quiz App!
      </h1>
      <p className="text-lg text-gray-700 mb-8 max-w-prose">
        Explore our collection of quizzes and test your knowledge. Click the button below to see available quizzes.
      </p>
      <section className="bg-white p-8 rounded-2xl shadow-xl flex flex-col items-center">
        <button
          onClick={() => navigateTo("/quizzes")}
          className="px-8 py-4 bg-blue-600 text-white text-xl font-bold rounded-full shadow-lg hover:bg-blue-700 transform hover:scale-105 transition-all duration-300 ease-in-out focus:outline-none focus:ring-4 focus:ring-blue-300 focus:ring-opacity-75"
        >
          Go to Quizzes
        </button>
      </section>
      <p className="mt-10 text-gray-500 text-sm">
        This is a simple demonstration of routing and API calls in React.
      </p>
    </main>
  );
};

export default HomePage;
