package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
}
