package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.StaticFS("/public", http.Dir("K:/practice/src/gin/web/static"))
	router.StaticFile("/favicon.ico", "web/favicon.ico")

	_ = router.Run(":80")
}
