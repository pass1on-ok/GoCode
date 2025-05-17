package mysql

import (
	"context"
	"errors"
	"online-learning-platform/domain"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlTransactionRepository struct {
	Conn *sql.DB
}

// NewMySQLTransactionRepository is constructor of MySQL repository
func NewMySQLTransactionRepository(Conn *sql.DB) domain.TransactionMySQLRepository {
	return &mysqlTransactionRepository{Conn}
}

func (db *mysqlTransactionRepository) InsertTransaction(ctx context.Context, req domain.TransactionRequest) (id int64, err error) {
	query := `INSERT INTO transaction (user_id, course_id, price, dtm_crt, dtm_upd) VALUES (?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.UserID, req.CourseID, req.Price)
	if err != nil {
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

func (db *mysqlTransactionRepository) SelectTransactionByID(ctx context.Context, id int64) (transaction domain.Transaction, err error) {
	query := `SELECT id, user_id, course_id, price, dtm_crt, dtm_upd FROM transaction WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&transaction.ID, &transaction.UserID, &transaction.CourseID, &transaction.Price, &transaction.DtmCrt, &transaction.DtmUpd)
	if err != nil {
		err = errors.New("transaction not found")
		log.Error(err)
		return
	}

	return
}
