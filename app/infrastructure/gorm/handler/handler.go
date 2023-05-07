package handler

import (
	"queryservice/apperror"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// *****
// コネクションプールの生成
// *****
type GORMConnector struct{}

/*
データベース接続結果を返す
*/
func (conn *GORMConnector) Open() (*gorm.DB, error) {
	// データベース接続URL
	url := "root:password@tcp(query_db:3307)/sample_db?charset=utf8mb4&parseTime=True&loc=Local"
	// データベース接続
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}
	// 接続プールを取得して設定する
	pool, err := db.DB()
	if err != nil {
		return nil, err
	}
	pool.SetMaxIdleConns(10)           // 最大接続数
	pool.SetMaxOpenConns(100)          // オープン接続最大数
	pool.SetConnMaxLifetime(time.Hour) //接続を再利用できる最大時間
	// 生成されたSQLをログ出力する
	db.Logger = db.Logger.LogMode(logger.Info)
	return db, nil
}

/*
コンストラクタ
*/
func NewGORMConnector() *GORMConnector {
	return &GORMConnector{}
}
