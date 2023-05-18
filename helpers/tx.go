package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		IsPanicErrors(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		IsPanicErrors(errCommit)
	}
}
