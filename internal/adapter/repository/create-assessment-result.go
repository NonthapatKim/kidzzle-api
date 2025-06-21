package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) CreateAssessmentResult(req domain.CreateAssessmentResultRequest) (domain.CreateAssessmentResultResponse, error) {
	query := `
		INSERT INTO assessment_result (
			kid_id,
			assessment_question_id,
			is_passed
		) VALUES (?, ?, ?)
	`
	isPassedStr := "false"
	if req.IsPassed {
		isPassedStr = "true"
	}

	_, err := r.db.Exec(
		query,
		req.KidId,
		req.AssessmentQuestionId,
		isPassedStr,
	)
	if err != nil {
		return domain.CreateAssessmentResultResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.CreateAssessmentResultResponse{}, nil
}
