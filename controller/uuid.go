package controller

import (
	"dsdCensus/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UIUD interface {
	Save(c *gin.Context)
	FindByDiscordId(c *gin.Context)
}

type uuId struct {
	service service.UUID
}

func (u *uuId) Save(c *gin.Context) {
	var discordId struct {
		DiscordId string `json:"discord_id"`
	}

	c.BindJSON(&discordId)

	discordIdInt, err := strconv.Atoi(discordId.DiscordId)

	if err != nil {
		log.Printf("controller save error: %s\n", err)
		c.Status(400)
		return
	}

	object, err := u.service.Save(discordIdInt)

	if err != nil {
		log.Printf("controller save error: %s\n: ", err)
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{"uuid": object.UUID, "discord_id": object.DiscordId})
}

func (u *uuId) FindByDiscordId(c *gin.Context) {
	discordId := c.Param("discordId")

	discordIdInt, err := strconv.Atoi(discordId)

	if err != nil {
		log.Printf("controller find by discord id error: %s\n", err)
		c.Status(400)
		return
	}

	object := u.service.FindByDiscordId(discordIdInt)

	if object == nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{ "uuid": object.UUID, "discord_id": object.DiscordId })
}

func NewUUID(service service.UUID) UIUD {
	return &uuId{service}
}
