package sakura_web

import "github.com/gin-gonic/gin"

func AuthGet(r *gin.Engine, relativePath string, handler gin.HandlerFunc) {
	r.GET(relativePath, JWTAuth(), handler)
}

func AuthPost(r *gin.Engine, relativePath string, handler gin.HandlerFunc) {
	r.POST(relativePath, JWTAuth(), handler)
}

func AuthDelete(r *gin.Engine, relativePath string, handler gin.HandlerFunc) {
	r.DELETE(relativePath, JWTAuth(), handler)
}

func AuthPut(r *gin.Engine, relativePath string, handler gin.HandlerFunc) {
	r.PUT(relativePath, JWTAuth(), handler)
}
