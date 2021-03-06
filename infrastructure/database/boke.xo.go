// Package database contains the types for schema 'ogiri'.
package database

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Boke represents a row from 'ogiri.boke'.
type Boke struct {
	BokeID      int       `json:"boke_id"`      // boke_id
	AnswererID  int       `json:"answerer_id"`  // answerer_id
	Boke        string    `json:"boke"`         // boke
	OgiriID     string    `json:"ogiri_id"`     // ogiri_id
	PublishedAt time.Time `json:"published_at"` // published_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Boke exists in the database.
func (b *Boke) Exists() bool {
	return b._exists
}

// Deleted provides information if the Boke has been deleted from the database.
func (b *Boke) Deleted() bool {
	return b._deleted
}

// Insert inserts the Boke to the database.
func (b *Boke) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO ogiri.boke (` +
		`boke_id, answerer_id, boke, ogiri_id, published_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, b.BokeID, b.AnswererID, b.Boke, b.OgiriID, b.PublishedAt)
	_, err = db.Exec(sqlstr, b.BokeID, b.AnswererID, b.Boke, b.OgiriID, b.PublishedAt)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Boke in the database.
func (b *Boke) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE ogiri.boke SET ` +
		`answerer_id = ?, boke = ?, ogiri_id = ?, published_at = ?` +
		` WHERE boke_id = ?`

	// run query
	XOLog(sqlstr, b.AnswererID, b.Boke, b.OgiriID, b.PublishedAt, b.BokeID)
	_, err = db.Exec(sqlstr, b.AnswererID, b.Boke, b.OgiriID, b.PublishedAt, b.BokeID)
	return err
}

// Save saves the Boke to the database.
func (b *Boke) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Delete deletes the Boke from the database.
func (b *Boke) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM ogiri.boke WHERE boke_id = ?`

	// run query
	XOLog(sqlstr, b.BokeID)
	_, err = db.Exec(sqlstr, b.BokeID)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// User returns the User associated with the Boke's AnswererID (answerer_id).
//
// Generated from foreign key 'boke_ibfk_1'.
func (b *Boke) User(db XODB) (*User, error) {
	return UserByID(db, b.AnswererID)
}

// Odai returns the Odai associated with the Boke's OgiriID (ogiri_id).
//
// Generated from foreign key 'boke_ibfk_2'.
func (b *Boke) Odai(db XODB) (*Odai, error) {
	return OdaiByNextOgiriID(db, b.OgiriID)
}

// BokeByBokeID retrieves a row from 'ogiri.boke' as a Boke.
//
// Generated from index 'boke_boke_id_pkey'.
func BokeByBokeID(db XODB, bokeID int) (*Boke, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`boke_id, answerer_id, boke, ogiri_id, published_at ` +
		`FROM ogiri.boke ` +
		`WHERE boke_id = ?`

	// run query
	XOLog(sqlstr, bokeID)
	b := Boke{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, bokeID).Scan(&b.BokeID, &b.AnswererID, &b.Boke, &b.OgiriID, &b.PublishedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// BokesByAnswererID retrieves a row from 'ogiri.boke' as a Boke.
//
// Generated from index 'fk_answerer_id'.
func BokesByAnswererID(db XODB, answererID int) ([]*Boke, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`boke_id, answerer_id, boke, ogiri_id, published_at ` +
		`FROM ogiri.boke ` +
		`WHERE answerer_id = ?`

	// run query
	XOLog(sqlstr, answererID)
	q, err := db.Query(sqlstr, answererID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Boke{}
	for q.Next() {
		b := Boke{
			_exists: true,
		}

		// scan
		err = q.Scan(&b.BokeID, &b.AnswererID, &b.Boke, &b.OgiriID, &b.PublishedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}

// BokesByOgiriID retrieves a row from 'ogiri.boke' as a Boke.
//
// Generated from index 'fk_ogiri_id'.
func BokesByOgiriID(db XODB, ogiriID string) ([]*Boke, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`boke_id, answerer_id, boke, ogiri_id, published_at ` +
		`FROM ogiri.boke ` +
		`WHERE ogiri_id = ?`

	// run query
	XOLog(sqlstr, ogiriID)
	q, err := db.Query(sqlstr, ogiriID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Boke{}
	for q.Next() {
		b := Boke{
			_exists: true,
		}

		// scan
		err = q.Scan(&b.BokeID, &b.AnswererID, &b.Boke, &b.OgiriID, &b.PublishedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}
