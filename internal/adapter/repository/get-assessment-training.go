package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetAssessmentTraining(req domain.GetAssessmentTrainingMethodsRequest) ([]domain.GetAssessmentTrainingMethodsResponse, error) {
	var result []domain.GetAssessmentTrainingMethodsResponse

	query := `
		SELECT 
			astm.training_methods_id,
			asq.assessment_question_id,
			asq.assessment_type_id,
			asq.assessment_no,
			astm.training_text,
			astm.training_required_tools
		FROM assessment_training_methods astm
		INNER JOIN assessment_question asq
			ON astm.assessment_question_id = asq.assessment_question_id
		WHERE asq.assessment_type_id = ?
		ORDER BY asq.assessment_no ASC;
	`
	rows, err := r.db.Query(query, req.AssessmentTypeId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Assessment domain.GetAssessmentTrainingMethodsResponse
		if err := rows.Scan(
			&Assessment.TrainingMethodsId,
			&Assessment.AssessmentQuestionId,
			&Assessment.AssessmentTypeId,
			&Assessment.AssessmentNo,
			&Assessment.TrainingText,
			&Assessment.TrainingRequiredTools,
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
