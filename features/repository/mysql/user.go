package mysql

import (
	"context"
	"errors"
	"online-learning-platform/domain"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMySQLUserRepository is constructor of MySQL repository
func NewMySQLUserRepository(Conn *sql.DB) domain.UserMySQLRepository {
	return &mysqlUserRepository{Conn}
}

func (db *mysqlUserRepository) SelectUserLogin(ctx context.Context, req domain.LoginRequest) (user domain.User, err error) {
	query := `SELECT id, email, password, name, picture, role, dtm_crt, dtm_upd FROM user WHERE email = ? AND deleted = 0`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, req.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Picture, &user.Role, &user.DtmCrt, &user.DtmUpd)
	if err != nil {
		err = errors.New("user not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) InsertUser(ctx context.Context, req domain.UserRequest) (id int64, err error) {
	query := `INSERT INTO user (email, password, name, picture, role, deleted, dtm_crt, dtm_upd) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.Email, req.Password, req.Name, req.Picture, req.Role, req.Deleted)
	if err != nil {
		err = errors.New("failed to create user")
		log.Error(err)
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) SelectUserByID(ctx context.Context, id int64) (user domain.User, err error) {
	query := `SELECT id, email, name, picture, role, dtm_crt, dtm_upd FROM user WHERE id = ? AND deleted = 0`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Picture, &user.Role, &user.DtmCrt, &user.DtmUpd)
	if err != nil {
		err = errors.New("user not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) RemoveUser(ctx context.Context, id int64) (err error) {
	query := `UPDATE user SET deleted = 1 WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		err = errors.New("failed to delete user")
		log.Error(err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}

	if affect == 0 {
		err = errors.New("not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) SelectTotalUser(ctx context.Context) (count int64, err error) {
	query := `SELECT COUNT(id) FROM user WHERE role = 'user'`
	log.Debug(query)

	row := db.Conn.QueryRowContext(ctx, query)
	err = row.Scan(&count)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
