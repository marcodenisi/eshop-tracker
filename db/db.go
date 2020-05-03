package db

import (
	"bytes"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/marcodenisi/eshop-tracker/model"
)

const dbName = "games.db"
const bucketName = "games"

// SaveGames saves games to DB
func SaveGames(games []model.EuGame) error {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal("Error while opening database", err)
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal("Error while opening games bucket", err)
			return err
		}

		for _, game := range games {
			obj, err := model.EncodeEuGame(game)
			if err != nil {
				continue
			}
			b.Put([]byte(strings.ToLower(game.Title)), obj)
		}

		return nil
	})
	return err
}

// GetGames retrieves games from database
func GetGames() ([]model.EuGame, error) {
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		log.Fatal("Error while opening database", err)
		return nil, err
	}
	defer db.Close()

	games := []model.EuGame{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(_, v []byte) error {
			g, err := model.DecodeEuGame(v)
			if err != nil {
				return err
			}
			games = append(games, *g)
			return nil
		})
		return nil
	})

	return games, err
}

// GetGamesFromName retrieves a list of games beginning with the provided name
func GetGamesFromName(name string) ([]model.EuGame, error) {
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		log.Fatal("Error while opening database", err)
		return nil, err
	}
	defer db.Close()

	games := []model.EuGame{}
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucketName)).Cursor()

		prefix := []byte(strings.ToLower(name))
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			g, err := model.DecodeEuGame(v)
			if err != nil {
				return err
			}
			games = append(games, *g)
		}

		return nil
	})

	return games, err
}
