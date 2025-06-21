package repository

import (
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) GetKidData(req domain.GetKidDataRequest) ([]domain.GetKidDataResponse, error) {
	var result []domain.GetKidDataResponse

	query := `
		SELECT 
			kid.kid_id,
			pregnant.pregnant_id,
			pregnant.mother_name,
			kid.kid_name,
			kid.kid_birthday,
			kid.kid_gender,
			kid.kid_gestational_age,
			kid.kid_oxygen,
			kid.kid_birth_weight,
			kid.kid_body_length,
			kid.kid_blood_type,
			kid.kid_congenital_disease,
			kid.created_at,
			kid.updated_at
		FROM kid_history kid
		LEFT JOIN mother_pregnant_history pregnant
        	ON kid.pregnant_id = pregnant.pregnant_id
        LEFT JOIN user
        	ON pregnant.user_id = user.user_id
            AND user.user_id = ?
		WHERE kid.pregnant_id = ?
	`
	rows, err := r.db.Query(query, req.UserId, req.PregnantId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var kid domain.GetKidDataResponse
		if err := rows.Scan(
			&kid.KidId,
			&kid.PregnantId,
			&kid.MotherName,
			&kid.KidName,
			&kid.KidBirthday,
			&kid.KidGender,
			&kid.KidGestationalAge,
			&kid.KidOxygen,
			&kid.KidBirthWeight,
			&kid.KidBodyLength,
			&kid.KidBloodType,
			&kid.KidCongenitalDisease,
			&kid.CreatedAt,
			&kid.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}

		result = append(result, kid)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
