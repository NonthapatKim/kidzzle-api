package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) CreateKidData(req domain.CreateKidDataRequest) (domain.CreateKidDataResponse, error) {
	query := `
		INSERT INTO kid_history (
			pregnant_id,
			kid_name,
			kid_birthday,
			kid_gender,
			kid_gestational_age,
			kid_oxygen,
			kid_congenital_disease,
			kid_birth_weight,
			kid_body_length,
			kid_blood_type
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.PregnantId,
		req.KidName,
		req.KidBirthday,
		req.KidGender,
		req.KidGestationalAge,
		req.KidOxygen,
		req.KidCongenitalDisease,
		req.KidBirthWeight,
		req.KidBodyLength,
		req.KidBloodType,
	)
	if err != nil {
		return domain.CreateKidDataResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.CreateKidDataResponse{}, nil
}
