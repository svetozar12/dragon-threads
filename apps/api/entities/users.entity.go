package entities

type User struct {
	Model       `swag:"-"`
	Username    string `json:"username" gorm:"unique;not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Active      bool   `json:"active"`
	SubDragonId int32  `json:"subDragonId"`
}
