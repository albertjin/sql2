package sql2

import (
    "database/sql"
)

// DB config struct for json friendly.
type Config struct {
    Driver string `json:"driver"`
    Connection string `json:"connection"`
    Init string `json:"init"`
}

// Open DB connection with config.
func (config *Config) Open() (db *sql.DB, err error) {
    if db, err = sql.Open(config.Driver, config.Connection); err == nil {
        if len(config.Init) > 0 {
            _, err = db.Exec(config.Init)
            if err != nil {
                db.Close()
                db = nil
            }
        }
    }
    return
}

// Implementation for DBI.DbUse(). A new connection is always opened.
func (config *Config) DbUse() (db *sql.DB, err error) {
    return config.Open()
}

// Implementation for DBI.DbRelease().
func (config *Config) DbRelease(db *sql.DB) {
    if (db != nil) {
        db.Close()
    }
}
