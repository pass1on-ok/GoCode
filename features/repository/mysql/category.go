package mysql

import (
	"context"
	"errors"
	"online-learning-platform/domain"
	"strconv"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

// NewMySQLCategoryRepository is constructor of MySQL repository
func NewMySQLCategoryRepository(Conn *sql.DB) domain.CategoryMySQLRepository {
	return &mysqlCategoryRepository{Conn}
}

func (db *mysqlCategoryRepository) InsertCategory(ctx context.Context, name string) (err error) {
	query := `INSERT INTO category (name, dtm_crt, dtm_upd) VALUES (?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		err = errors.New("failed to create category")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCategoryRepository) SelectCategoryByID(ctx context.Context, id int64) (category domain.Category, err error) {
	query := `SELECT id, name, dtm_crt, dtm_upd FROM category WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&category.ID, &category.Name, &category.DtmCrt, &category.DtmUpd)
	if err != nil {
		err = errors.New("no category found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCategoryRepository) SelectAllCategory(ctx context.Context, page, limit int64, sort string) (category []domain.Category, err error) {
	var query string
	if sort == "popular" || sort == "unpopular" {
		query = `SELECT c.name, COUNT(co.category_id) AS most_popular FROM transaction t
		INNER JOIN course co ON t.course_id = co.id
		INNER JOIN category c ON co.category_id = c.id
		GROUP BY c.name ORDER BY most_popular`

		if sort != "" {
			if sort == "popular" {
				query += ` ASC`
			} else if sort == "unpopular" {
				query += ` DESC`
			} else {
				query += ` ASC`
			}
		}

		if limit > 0 {
			query += ` LIMIT ` + strconv.Itoa(int(limit))
		}

		if page > 0 {
			query += ` OFFSET ` + strconv.Itoa(int(page))
		}

		log.Debug(query)

		rows, er := db.Conn.QueryContext(ctx, query)
		if er != nil {
			if er == sql.ErrNoRows {
				err = errors.New("category not found")
			}
		}
		defer rows.Close()

		for rows.Next() {
			var i domain.Category
			err = rows.Scan(&i.Name, &i.Count)
			if err != nil {
				log.Error(err)
				return
			}

			category = append(category, i)
		}

	} else {
		query = `SELECT id, name, dtm_crt, dtm_upd FROM category`

		if limit > 0 {
			query += ` LIMIT ` + strconv.Itoa(int(limit))
		}

		if page > 0 {
			query += ` OFFSET ` + strconv.Itoa(int(page))
		}

		log.Debug(query)

		rows, er := db.Conn.QueryContext(ctx, query)
		if er != nil {
			if er == sql.ErrNoRows {
				err = errors.New("category not found")
			}
		}
		defer rows.Close()

		for rows.Next() {
			var i domain.Category
			err = rows.Scan(&i.ID, &i.Name, &i.DtmCrt, &i.DtmUpd)
			if err != nil {
				log.Error(err)
				return
			}

			category = append(category, i)
		}
	}

	return
}

func (db *mysqlCategoryRepository) EditCategory(ctx context.Context, name string, id int64) (err error) {
	query := `UPDATE category SET name = ? WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Debug(err)
		return
	}

	res, err := stmt.ExecContext(ctx, name, id)
	if err != nil {
		err = errors.New("failed to update")
		log.Debug(err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Debug(err)
		return
	}

	if affect == 0 {
		err = errors.New("no updated data")
		return
	}

	return
}
