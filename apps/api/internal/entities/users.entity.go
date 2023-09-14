package entities

type User struct {
	Model    `swag:"-"`
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=30"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"omitempty"`
	Bio      string `validate:"max=250"`
	Active   bool   `swag:"-"`
}
