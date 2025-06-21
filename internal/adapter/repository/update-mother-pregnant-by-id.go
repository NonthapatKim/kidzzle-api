package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) UpdateMotherPregnantById(req domain.UpdateMotherPregnantByIdRequest) (domain.UpdateMotherPregnantByIdResponse, error) {
	query := `
		UPDATE mother_pregnant_history SET
			mother_name = ?,
			mother_birthday = ?,
			pregnant_congenital_disease = ?,
			pregnant_complications = ?,
			pregnant_drug_history = ?,
			updated_at = ?
		WHERE pregnant_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.MotherName,
		req.MotherBirthday,
		req.PregnantCongenitalDisease,
		req.PregnantComplications,
		req.PregnantDrugHistory,
		time.Now(),
		req.PregnantId,
	)
	if err != nil {
		return domain.UpdateMotherPregnantByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateMotherPregnantByIdResponse{}, nil
}
