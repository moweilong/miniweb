// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package model

import (
	"time"

	"github.com/moweilong/miniweb/pkg/auth"
	"gorm.io/gorm"
)

// UserM 是数据库中 user 记录 struct 格式的映射.
type UserM struct {
	CreatedAt time.Time `gorm:"column:createdAt"`         //
	Email     string    `gorm:"column:email"`             //
	ID        int64     `gorm:"column:id;primary_key"`    //
	Nickname  string    `gorm:"column:nickname"`          //
	Password  string    `gorm:"column:password;not null"` //
	Phone     string    `gorm:"column:phone"`             //
	UpdatedAt time.Time `gorm:"column:updatedAt"`         //
	Username  string    `gorm:"column:username;not null"` //
}

// TableName 用来指定映射的 MySQL 表名.
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}