// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package miniweb

import (
	"github.com/gin-gonic/gin"
	"github.com/moweilong/miniweb/internal/miniweb/controller/v1/user"
	"github.com/moweilong/miniweb/internal/miniweb/store"
	"github.com/moweilong/miniweb/internal/pkg/core"
	"github.com/moweilong/miniweb/internal/pkg/errno"
	"github.com/moweilong/miniweb/internal/pkg/log"
	mw "github.com/moweilong/miniweb/internal/pkg/middleware"
	"github.com/moweilong/miniweb/pkg/auth"
)

// installRouters 安装 miniweb 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 无需认证和授权
			userv1.PUT(":name/change-password", uc.ChangePassword) // 已经有认证，无需授权
			userv1.Use(mw.Authn(), mw.Authz(authz))                // 后面的接口需要认证和授权
			userv1.GET(":name", uc.Get)                            // 获取用户详情
		}
	}

	return nil
}
