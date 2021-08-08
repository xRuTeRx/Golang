package adress_book

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

type Contact struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	GroupId int    `db:"group_id"`
}

func connString(dbType, user, pwd, dbHost string, port int, dbName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", dbType, user, pwd, dbHost, port, dbName)
}

func (db *DB) AddContact(name string, phone string, groupId int) error {
	q := "INSERT INTO contacts (name, phone, group_id) VALUES ($1, $2, $3);"
	if _, err := db.Conn.Exec(q, name, phone, groupId); err != nil {
		return errors.Wrap(err, "failed to add contact")
	}
	return nil
}

func (db *DB) AssignToGroup(Id int, groupId int) error {
	q := "update contacts set  group_id=$1 where id=$2;"
	if _, err := db.Conn.Exec(q, groupId, Id); err != nil {
		return errors.Wrap(err, "failed to update contacts")
	}
	return nil
}

func rowsToContact(rows *sql.Rows) (contacts []Contact, err error) {
	contacts = make([]Contact, 0)
	for rows.Next() {
		u := &Contact{}
		if err = rows.Scan(&u.ID, &u.Name, &u.Phone, &u.GroupId); err != nil {
			return nil, errors.Wrap(err, "failed to List Contacts (scan)")
		}
		contacts = append(contacts, *u)
	}
	return contacts, nil
}

func (db *DB) ListAllByGroup(groupId int) ([]Contact, error) {
	q := "select * from contacts where group_id=$1"
	rows, err := db.Conn.Query(q, groupId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to List contacts")
	}
	return rowsToContact(rows)
}

func ConnToDB() (*DB, error) {
	connStr := connString("postgres", "root", "root", "localhost", 5432, "adress_book")
	c, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to DB")
	}
	db := &DB{
		Conn: c,
	}
	return db, nil

}
