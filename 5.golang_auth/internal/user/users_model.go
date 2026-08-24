package user

import "time"

// type Role string

// const (
// 	RoleUser  Role = "user"
// 	RoleAdmin Role = "admin"
// )

type User struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	Email          string    `json:"email" bson:"email"`
	HashedPassword string    `json:"-" bson:"hashed_password"`
	Role           string    `json:"role" bson:"role"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
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


