package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	query := `
		INSERT INTO user (
			user_id,
			email, 
			password, 
			created_at
		) VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.UserId,
		req.Email,
		req.Password,
		time.Now(),
	)
	if err != nil {
		return domain.UserRegisterResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	return domain.UserRegisterResponse{}, nil
}
