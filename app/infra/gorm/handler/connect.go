package handler

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// query_dbに接続
func ConnectDB() (*gorm.DB, error) {
	// データベースへ接続
	dcs := "root:password@tcp(query_db:3306)/sample_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dcs), &gorm.Config{})
	if err != nil {
		return nil, DBErrHandler(err)
	}
	// *sql.DBの取得
	if db, err := conn.DB(); err != nil {
		return nil, DBErrHandler(err)
	} else {
		// データベース接続の確認
		if err := db.Ping(); err != nil {
			return nil, DBErrHandler(err)
		}
		// コネクションプールの設定
		db.SetConnMaxIdleTime(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)
		// 生成されたSQLをログ出力する
		conn.Logger = conn.Logger.LogMode(logger.Info)
		return conn, nil
	}
}
