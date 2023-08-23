package model

type Person struct {
	ID          int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
