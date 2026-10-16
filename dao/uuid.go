package dao

import (
	"dsdCensus/base/db"
	object "dsdCensus/object"
	"log"
)

type UUID interface {
	Save(object.UUID) error
	IfExists(object.UUID) (bool, error)
	FindByDiscordId(int) *object.UUID
}

type uuId struct {
	db db.DAO
}

func (u *uuId) Save(o object.UUID) error {
	_, err := u.db.Exec("INSERT INTO uuid ( uuid, discord_id ) values (?, ?)", o.UUID, o.DiscordId)
	if err != nil {
		log.Printf("uuid save error: %s\n", err)
		return err
	}

	return nil
}

func (u *uuId) IfExists(o object.UUID) (bool, error) {
	rows, err := u.db.Query("SELECT COUNT(*) FROM uuid WHERE uuid=? AND discord_id=?", o.UUID, o.DiscordId);

	if err != nil {
		log.Printf("uuid ifexists error: %s\n", err)
		return false, err
	}

	var count struct {
		Count int
	}

	rows.Next()
	err = rows.Scan(&count.Count)

	if err != nil {
		log.Printf("uuid dao if exists error: %s\n", err)
		return false, err
	}

	if count.Count != 0 {
		return false, nil
	}
	return true, nil

}

func (u *uuId) FindByDiscordId(discord_id int) (*object.UUID) {
	row := u.db.QueryRow("SELECT * from uuid WHERE discord_id=?", discord_id)

	if row == nil {
		return nil
	}

	var uuid object.UUID

	err := row.Scan(&uuid.UUID, &uuid.DiscordId)

	if err != nil {
		log.Printf("uuid dao scan error: %s\n", err)
		return nil
	}

	return &uuid
}

func NewUuId(db db.DAO) UUID {
	return &uuId{db};
}
