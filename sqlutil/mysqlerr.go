package sqlutil

import "github.com/go-sql-driver/mysql"

// MysqlErrCode mysql error code
type MysqlErrCode uint16

const (
	MysqlDupEntry        MysqlErrCode = 1062
	MysqlRowIsReferenced MysqlErrCode = 1451
	MysqlNoReferencedRow MysqlErrCode = 1452
)

// GetMysqlErrCode get mysql error code, if error is nil or not a mysql error, return 0
func GetMysqlErrCode(err error) MysqlErrCode {
	e, _ := err.(*mysql.MySQLError)
	if e == nil {
		return 0
	}
	return MysqlErrCode(e.Number)
}
