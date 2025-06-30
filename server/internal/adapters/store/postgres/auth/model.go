package auth

import (
	"database/sql"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/utils"
)

type AuthModel struct {
	Id 		 		utils.UUID			`db:"auth_id"`
	Email    		domain.Email		`db:"email"`
	Password 		domain.Password		`db:"password"`
	Version  		int					`db:"version"`
	MajicActive  	bool				`db:"majic_active"`
	OTPActive 		bool				`db:"otp_active"`
	OIDCActive		bool				`db:"oidc_active"`
	PasswordActive  bool				`db:"password_active"`
	UpdatedBy 		sql.NullString 		`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		sql.NullString 		`db:"created_by"`
	CreatedDT 		sql.NullTime  		`db:"created_dt"`
}

func (s AuthModel) Transform() domain.SanitizedAuth {
	return domain.SanitizedAuth{
		Id: 		s.Id,
		Email: 		s.Email,		
		Password: 	s.Password,
		Version: 	s.Version,
		MajicActive: s.MajicActive,			
		OTPActive: s.OTPActive, 					
		OIDCActive: s.OIDCActive,						
		PasswordActive: s.PasswordActive,  	
	}
}