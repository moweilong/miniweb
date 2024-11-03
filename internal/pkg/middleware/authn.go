// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/moweilong/miniweb/internal/pkg/core"
	"github.com/moweilong/miniweb/internal/pkg/errno"
	"github.com/moweilong/miniweb/internal/pkg/known"
	"github.com/moweilong/miniweb/pkg/token"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(ctx)
		if err != nil {
			core.WriteResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}

		ctx.Set(known.XUsernameKey, username)
		ctx.Next()
	}
}
