package model

type Comment struct {
	BaseModel
	Content   string `gorm:"type:text;not null" json:"content"`
	ArticleID uint   `gorm:"index;not null" json:"article_id"`
	UserID    uint   `gorm:"index;not null" json:"user_id"`
}
