package main

import (
	_ "WorthRecord-gf/boot"
	_ "WorthRecord-gf/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
