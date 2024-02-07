package handler_test

import (
	"queryservice/infra/gorm/handler"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/gorm/handlerパッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("接続が成功した場合、*gorm.DBが返る", Label("DB接続"), func() {
		conn, err := handler.ConnectDB()
		if err != nil {
			Fail(err.Error())
		}
		Expect(conn).ToNot(BeNil())
	})
})
