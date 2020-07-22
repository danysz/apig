package server

import (
	"github.com/danysz/apig/_example/middleware"
	"github.com/danysz/apig/_example/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
