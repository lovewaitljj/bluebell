package models

import "time"

type User struct {
	ID         uint64    `gorm:"primaryKey;column:id;autoIncrement"`                                                      // 主键 ID
	UserID     uint64    `gorm:"column:user_id;not null;uniqueIndex:idx_user_id"`                                         // 用户唯一 ID
	Username   string    `gorm:"column:username;type:varchar(64);not null;uniqueIndex"`                                   // 用户名，唯一
	Password   string    `gorm:"column:password;type:varchar(64);not null"`                                               // 密码
	Email      string    `gorm:"column:email;type:varchar(64);"`                                                          // 邮箱
	Gender     int8      `gorm:"column:gender;type:tinyint;default:0;not null"`                                           // 性别，0: 未指定, 1: 男, 2: 女
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"` // 更新时间
}

func (u User) TableName() string {
	return "user"
}
