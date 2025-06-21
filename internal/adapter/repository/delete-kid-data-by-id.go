package repository

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (r *repository) DeleteKidDataById(req domain.DeleteKidDataByIdRequest) (domain.DeleteKidDataByIdResponse, error) {
	query := `DELETE FROM kid_history WHERE kid_id = ?`
	_, err := r.db.Exec(query, req.KidId)
	if err != nil {
		return domain.DeleteKidDataByIdResponse{}, err
	}

	return domain.DeleteKidDataByIdResponse{}, nil
}
