package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

func (r *repository) UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResult, error) {
	var result domain.UserRequestResetPasswordResult

	err := r.db.QueryRow(
		`SELECT 
			CASE 
				WHEN EXISTS(SELECT 1 FROM user WHERE email = ?) 
				THEN (SELECT user_id FROM user WHERE email = ? LIMIT 1) 
				ELSE NULL 
			END AS user_id;
		`,
		req.Email,
		req.Email,
	).Scan(&result.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("not found this email")
		}
		return result, err
	}

	_, err = r.db.Exec(
		`UPDATE user SET reset_password_token = ? WHERE user_id = ?`,
		req.Token,
		result.UserId,
	)
	if err != nil {
		return result, fmt.Errorf("failed to update user reset password token: %w", err)
	}

	return result, nil
}
