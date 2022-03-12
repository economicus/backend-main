package table

type Profile struct {
	UserID       uint     `gorm:"primaryKey" json:"user_id"`
	Nickname     string   `gorm:"column:nickname;not null" json:"nickname,omitempty"`
	ProfileImage string   `gorm:"column:profile_image;default:'photo/no_image.png'" json:"profile_image,omitempty"`
	Email        string   `gorm:"column:email" json:"email,omitempty"`
	Phone        string   `gorm:"column:phone" json:"phone,omitempty"`
	UserURL      string   `gorm:"column:user_url" json:"user_url,omitempty"`
	IntroMessage string   `gorm:"column:intro_message" json:"intro_message"`
	Location     Location `gorm:"embedded;embeddedPrefix:location_" json:"location,omitempty"`
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
