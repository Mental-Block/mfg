package domain

import "github.com/server/pkg/utils"

type Auth struct {
	Id 		 		string			
	Email    		string		
	Password 		string	
	MajicActive		bool
	OtpActive		bool
	OIDCActive		bool
	PasswordActive 	bool
	Version  		int	
}

func (auth Auth) Transform() (SanitizedAuth, error) {
	
	UUID, err := utils.ConvertStringToUUID(auth.Id) 
	
	if err != nil {
		return SanitizedAuth{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	validEmail, err := Email(auth.Email).NewEmail()

	if err != nil {
		return SanitizedAuth{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "invalid email supplied: %s", err)
	}

	validPassword, err := Password(auth.Password).NewPassword()

	if err != nil {
		return SanitizedAuth{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "invalid password supplied: %s", err)
	}

	return SanitizedAuth{
		Id: 		UUID,
		Email: 		validEmail,		
		Password: 	validPassword,
		Version: 	auth.Version,
		MajicActive: auth.MajicActive,		
		OTPActive: auth.OtpActive,		
		OIDCActive: auth.OIDCActive,		
		PasswordActive: auth.PasswordActive, 	
	}, nil
}

type SanitizedAuth struct {
	Id 		 		utils.UUID				
	Email    		Email		
	Password 		Password
	MajicActive		bool
	OTPActive		bool
	OIDCActive		bool
	PasswordActive 	bool		
	Version  	    int		
}

func (s SanitizedAuth) Transform() *Auth {
	return &Auth{
		Id: 		s.Id.String(),
		Email: 		s.Email.String(),		
		Password: 	s.Password.String(),
		Version: 	s.Version,
		MajicActive: s.MajicActive,		
		OtpActive: s.OTPActive,		
		OIDCActive: s.OIDCActive,		
		PasswordActive: s.PasswordActive,
	}
}