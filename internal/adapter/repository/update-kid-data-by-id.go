package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) UpdateKidDataById(req domain.UpdateKidDataByIdRequest) (domain.UpdateKidDataByIdResponse, error) {
	query := `
		UPDATE kid_history SET
			kid_name = ?,
			kid_birthday = ?,
			kid_gender = ?,
			kid_gestational_age = ?,
			kid_oxygen = ?,
			kid_congenital_disease = ?,
			kid_birth_weight = ?,
			kid_body_length = ?,
			kid_blood_type = ?,
			updated_at = ?
		WHERE kid_id = ?
		AND pregnant_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.KidName,
		req.KidBirthday,
		req.KidGender,
		req.KidGestationalAge,
		req.KidOxygen,
		req.KidCongenitalDisease,
		req.KidBirthWeight,
		req.KidBodyLength,
		req.KidBloodType,
		time.Now(),
		req.KidId,
		req.PregnantId,
	)
	if err != nil {
		return domain.UpdateKidDataByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateKidDataByIdResponse{}, nil
}
