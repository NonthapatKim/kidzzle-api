package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetPromotionColor() ([]domain.GetPromotionColorResponse, error) {
	var result []domain.GetPromotionColorResponse

	query := `
		SELECT
			color_id,
			color_name,
			color_shape,
			color_image_url,
			color_sound_file_url
		FROM promotion_color;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Promotion domain.GetPromotionColorResponse
		if err := rows.Scan(
			&Promotion.ColorId,
			&Promotion.ColorName,
			&Promotion.ColorShape,
			&Promotion.ColorImageUrl,
			&Promotion.ColorSoundFileUrl,
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
