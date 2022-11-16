package mysqlpkg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mikejk8s/gmud/logger"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "1234"
	//hostname = "docker.for.mac.localhost:3306"
	hostname = "127.0.0.1:3306"
	dbname   = "characters"
)

func Dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE SCHEMA IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", Dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	//defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

func createCharacterTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS characters (
		id BIGINT UNIQUE NOT NULL PRIMARY KEY,
		name VARCHAR(30) UNIQUE NOT NULL,
		class VARCHAR(15) NOT NULL,
		race VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
		level INT(3) NOT NULL DEFAULT '1',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		alive BOOLEAN NOT NULL DEFAULT '1',
        characterowner VARCHAR(20) NOT NULL DEFAULT 'player'
	) ENGINE=INNODB;`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating Character table", err)
		panic(err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		panic(err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func Connect() {
	// Get a new logger instance for database
	DBConnectionLogger := logger.GetNewLogger()
	logDirectory := fmt.Sprintf("./logs/dbconn")
	err := DBConnectionLogger.AssignOutput("dbLog", logDirectory)
	if err != nil {
		log.Printf("Error %s when assigning output to logger", err)
	}
	db, err := DbConnection()
	if err != nil {
		DBConnectionLogger.LogUtil.Errorln("Error connecting to DB: ", err)
		panic(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			DBConnectionLogger.LogUtil.Errorf("Error %s while connecting to DB: ", err)
			return
		}
	}(db)
	DBConnectionLogger.LogUtil.Infoln("Connected to DB successfully")
	err = createCharacterTable(db)
	if err != nil {
		DBConnectionLogger.LogUtil.Errorf("Error %s while creating Character table", err)
		panic(err)
		return
	}
}

// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// defer cancelfunc()
