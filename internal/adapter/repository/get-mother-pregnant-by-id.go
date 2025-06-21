package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetMotherPregnantById(req domain.GetMotherPregnantByIdRequest) ([]domain.GetMotherPregnantByIdResponse, error) {
	var result []domain.GetMotherPregnantByIdResponse

	query := `
		SELECT
			mp.pregnant_id,
			mp.mother_name,
			mp.mother_birthday,
			mp.pregnant_congenital_disease,
			mp.pregnant_complications,
			mp.pregnant_drug_history,
			mp.created_at,
			mp.updated_at
		FROM mother_pregnant_history mp
		WHERE mp.user_id = ?;
	`
	rows, err := r.db.Query(query, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Mother domain.GetMotherPregnantByIdResponse
		if err := rows.Scan(
			&Mother.PregnantId,
			&Mother.MotherName,
			&Mother.MotherBirthday,
			&Mother.PregnantCongenitalDisease,
			&Mother.PregnantComplications,
			&Mother.PregnantDrugHistory,
			&Mother.CreatedAt,
			&Mother.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}

		result = append(result, Mother)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
