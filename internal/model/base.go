package model

import (
	"time"

	"gorm.io/gorm"
)

// model 包用于定义 GORM 数据库模型，此处先占位，后续可添加 User、Post 等结构体。

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除字段,前端不可见
}
