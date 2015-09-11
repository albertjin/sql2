package sql2

import (
    "database/sql"
)

// Config struct for writing in json.
type Config struct {
    Driver string `json:"driver"`
    Connection string `json:"connection"`
    Init string `json:"init"`
}

// Open db connection with config.
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

// Implementation for DBI.DbUse().
func (config *Config) DbUse() (db *sql.DB, err error) {
    return config.Open()
}

// Implementation for DBI.DbRelease().
func (config *Config) DbRelease(db *sql.DB) {
    if (db != nil) {
        db.Close()
    }
}
