package models

import (
	"time"
)

type Users struct {
	Model
	Name     string     `json:"name"`
	Age      int64      `json:"age"`
	Birthday *time.Time `json:"birthday"`
	Email    string     `json:"email" gorm:"type:varchar(100);unique_index"`
	Role     string     `json:"role" gorm:"size:255"`      // 设置字段大小为255
	Num      int        `json:"num" gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address  string     `json:"address" gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe int        `json:"ignore_me" gorm:"-"`        // 忽略本字段
}

func (Users) TableName() string {
	// 自定义表名
	return "profiles"
}
