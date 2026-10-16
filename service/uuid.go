package service

import (
	"dsdCensus/dao"
	"dsdCensus/object"
	"log"

	"github.com/google/uuid"
)

type UUID interface {
	Save(discordId int) (*object.UUID, error)
	IfExists(object.UUID) (bool, error)
	FindByDiscordId(int) *object.UUID
}

type uuId struct {
	dao dao.UUID
}

func (u *uuId) Save(discordId int) (*object.UUID, error) {

	uuid := uuid.New()

	obj := &object.UUID{UUID: uuid.ID(), DiscordId: uint32(discordId)}

	err := u.dao.Save(*obj)
	if err != nil {
		log.Panicf("service save error %s\n", err)
		return nil, err
	}
	return obj, nil
}

func (u *uuId) IfExists(o object.UUID) (bool, error) {
	e, err := u.dao.IfExists(o)
	if err != nil {
		log.Panicf("uuid service IfExists error: %s\n", err)
		return false, err
	}

	return e, nil
}

func (u *uuId) FindByDiscordId(discord_id int) *object.UUID {
	uuid := u.dao.FindByDiscordId(discord_id)

	if uuid == nil {
		return nil
	}

	return uuid
}

func NewUUID(dao dao.UUID) UUID {
	return &uuId{dao}
}
