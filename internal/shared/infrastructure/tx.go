package infrastructure

import "database/sql"

type TxManager interface {
	WithinTx(fn func(tx *sql.Tx) error) error
}
