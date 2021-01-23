package app

import (
	"log"
	"strconv"

	cr "github.com/brkelkar/common_utils/configreader"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Start_Application(cfg cr.Config) {
	port := strconv.Itoa(cfg.Server.Port)
	host := cfg.Server.Host
	serverUrl := host + ":" + port

	mapUrls()

	log.Printf("connect to http://%s/ for GraphQL playground", serverUrl)
	log.Fatal(router.Run(serverUrl))
}
