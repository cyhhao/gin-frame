package models

import (
	"database/sql"
	"time"
)

type Users struct {
	Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     string `gorm:"size:255"`       // 设置字段大小为255
	Num      int    `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address  string `gorm:"index:addr"`     // 给address字段创建名为addr的索引
	IgnoreMe int    `gorm:"-"`              // 忽略本字段
}

func (Users) TableName() string {
	// 自定义表名
	return "profiles"
}
