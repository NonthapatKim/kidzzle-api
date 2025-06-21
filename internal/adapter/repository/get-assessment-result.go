package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetAssessmentResultById(req domain.GetAssessmentResultByIdRequest) ([]domain.GetAssessmentResultByIdResponse, error) {
	var result []domain.GetAssessmentResultByIdResponse

	query := `
		SELECT
			assessment_result_id,
			kid_id,
			assessment_question_id,
			is_passed,
			created_at
		FROM assessment_result
		WHERE kid_id = ?
			AND assessment_question_id = ?
		ORDER BY created_at ASC;
	`
	rows, err := r.db.Query(query, req.KidId, req.AssessmentQuestionId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Assessment domain.GetAssessmentResultByIdResponse
		if err := rows.Scan(
			&Assessment.AssessmentResultId,
			&Assessment.KidId,
			&Assessment.AssessmentQuestionId,
			&Assessment.IsPassed,
			&Assessment.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, Assessment)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
