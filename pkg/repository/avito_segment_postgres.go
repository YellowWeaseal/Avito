package repository

import (
	"awesomeProject1"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AvitoSegmentPostgres struct {
	db *sqlx.DB
}

func NewAvitoSegmentPostgres(db *sqlx.DB) *AvitoSegmentPostgres {
	return &AvitoSegmentPostgres{db: db}
}

func (r *AvitoSegmentPostgres) Create(segment awesomeProject1.SegmentsUpdate) (int, error) {
	var id int

	createSegmentQuery := fmt.Sprintf("INSERT INTO %s (name, description) VALUES ($1, $2) RETURNING id",
		segmentsTable)
	_, err := r.db.Exec(createSegmentQuery, segment.Name, segment.Description)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *AvitoSegmentPostgres) GetAll() ([]awesomeProject1.Segment, error) {
	var segments []awesomeProject1.Segment

	query := fmt.Sprintf("SELECT id, name, description FROM %s",
		segmentsTable)
	err := r.db.Select(&segments, query)

	return segments, err
}

func (r *AvitoSegmentPostgres) GetById(segmentId int) (awesomeProject1.Segment, error) {
	var segment awesomeProject1.Segment

	query := fmt.Sprintf(`SELECT id, name, description FROM %s  WHERE id = $1 `,
		segmentsTable)
	err := r.db.Get(&segment, query, segmentId)

	return segment, err
}

func (r *AvitoSegmentPostgres) Delete(segmentId int) error {
	/*	tx, err := r.db.Begin()
		if err != nil {
			return 0, err
		}

		var userId int
		deleteSegmentQuery := fmt.Sprintf("DELETE FROM %s  WHERE id=$1",
			segmentsTable)

		row := tx.QueryRow(deleteSegmentQuery, segmentId)
		err = row.Scan(&userId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		deleteSegmentFromUserQuery := fmt.Sprintf("INSERT INTO %s (update_user_segment, updated_at) values ($1, $2)", updatedUserSegmentsTable)
		_, err = tx.Exec(createUserUpdatedIngoQuery, segment.SegmentId, time.Now())
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		return userId, tx.Commit()
		query := fmt.Sprintf("DELETE FROM %s  WHERE id=$1",
			todoListsTable, usersListsTable)
		_, err := r.db.Exec(query, userId, listId)*/

	return nil
}

func (r *AvitoSegmentPostgres) Update(segmentId int, input awesomeProject1.SegmentsUpdate) (int, error) {
	/*setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)*/
	return 0, nil
}
