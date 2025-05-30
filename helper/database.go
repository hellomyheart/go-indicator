package helper

import (
	"database/sql"
	"fmt"
	"log"
)

// CloseDatabaseWithError 安全关闭数据库连接
func CloseDatabaseWithError(db *sql.DB, err error) error {
	closeErr := db.Close()
	if closeErr == nil {
		return err
	}

	closeErr = fmt.Errorf("unable to close database: %w", closeErr)

	if err != nil {
		log.Println(closeErr)
		return err
	}

	return closeErr
}

// CloseDatabaseRows 安全关闭数据库查询结果集
func CloseDatabaseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Println(err)
	}
}
