package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			log.Panic(errRollBack)
		}
		log.Panic(err)
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			log.Panic(errCommit)
		}
	}
}
