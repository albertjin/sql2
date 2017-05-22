package sql2

import (
    "crypto/sha256"
    "database/sql"
    "regexp"
    "sync"

    "github.com/albertjin/ec"
)

var _IsSafeDbNameOnce sync.Once
var _IsSafeDbNameRe *regexp.Regexp

func IsSafeDbName(s string) bool {
    _IsSafeDbNameOnce.Do(func() {
        _IsSafeDbNameRe = regexp.MustCompile(`[a-zA-Z0-9_]+`)
    })
    return _IsSafeDbNameRe.MatchString(s)
}

// Convert text to id hashed with SHA256.
func Text2Id(db *sql.DB, dbName, text string, lock *sync.Mutex, toFulltext func (text string) string) (id int, err error) {
    if lock != nil {
        lock.Lock(); defer lock.Unlock()
    }

    hash0 := sha256.Sum256([]byte(text))
    hash := hash0[:]
    if db.QueryRow("select `id` from `" + dbName + "`.`String` where `sha256`=?", hash).Scan(&id) == nil {
        return
    }

    var result sql.Result
    if toFulltext != nil {
        result, err = db.Exec("insert into `" + dbName + "`.`String`(`sha256`, `words`, `original`) values(?,?,?)", hash, toFulltext(text), text)
    } else {
        result, err = db.Exec("insert into `" + dbName + "`.`String`(`sha256`, `original`) values(?,?)", hash, text)
    }
    if err != nil {
        err = ec.Wrapf(err, "fail at db.Exec() with db name: %v", dbName)
        return
    }

    x, err := result.LastInsertId()
    if err != nil {
        err = ec.Wrap(err, "fail at LastInsertId()")
        return
    }

    id = int(x)
    return
}
