package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:root@tcp(tasker_dev_db)/tasker_dev")

type User struct {
	ID       uint64
	Email    string
	Password string
}

func selectUserList() {

}

func FindById(id uint64) (*User, error) {
	sql := "select id, email from users where id = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := User{}
	if err := stmt.QueryRow(id).Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
