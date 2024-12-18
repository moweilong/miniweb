// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moweilong/miniweb/internal/pkg/known"
)

// RequestID 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 保存在 gin.Context 中，方便后边程序使用
		c.Set(known.XRequestIDKey, requestID)
		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`
		c.Writer.Header().Add(known.XRequestIDKey, requestID)

		c.Next()
		// 请求方法处理后执行的逻辑
	}
}
