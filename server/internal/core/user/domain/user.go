package domain

import (
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type User struct {
	Id string
	Username string
	Active bool
	Title  string  		
	Avatar string    		
	Metadata metadata.Metadata  		
}

func (user User) Transform()(SanitizedUser, error) {	

	UUID, err := utils.ConvertStringToUUID(user.Id) 
	
	if err != nil {
		return SanitizedUser{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	validUsername, err := NewUsername(user.Username)

	if err != nil {
		return SanitizedUser{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidUsernameFormat, err)
	}

	validTitle, err := NewTitle(user.Title)

	if err != nil {
		return   SanitizedUser{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidTileFormat, err)
	}

	validAvtar, err := NewAvtar(user.Avatar)

	if err != nil {
		return  SanitizedUser{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidAvtarFormat, err)
	}

	return SanitizedUser{
		Id:           	UUID,
		Username: 		validUsername,
		Active:         user.Active,
		Title:    		validTitle,
		Avatar:    		validAvtar ,
		Metadata:  		user.Metadata,
	}, nil
}

type SanitizedUser struct {
	Id utils.UUID
	Username Username
	Active bool
	Title  Title  		
	Avatar Avtar    		
	Metadata metadata.Metadata    	
}

func (user SanitizedUser) Transform() (*User){
	return &User{
		Id: user.Id.String(),
		Username: user.Username.String(),
		Active: user.Active,
		Title: user.Title.String(),
		Avatar: user.Avatar.String(),
		Metadata: user.Metadata,
	}
}
