package handler_test

import (
	"log"
	"queryservice/infrastructure/gorm/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	db, err := handler.NewGORMConnector().Open()
	if err != nil {
		assert.Error(t, err)
		return
	}
	log.Println("db = ", db)
	assert.NotNil(t, db)
}
