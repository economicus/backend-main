package table

import (
	"gorm.io/gorm"
	"time"
)

type Users []*User

type User struct {
	gorm.Model     `json:"-" swaggerignore:"true"`
	Email          string    `gorm:"column:email;not null;unique" json:"email,omitempty"`
	Password       []byte    `gorm:"column:password;not null" json:"-"`
	UserActive     bool      `gorm:"column:user_active;default:true;" json:"-"`
	AccessLevel    int       `gorm:"column:access_level;default:1" json:"-"`
	LastLogin      time.Time `gorm:"column:last_login;default:null" json:"-"`
	AuthResource   string    `gorm:"column:auth_resource;default:'local'" json:"-"`
	Profile        Profile   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"profile,omitempty"`
	Quants         []Quant   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	FavoriteQuants []*Quant  `gorm:"many2many:user_favorite_quants;" json:"-"`
	Comments       []Comment `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Replies        []Reply   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Followings     Users     `gorm:"many2many:followings;" json:"-"`
}
