package list

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/list"
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

	query := fmt.Sprintf("SELECT COUNT(*) count FROM %s WHERE deleted_at is null", entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return 0, errors.WithStack(err)
	}
	defer stmt.Close()

	var res uint64
	if err = stmt.Get(&res); err != nil {
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

	query := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at is null ORDER BY id LIMIT ? OFFSET ?", entity.PrimaryKey, entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	var res []uint64
	if err = stmt.Select(&res, limit, offset); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return make([]uint64, 0), nil
		default:
			return nil, errors.WithStack(err)
		}
	}

	return res, nil
}

func Create(title string) (*entity.Entity, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("INSERT INTO %s (title, created_at, updated_at) VALUES (?, ?, ?)", entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	now := common.NewDatetime(conf.NowDatetime)
	var result sql.Result
	if result, err = stmt.Exec(title, now.String(), now.String()); err != nil {
		return nil, errors.WithStack(err)
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return nil, errors.WithStack(err)
	}

	res := &entity.Entity{
		ID:        uint64(id),
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return res, nil
}

func Update(id uint64, title string) (*entity.Entity, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res *entity.Entity
	if res, err = FindOne(id); err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("UPDATE %s set title = ?, updated_at = ? WHERE id = ?", entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	now := common.NewDatetime(conf.NowDatetime)
	if _, err = stmt.Exec(title, now.String(), id); err != nil {
		return nil, errors.WithStack(err)
	}

	res.Title = title
	res.UpdatedAt = now

	return res, nil
}

func Delete(id uint64) (*entity.Entity, error) {
	var err error

	conf := config.GetConfig()
	db, err := database.Open(conf.DatabaseSlave.URL, entity.Database)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res *entity.Entity
	if res, err = FindOne(id); err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("UPDATE %s set deleted_at = ? WHERE id = ?", entity.Table)

	var stmt *sqlx.Stmt
	if stmt, err = db.Preparex(query); err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	now := common.NewDatetime(conf.NowDatetime)
	if _, err = stmt.Exec(now.String(), id); err != nil {
		return nil, errors.WithStack(err)
	}

	res.DeletedAt = now

	return res, nil
}
