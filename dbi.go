package sql2
import "database/sql"

type DBI interface {
    DbUse() (db *sql.DB, err error)
    DbRelease(db *sql.DB)
}
