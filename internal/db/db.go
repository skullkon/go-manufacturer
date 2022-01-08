package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/skullkon/go-manufacturer/internal/models"
	"strconv"
)

type Database struct {
	DB *sqlx.DB
}

func NewDB(dsn string) (*Database, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
	}, nil
}

func (db *Database) GetPosts(N int) ([]models.Manufacturer, error) {
	var manyPosts []models.Manufacturer
	i := 0

	result, err := db.DB.Query("SELECT id, details FROM manufacturer WHERE details->>'needUpdate' = 'true';")
	if err != nil {
		return nil, err
	}

	for result.Next() {
		if i >= N {
			break
		}
		i++
		//один большой костыль
		type locDetails struct {
			NeedUpdate string `json:"needUpdate"`
		}
		var locData locDetails
		var locDetailed models.Detail
		var post models.Manufacturer
		var b types.JSONText

		err := result.Scan(&post.Id, &b)
		if err != nil {
			return nil, err
		}
		err = b.Unmarshal(&locData)
		if err != nil {
			return nil, err
		}
		booledStr, err := strconv.ParseBool(locData.NeedUpdate)
		if err != nil {
			return nil, err
		}
		locDetailed.NeedUpdate = booledStr
		manyPosts = append(manyPosts, post)

	}

	return manyPosts, nil
}

func (db *Database) Update(id []int64) error {
	_, err := db.DB.Exec("UPDATE manufacturer SET details = jsonb_set(details, '{needUpdate}', '\"false\"') WHERE id = any($1);", pq.Array(id))
	if err != nil {
		return err
	}

	return nil
}
