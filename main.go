package main

import (
	appconfig "dsdCensus/base/config"
	"dsdCensus/base/db"
	"dsdCensus/controller"
	"dsdCensus/dao"
	"dsdCensus/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg, err := appconfig.NewConfig("conf/conf.yml")

	if err != nil {
		log.Panicln(err)
		return
	}

	sqlite := db.Init(cfg)

	uuidDao := dao.NewUuId(sqlite)

	uuidService := service.NewUUID(uuidDao)

	uuidController := controller.NewUUID(uuidService)

	r.GET("/uuid/:discordId", uuidController.FindByDiscordId)
	r.POST("/uuid", uuidController.Save)

	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
