package database

import (
	"fmt"
	"math/rand"
	"sync"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var dbs = map[string]*sqlx.DB{}

func Open(seed []string, dbname string) (*sqlx.DB, error) {
	var err error

	var url string
	if len(seed) > 1 {
		url = fmt.Sprintf(seed[rand.Intn(len(seed))], dbname)
	} else {
		url = fmt.Sprintf(seed[0], dbname)
	}

	var mutex sync.Mutex
	mutex.Lock()
	db, ok := dbs[url]
	if ok == false {
		db, err = sqlx.Open("mysql", url)
		if err != nil {
			mutex.Unlock()
			return nil, errors.WithStack(err)
		}
		dbs[url] = db
	}
	mutex.Unlock()

	return db, nil
}
