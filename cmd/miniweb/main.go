// Copyright 2024 Noah Long(莫维龙) <kalandramo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/moweilong/miniweb.

package main

import (
	"os"

	"github.com/moweilong/miniweb/internal/miniweb"
	_ "go.uber.org/automaxprocs"
)

func main() {
	command := miniweb.NewMiniWebCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
