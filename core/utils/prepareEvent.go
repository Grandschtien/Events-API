package utils

import "database/sql"

func PrepareEvent(
	conn *sql.DB,
	operation string,
) *sql.Stmt {
	stmt, err := conn.Prepare(
		operation,
	)

	if err != nil {
		panic("Operation did not prepare")
	}

	return stmt
}
