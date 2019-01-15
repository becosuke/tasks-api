package task

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/task"
	"github.com/becosuke/tasks-api/infrastructure/database"
)

func FindOne(id uint64) (*entity.Entity, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", entity.Table, entity.PrimaryKey)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	res := &entity.Entity{}
	if err = stmt.Get(res, id); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, nil
		default:
			return nil, errors.WithStack(err)
		}
	}

	return res, nil
}

func CountAll() (uint64, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	query := fmt.Sprintf("SELECT COUNT(*) count FROM %s WHERE deleted = ?", entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return 0, errors.WithStack(err)
	}
	defer stmt.Close()

	var res uint64
	if err = stmt.Get(&res, common.DELETED_STATUS_ALIVE); err != nil {
		return 0, errors.WithStack(err)
	}

	return res, nil
}

func FindPrimaryKeyAll(limit int32, offset int32) ([]uint64, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE deleted = ? ORDER BY id LIMIT ? OFFSET ?", entity.PrimaryKey, entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	var res []uint64
	if err = stmt.Select(&res, common.DELETED_STATUS_ALIVE, limit, offset); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return make([]uint64, 0), nil
		default:
			return nil, errors.WithStack(err)
		}
	}

	return res, nil
}
