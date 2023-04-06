package config

import (
	model "challenges_9/module/model/book"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	// minimum config untuk postgres
	Port              uint
	Host              string
	Username          string
	Password          string
	DBName            string
	MaxIdleConnection int
	MaxOpenConnection int
	MaxIdleTime       int
}

func NewDBPostges() *sql.DB {
	pgConf := PostgresConfig{
		Port:              5432,
		Host:              "127.0.0.1",
		Username:          "postgres",
		Password:          "billy",
		DBName:            "DB_Books",
		MaxOpenConnection: 7,
		MaxIdleConnection: 5,
		MaxIdleTime:       int(30 * time.Minute),
	}

	connString := fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
	`,
		pgConf.Host,
		pgConf.Port,
		pgConf.Username,
		pgConf.Password,
		pgConf.DBName,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	// test connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// set extended config
	db.SetMaxIdleConns(pgConf.MaxIdleConnection)
	db.SetMaxOpenConns(pgConf.MaxOpenConnection)
	db.SetConnMaxIdleTime(time.Duration(pgConf.MaxIdleTime))

	return db

}

func NewDBPostgesGormConn() *gorm.DB {
	pgConf := PostgresConfig{
		Port:              5432,
		Host:              "127.0.0.1",
		Username:          "postgres",
		Password:          "billy",
		DBName:            "db_books_tes",
		MaxOpenConnection: 7,
		MaxIdleConnection: 5,
		MaxIdleTime:       int(30 * time.Minute),
	}

	connString := fmt.Sprintf(`host=%v port=%v user=%v password=%v dbname=%v`,
		pgConf.Host,
		pgConf.Port,
		pgConf.Username,
		pgConf.Password,
		pgConf.DBName,
	)
	fmt.Println(connString)
	//	dsn := "host=127.0.0.1 user=postgres password=billy dbname=db_books_tes port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.SetMaxIdleConns(pgConf.MaxIdleConnection)
	dbSQL.SetMaxOpenConns(pgConf.MaxOpenConnection)
	//tenggat waktu apabila koneksi tidak digunakan akan di matikan
	dbSQL.SetConnMaxIdleTime(time.Duration(pgConf.MaxIdleTime))
	dbSQL.SetConnMaxLifetime(60 * time.Minute)

	if err = db.AutoMigrate(&model.Book{}); err != nil {
		panic(err)
	}

	return db
}
