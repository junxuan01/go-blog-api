package model

type Article struct {
	BaseModel
	Title   string `gorm:"type:varchar(255);not null;index" json:"title"` // 标题加索引，方便搜索
	Content string `gorm:"type:longtext" json:"content"`
	UserID  uint   `gorm:"index;not null" json:"user_id"`           // 逻辑外键
	User    *User  `gorm:"foreignKey:UserID" json:"user,omitempty"` // 关联关系
}
