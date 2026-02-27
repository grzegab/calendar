package db

import "database/sql"

type PostgresTxManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *PostgresTxManager {
	return &PostgresTxManager{db: db}
}

func (m *PostgresTxManager) WithinTx(fn func(tx *sql.Tx) error) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
