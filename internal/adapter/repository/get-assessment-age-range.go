package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetAssessmentAgeRange(req domain.GetAssessmentAgeRangeRequest) ([]domain.GetAssessmentAgeRangeResponse, error) {
	var result []domain.GetAssessmentAgeRangeResponse

	query := `
		SELECT 
			age_range_id,
			age_range,
			min_months,
			max_months,
			assessment_video_url_dspm,
			assessment_video_url_daim
		FROM age_range
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ageRange domain.GetAssessmentAgeRangeResponse
		if err := rows.Scan(
			&ageRange.AgeRangeId,
			&ageRange.AgeRange,
			&ageRange.MinMonths,
			&ageRange.MaxMonths,
			&ageRange.AssessmentVideoUrlDspm,
			&ageRange.AssessmentVideoUrlDaim,
		); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, ageRange)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
