package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"` // 密码不返回给前端
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
}
