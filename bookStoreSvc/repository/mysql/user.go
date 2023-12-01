package mysql

import (
	"database/sql"
	"github.com/thtrangphu/bookStoreSvc/entity"
	"time"
)

type User struct {
	db *sql.DB
}

func CreateNewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(e *entity.User) (entity.ID, error) {
	stmt, err := u.db.Prepare(`insert into user (id, email, password, full_name, create_at) values (?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Email,
		e.Password,
		e.FullName,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return e.ID, nil
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

func (u *User) Get(id entity.ID) (*entity.User, error) {
	return getUser(id, u.db)
}

func getUser(id entity.ID, db *sql.DB) (*entity.User, error) {
	stmt, err := db.Prepare(`select * from user where id = ?`)
	if err != nil {
		return nil, err
	}
	var u entity.User
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email, &u.FullName, &u.CreatedAt)
	}

	return &u, nil
}

func (u *User) Update(e *entity.User) error {
	e.UpdatedAt = time.Now()
	_, err := u.db.Exec("update user set email = ?, password = ?, full_name = ?, updated_at = ? where id = ?", e.Email, e.Password, e.FullName, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) List() ([]*entity.User, error) {
	stmt, err := u.db.Prepare(`select id from user`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []entity.ID
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i entity.ID
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, entity.ErrNotFound
	}
	var users []*entity.User
	for _, id := range ids {
		u, err := getUser(id, u.db)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Delete an user
func (u *User) Delete(id entity.ID) error {
	_, err := u.db.Exec("delete from user where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
