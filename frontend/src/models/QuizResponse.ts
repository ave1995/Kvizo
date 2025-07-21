export interface QuizResponse {
  id: string;             // UUID as string
  title: string;
  description: string;
  created_at: string;     // ISO date string (from time.Time)
  updated_at: string;     // ISO date string (from time.Time)
}
