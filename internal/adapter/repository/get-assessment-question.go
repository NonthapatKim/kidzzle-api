package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetAssessmentQuestion(req domain.GetAssessmentQuestionRequest) ([]domain.GetAssessmentQuestionResponse, error) {
	var result []domain.GetAssessmentQuestionResponse

	query := `
		SELECT
			asq.assessment_question_id,
			asq.assessment_type_id,
			asq.age_range_id,
			asq.dev_type_id,
			asq.assessment_no,
			ast.type_name AS assessment_type_name,
			age.age_range AS age_range_name,
			dev.development_type,
			asq.question_text,
			asq.assessment_method,
			asq.assessment_required_tool,
			asq.pass_criteria
		FROM assessment_question asq
		INNER JOIN assessment_type ast
			ON asq.assessment_type_id = ast.assessment_type_id
		INNER JOIN age_range age
			ON asq.age_range_id = age.age_range_id
		INNER JOIN development_type dev
			ON asq.dev_type_id = dev.dev_type_id
		WHERE asq.assessment_type_id = ?
		ORDER BY asq.assessment_no ASC;
	`
	rows, err := r.db.Query(query, req.AssessmentTypeId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Assessment domain.GetAssessmentQuestionResponse
		if err := rows.Scan(
			&Assessment.AssessmentQuestionId,
			&Assessment.AssessmentTypeId,
			&Assessment.AgeRangeId,
			&Assessment.DevTypeId,
			&Assessment.AssessmentNo,
			&Assessment.AssessmentTypeName,
			&Assessment.AgeRangeName,
			&Assessment.DevTypeName,
			&Assessment.QuestionText,
			&Assessment.AssessmentMethod,
			&Assessment.AssessmentRequiredtTool,
			&Assessment.PassCriteria,
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
