package repository

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) DeleteMotherPregnantById(req domain.DeleteMotherPregnantByIdRequest) (domain.DeleteMotherPregnantByIdResponse, error) {
	query := `DELETE FROM mother_pregnant_history WHERE pregnant_id = ?`
	_, err := r.db.Exec(query, req.PregnantId)
	if err != nil {
		return domain.DeleteMotherPregnantByIdResponse{}, err
	}

	return domain.DeleteMotherPregnantByIdResponse{}, nil
}
