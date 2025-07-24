package dto

import "kvizo-api/internal/repositories"

type QuestionRequest struct {
	Title   string                    `json:"title" binding:"required"`
	OptionA string                    `json:"option_a" binding:"required"`
	OptionB string                    `json:"option_b" binding:"required"`
	OptionC string                    `json:"option_c" binding:"required"`
	OptionD string                    `json:"option_d" binding:"required"`
	Answer  repositories.AnswerOption `json:"answer" binding:"required,oneof=1 2 3 4"`
}
