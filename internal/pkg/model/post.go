// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package model

import "time"

// PostM 是数据库中 post 记录 struct 格式的映射.
type PostM struct {
	Content   string    `gorm:"column:content"`           //
	CreatedAt time.Time `gorm:"column:createdAt"`         //
	ID        int64     `gorm:"column:id;primary_key"`    //
	PostID    string    `gorm:"column:postID;not null"`   //
	Title     string    `gorm:"column:title;not null"`    //
	UpdatedAt time.Time `gorm:"column:updatedAt"`         //
	Username  string    `gorm:"column:username;not null"` //
}

// TableName 用来指定映射的 MySQL 表名.
func (p *PostM) TableName() string {
	return "post"
}
