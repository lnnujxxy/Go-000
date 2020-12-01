package dao

import (
	"database/sql"
	"fmt"
	"gogeekbang/internal/pkg/config"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup(config config.Database) {
	var err error
	db, err = sql.Open(config.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Host, config.Database))
	if err != nil {
		panic("dao: init fail err " + err.Error())
	}
}

func GetName(id int) (name string, err error) {
	err = db.QueryRow("SELECT name FROM test WHERE id = ?", id).Scan(&name)
	// 没有记录不作为错误返回
	if err != nil && err != sql.ErrNoRows {
		return name, errors.Wrap(err, fmt.Sprintf("dao: GetName id=%d", id))
	}
	return name, nil
}
