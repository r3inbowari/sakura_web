package main

import (
	"github.com/gin-gonic/gin"
	"sakura_web"
	"sakura_web/rest"
	"sakura_web/utils"
)

func main() {
	r := gin.Default()
	r.Use(sakura_web.Cors())
	rest.User(r)

	_ = r.Run(utils.GetConfig().ServerAddr)
}



