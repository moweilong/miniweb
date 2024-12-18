// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package user

import (
	"github.com/moweilong/miniweb/internal/miniweb/biz"
	"github.com/moweilong/miniweb/internal/miniweb/store"
	"github.com/moweilong/miniweb/pkg/auth"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{
		a: a,
		b: biz.NewBiz(ds),
	}
}
