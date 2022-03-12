package model

import (
	"main/internal/core/model/table"
	"main/internal/pkg/objconv"
)

type Profile struct {
	table.Profile
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func NewProfile(userID uint, nickname string) *Profile {
	return &Profile{
		table.Profile{
			UserID:   userID,
			Nickname: nickname,
		},
	}
}

func (p *Profile) GetID() uint {
	return p.UserID
}

func (p *Profile) GetOwnerID() uint {
	return p.UserID
}

func (p *Profile) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"user_id":       p.UserID,
		"nickname":      p.Nickname,
		"profile_image": p.ProfileImage,
		"email":         p.Email,
		"user_url":      p.UserURL,
		"intro_message": p.IntroMessage,
		"location": map[string]string{
			"country": p.Location.Country,
			"city":    p.Location.City,
		},
	}
	return data
}

func (p *Profile) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(p, fields)
}
