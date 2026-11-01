package user

import "time"

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	Role           string    `json:"role"`
	RefreshToken   string    `json:"-"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type UserRecord struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	Email          string `gorm:"type:varchar(255);not null;uniqueIndex"`
	HashedPassword string `gorm:"type:text;not null"`
	Role           string `gorm:"type:varchar(50);not null;default:user"`
	RefreshToken   string `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (UserRecord) TableName() string {
	return "users"
}

type PublicUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) ToPublicUser() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
