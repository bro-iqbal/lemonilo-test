package models

type User struct {
	UserID   int    `json:"UserID" gorm:"column:userid;primaryKey;autoIncrement"`
	Username string `json:"Username" gorm:"column:username"`
	Email    string `json:"Email" gorm:"column:email"`
	Address  string `json:"Address" gorm:"column:address"`
	Password string `json:"Password" gorm:"column:password"`
	Token    string `json:"Token" gorm:"column:token"`
}

type Users []User

func (User) TableName() string {
	return "user"
}
