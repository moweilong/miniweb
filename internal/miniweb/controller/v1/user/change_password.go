// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/moweilong/miniweb/internal/pkg/core"
	"github.com/moweilong/miniweb/internal/pkg/errno"
	"github.com/moweilong/miniweb/internal/pkg/log"
	v1 "github.com/moweilong/miniweb/pkg/api/miniweb/v1"
)

// ChangePassword 用来修改指定用户的密码.
func (c *UserController) ChangePassword(ctx *gin.Context) {
	log.C(ctx).Infow("Change password function called")

	var r v1.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(ctx, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := c.b.Users().ChangePassword(ctx, ctx.Param("name"), &r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
