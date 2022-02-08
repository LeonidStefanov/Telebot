package database

import (
	"database/sql"
	"fmt"
	"home/leonid/Git/Pract/telegram_bot/pkg/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	base *sql.DB
}

func NewBD(host string) (*DB, error) {
	database, err := sql.Open("mysql", host)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	database.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &DB{
		base: database,
	}, nil

}

func (db *DB) Close() error {
	return db.base.Close()
}

func (db *DB) AddRequest(userId int, userName, response string) error {
	add := fmt.Sprintf("insert into telegram_bot(user_id, user_name, response) values(%d,'%v',%v)", userId, userName, response)

	_, err := db.base.Exec(add)
	if err != nil {
		return err
	}

	return nil
}
func (db *DB) GetRequest() ([]models.Request, error) {

	rep, err := db.base.Query("SELECT id,name,age FROM telegram_bot ")
	if err != nil {
		return nil, err
	}

	report := []models.Request{}

	for rep.Next() {
		var r models.Request
		err := rep.Scan(&r.ID, &r.UserID, &r.UserName, r.Respons)
		if err != nil {
			return nil, err
		}

		report = append(report, r)
	}
	if rep.Err() != nil {
		return nil, err
	}
	fmt.Println(report)
	return report, nil

}
