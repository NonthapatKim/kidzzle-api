package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetAssessmentResultDataById(req domain.GetAssessmentResultDataByIdRequest) ([]domain.GetAssessmentResultDataByIdResponse, error) {
	var result []domain.GetAssessmentResultDataByIdResponse

	query := `
		SELECT
			assr.assessment_result_id,
			assr.kid_id,
			assr.assessment_question_id,
			asq.assessment_type_id,
			asq.age_range_id,
			assr.is_passed,
			assr.created_at
		FROM assessment_result assr
		INNER JOIN assessment_question asq
			ON assr.assessment_question_id = asq.assessment_question_id
		WHERE assr.kid_id = ?
			AND asq.age_range_id = ?
			AND asq.assessment_type_id = ?
		ORDER BY asq.assessment_type_id, assr.created_at ASC;
	`
	rows, err := r.db.Query(query, req.KidId, req.AgeRageId, req.AssessmentTypeId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Assessment domain.GetAssessmentResultDataByIdResponse
		if err := rows.Scan(
			&Assessment.AssessmentResultId,
			&Assessment.KidId,
			&Assessment.AssessmentQuestionId,
			&Assessment.AssessmentTypeId,
			&Assessment.AgeRangeId,
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
