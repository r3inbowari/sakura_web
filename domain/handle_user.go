package domain

import (
	"github.com/gin-gonic/gin"
	Sakura "github.com/r3inbowari/sakura_go"
)

func Post(c *gin.Context) {
	ret := Sakura.RankUpdate()

	c.JSON(200, gin.H{
		"data": ret,
	})
}
