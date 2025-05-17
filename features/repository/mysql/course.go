package mysql

import (
	"context"
	"errors"
	"online-learning-platform/domain"
	"strings"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlCourseRepository struct {
	Conn *sql.DB
}

// NewMySQLCourseRepository is constructor of MySQL repository
func NewMySQLCourseRepository(Conn *sql.DB) domain.CourseMySQLRepository {
	return &mysqlCourseRepository{Conn}
}

func (db *mysqlCourseRepository) InsertCourse(ctx context.Context, req domain.CourseRequest) (id int64, err error) {
	query := `INSERT INTO course (category_id, name, detail, price, picture, dtm_crt, dtm_upd) VALUES (?, ?, ?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.CategoryID, req.Name, req.Detail, req.Price, req.Picture)
	if err != nil {
		err = errors.New("failed to create course")
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

func (db *mysqlCourseRepository) SelectCourseByID(ctx context.Context, id int64) (course domain.Course, err error) {
	query := `SELECT id, category_id, name, detail, price, picture, dtm_crt, dtm_upd FROM course WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&course.ID, &course.CategoryID, &course.Name, &course.Detail, &course.Price, &course.Picture, &course.DtmCrt, &course.DtmUpd)
	if err != nil {
		err = errors.New("course not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCourseRepository) SelectAllCourse(ctx context.Context, page, limit int64, search string, sort string, categoryId ...*string) (course []domain.Course, err error) {
	query := `SELECT id, category_id, name, detail, price, picture, dtm_crt, dtm_upd FROM course`

	if categoryId[0] != nil {
		var i []string
		for _, v := range categoryId {
			i = append(i, *v)
		}
		query += ` WHERE category_id IN (` + strings.Join(i, ",") + `)`
	}

	if search != "" {
		query += " WHERE MATCH(name) AGAINST('" + search + "')"
	}

	if sort == "lowest" {
		query += " ORDER BY price > 0 ASC"

	} else if sort == "highest" {
		query += " ORDER BY price > 0 DESC"

	} else if sort == "free" {
		query += " ORDER BY price = 0 ASC"

	} else {
		query += " ORDER BY id"

	}

	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Course
		err = rows.Scan(&i.ID, &i.CategoryID, &i.Name, &i.Detail, &i.Price, &i.Picture, &i.DtmCrt, &i.DtmUpd)
		if err != nil {
			log.Error(err)
			return
		}

		course = append(course, i)
	}

	return
}

func (db *mysqlCourseRepository) EditCourse(ctx context.Context, req domain.CourseRequest, id int64) (err error) {
	query := `UPDATE course SET category_id = ?, name = ?, detail = ?, price = ?, picture = ? WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.CategoryID, req.Name, req.Detail, req.Price, req.Picture, id)
	if err != nil {
		log.Error(err)
		return
	}

	updated, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}

	if updated == 0 {
		err = errors.New("no updated data")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCourseRepository) RemoveCourse(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM course WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}

	if affect == 0 {
		err = errors.New("data not found")
		return
	}

	return
}

func (db *mysqlCourseRepository) SelectTotalCourse(ctx context.Context) (count int64, err error) {
	query := `SELECT COUNT(id) FROM course`
	log.Debug(query)

	row := db.Conn.QueryRowContext(ctx, query)
	err = row.Scan(&count)
	if err != nil {
		err = errors.New("course not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCourseRepository) SelectTotalFreeCourse(ctx context.Context) (count int64, err error) {
	query := `SELECT COUNT(id) FROM course WHERE price = 0`
	log.Debug(query)

	row := db.Conn.QueryRowContext(ctx, query)
	err = row.Scan(&count)
	if err != nil {
		err = errors.New("course not found")
		log.Error(err)
		return
	}

	return
}
