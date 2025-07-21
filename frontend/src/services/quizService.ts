import type { QuizResponse } from "../models/QuizResponse";

export const fetchQuizzes = async (): Promise<QuizResponse[]> => {
  try {
    const response = await fetch("/api/quizzes"); // adjust URL as needed

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    // Assuming your API returns an array with the correct fields,
    // otherwise, map and transform as needed
    const quizzes: QuizResponse[] = data.map((item: any) => ({
      id: item.id,
      title: item.title,
      description: item.description,
      created_at: item.created_at,
      updated_at: item.updated_at,
    }));

    return quizzes;
  } catch (error) {
    console.error("Error fetching quizzes:", error);
    throw error;
  }
};
