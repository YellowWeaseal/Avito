package repository

import (
	"awesomeProject1"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AvitoUserPostgres struct {
	db *sqlx.DB
}

func NewAvitoUserPostgres(db *sqlx.DB) *AvitoUserPostgres {
	return &AvitoUserPostgres{db: db}
}

func (r *AvitoUserPostgres) Create(user awesomeProject1.UserUpdate) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var userId int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (user_segments) values ($1) RETURNING id", usersTable)

	row := tx.QueryRow(createUserQuery, user.UserSegments)
	err = row.Scan(&userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserUpdatedIngoQuery := fmt.Sprintf("INSERT INTO %s (update_user_segment, updated_at) values ($1, $2)", updatedUserSegmentsTable)
	_, err = tx.Exec(createUserUpdatedIngoQuery, user.UserSegments, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return userId, tx.Commit()
}

func (r *AvitoUserPostgres) GetAll() ([]awesomeProject1.User, error) {
	var users []awesomeProject1.User
	query := fmt.Sprintf("SELECT id, user_segments FROM %s", usersTable)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *AvitoUserPostgres) GetById(userId int) (awesomeProject1.User, error) {
	var user awesomeProject1.User
	query := fmt.Sprintf(`SELECT id, user_segments FROM %s WHERE id = $1 `,
		usersTable)
	if err := r.db.Get(&user, query, user.UserId, user.UserSegments); err != nil {
		return user, err
	}

	return user, nil
}

func (r *AvitoUserPostgres) Delete(userId int) error {
	query := fmt.Sprintf(`DELETE FROM %s  WHERE id=$1`,
		usersTable)
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *AvitoUserPostgres) Update(userId int, input awesomeProject1.UserUpdate) (int, error) {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)*/
	return 0, nil
}
