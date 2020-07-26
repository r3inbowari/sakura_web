package da

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sakura_web/utils"
)

var db *gorm.DB

func DBC() *gorm.DB {
	if db == nil {
		lc := utils.GetConfig()
		var err error
		db, err = gorm.Open("mysql", lc.GetDatabaseUsername()+":"+lc.GetDatabasePassword()+"@tcp("+lc.GetDatabaseHost()+")/"+lc.GetDatabaseTable()+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("error dbs")
		}

		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(true)
	}
	return db
}
