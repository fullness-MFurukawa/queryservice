package handler_test

import (
	"log"
	"queryservice/apperror"
	"queryservice/infrastructure/gorm/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	db, err := handler.NewGORMConnector().Open()
	if err != nil {
		e := err.(*apperror.AppError)
		assert.Equal(t, e.Error(), "status:500,message:現在サービスの提供ができません。")
		return
	}
	log.Println("db = ", db)
	assert.NotNil(t, db)
}
