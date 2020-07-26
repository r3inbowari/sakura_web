package rest

import (
	"github.com/gin-gonic/gin"
	"sakura_web/domain"
)

func User(r *gin.Engine)  {
	r.GET("v1/rank", domain.Rank)
}