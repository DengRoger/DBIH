package controller

import (
	"DBEE_HW/DB"
	"DBEE_HW/module"
	"DBEE_HW/encryption"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func RouterController() {
	router := gin.Default()
	// router.GET("/page", GetHead)
	// router.GET("/list", GetLists)
	// router.POST("/post", Post)
	router.Run(":8443")
}
