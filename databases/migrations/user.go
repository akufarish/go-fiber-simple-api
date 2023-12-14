package migrations

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(255)"  json:"email" form:"email"`
	Username string `gorm:"type:varchar(255)" json:"username" form:"username"`
	Password string `gorm:"type:varchar(255)" json:"password" form:"password"`
}