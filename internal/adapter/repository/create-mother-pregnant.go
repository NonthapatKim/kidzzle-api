package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) CreateMotherPregnant(req domain.CreateMotherPregnantRequest) (domain.CreateMotherPregnantResponse, error) {
	query := `
		INSERT INTO mother_pregnant_history (
			user_id,
			mother_name,
			mother_birthday, 
			pregnant_congenital_disease, 
			pregnant_complications,
			pregnant_drug_history
		) VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.UserId,
		req.MotherName,
		req.MotherBirthday,
		req.PregnantCongenitalDisease,
		req.PregnantComplications,
		req.PregnantDrugHistory,
	)
	if err != nil {
		return domain.CreateMotherPregnantResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.CreateMotherPregnantResponse{}, nil
}
