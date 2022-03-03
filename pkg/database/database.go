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

func (db *DB) AddRequest(userId int, userName, response string, resTime string) error {

	add := fmt.Sprintf("insert into telegram_bot(user_id, user_name, response, date_and_time ) values(%d,'%v','%v','%v')", userId, userName, response, resTime)

	_, err := db.base.Exec(add)
	if err != nil {
		return err
	}

	return nil
}
func (db *DB) GetRequest() ([]models.Request, error) {

	rep, err := db.base.Query("SELECT id,user_id ,user_name,response, date_and_time FROM telegram_bot ")
	if err != nil {
		return nil, err
	}

	report := []models.Request{}

	for rep.Next() {
		var r models.Request
		err := rep.Scan(&r.ID, &r.UserID, &r.UserName, &r.Respons, &r.ResTime)
		if err != nil {
			return nil, err
		}

		report = append(report, r)
	}
	if rep.Err() != nil {
		return nil, err
	}

	return report, nil

}

func (db *DB) GetUsers() ([]models.Request, error) {

	rep, err := db.base.Query("SELECT user_id , user_name  from  telegram_bot   ")
	if err != nil {
		return nil, err
	}

	report := []models.Request{}

	for rep.Next() {
		var r models.Request
		err := rep.Scan(&r.UserID, &r.UserName)
		if err != nil {
			return nil, err
		}

		report = append(report, r)
	}
	if rep.Err() != nil {
		return nil, err
	}

	return report, nil

}

func (db *DB) GetUserRequests(user_name string) ([]models.Request, error) {

	rep, err := db.base.Query("Select response,date_and_time from telegram_bot WHERE user_name =?", user_name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	report := []models.Request{}

	for rep.Next() {
		var r models.Request
		err := rep.Scan(&r.Respons, &r.ResTime)
		if err != nil {
			return nil, err
		}

		fmt.Println(r)
		report = append(report, r)
	}
	if rep.Err() != nil {
		return nil, err
	}

	return report, nil
}

func (db *DB) DeleteRequst(id int) error {
	str := fmt.Sprintf("DELETE FROM telegram_bot WHERE id =%v", id)

	add, err := db.base.Exec(str)
	if err != nil {
		return err
	}
	fmt.Println(add)
	return nil
}
