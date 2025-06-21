package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetPromotionBodyPart() ([]domain.GetPromotionBodyPartResponse, error) {
	var result []domain.GetPromotionBodyPartResponse

	query := `
		SELECT
			body_id,
			body_name,
			body_image_url,
			body_sound_file_url
		FROM promotion_body_part;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Promotion domain.GetPromotionBodyPartResponse
		if err := rows.Scan(
			&Promotion.BodyId,
			&Promotion.BodyName,
			&Promotion.BodyImageUrl,
			&Promotion.BodySoundFileUrl,
		); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, Promotion)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
