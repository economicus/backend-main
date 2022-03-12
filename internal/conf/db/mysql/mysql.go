package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const maxIter = 20

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

type DB struct {
	conf
	DB *gorm.DB
}

func New() *DB {
	ms := DB{}
	ms.conf = newConf()
	ms.openGorm()
	ms.setup()
	ms.testDBConnection()
	ms.Info()
	return &ms
}

func (ms *DB) GetSqlDB() *sql.DB {
	db, err := ms.DB.DB()
	if err != nil {
		log.Panicf("error while getting sql db: %v", err)
	}
	return db
}

func (ms *DB) openGorm() {
	var err error
	dsn := ms.conf.GetDSN()

	for i := 0; i < maxIter; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			ms.DB = db
			return
		} else {
			log.Println("MySQL connection has failed. Sleeping for a second...")
			time.Sleep(1 * time.Second)
		}
	}
	log.Panicf("error while connecting mysql with dsn '%s': %v", dsn, err)
}

func (ms *DB) setup() {
	mysqlDB := ms.GetSqlDB()
	mysqlDB.SetMaxOpenConns(maxOpenDBConn)
	mysqlDB.SetMaxIdleConns(maxIdleDBConn)
	mysqlDB.SetConnMaxLifetime(maxDBLifeTime)
}

func (ms *DB) testDBConnection() {
	for i := 0; i < maxIter; i++ {
		if err := ms.GetSqlDB().Ping(); err == nil {
			fmt.Println("[INFO] Pinged mysql successfully")
			return
		} else {
			fmt.Println("[INFO] MySQL has not been prepared yet. Sleeping for a second...")
			time.Sleep(1 * time.Second)
		}
	}
	log.Panicf("error while testing connection: tried %d times but failed", maxIter)
}

func (ms *DB) Migrate(objs []interface{}) {
	for i := range objs {
		if err := ms.DB.AutoMigrate(objs[i]); err != nil {
			log.Fatalf("error while auto migration: %v", err)
		}
	}
}
