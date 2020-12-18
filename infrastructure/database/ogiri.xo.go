// Package database contains the types for schema 'ogiri'.
package database

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// Ogiri represents a row from 'ogiri.ogiri'.
type Ogiri struct {
	OgiriID          string `json:"ogiri_id"`          // ogiri_id
	OdaiID           int    `json:"odai_id"`           // odai_id
	AnswerDuration   int    `json:"answer_duration"`   // answer_duration
	VoteDuration     int    `json:"vote_duration"`     // vote_duration
	QuestionDuration int    `json:"question_duration"` // question_duration

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ogiri exists in the database.
func (o *Ogiri) Exists() bool {
	return o._exists
}

// Deleted provides information if the Ogiri has been deleted from the database.
func (o *Ogiri) Deleted() bool {
	return o._deleted
}

// Insert inserts the Ogiri to the database.
func (o *Ogiri) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO ogiri.ogiri (` +
		`ogiri_id, odai_id, answer_duration, vote_duration, question_duration` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, o.OgiriID, o.OdaiID, o.AnswerDuration, o.VoteDuration, o.QuestionDuration)
	_, err = db.Exec(sqlstr, o.OgiriID, o.OdaiID, o.AnswerDuration, o.VoteDuration, o.QuestionDuration)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Update updates the Ogiri in the database.
func (o *Ogiri) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE ogiri.ogiri SET ` +
		`odai_id = ?, answer_duration = ?, vote_duration = ?, question_duration = ?` +
		` WHERE ogiri_id = ?`

	// run query
	XOLog(sqlstr, o.OdaiID, o.AnswerDuration, o.VoteDuration, o.QuestionDuration, o.OgiriID)
	_, err = db.Exec(sqlstr, o.OdaiID, o.AnswerDuration, o.VoteDuration, o.QuestionDuration, o.OgiriID)
	return err
}

// Save saves the Ogiri to the database.
func (o *Ogiri) Save(db XODB) error {
	if o.Exists() {
		return o.Update(db)
	}

	return o.Insert(db)
}

// Delete deletes the Ogiri from the database.
func (o *Ogiri) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return nil
	}

	// if deleted, bail
	if o._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM ogiri.ogiri WHERE ogiri_id = ?`

	// run query
	XOLog(sqlstr, o.OgiriID)
	_, err = db.Exec(sqlstr, o.OgiriID)
	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

// Odai returns the Odai associated with the Ogiri's OgiriID (ogiri_id).
//
// Generated from foreign key 'ogiri_ibfk_1'.
func (o *Ogiri) Odai(db XODB) (*Odai, error) {
	return OdaiByNextOgiriID(db, o.OgiriID)
}

// OgiriByOgiriID retrieves a row from 'ogiri.ogiri' as a Ogiri.
//
// Generated from index 'ogiri_ogiri_id_pkey'.
func OgiriByOgiriID(db XODB, ogiriID string) (*Ogiri, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ogiri_id, odai_id, answer_duration, vote_duration, question_duration ` +
		`FROM ogiri.ogiri ` +
		`WHERE ogiri_id = ?`

	// run query
	XOLog(sqlstr, ogiriID)
	o := Ogiri{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, ogiriID).Scan(&o.OgiriID, &o.OdaiID, &o.AnswerDuration, &o.VoteDuration, &o.QuestionDuration)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
