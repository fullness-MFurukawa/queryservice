package repository_test

import (
	"context"
	"log"
	"queryservice/domain/models/categories"
	"queryservice/infra/gorm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("CategoryRepositoryGORM構造体", Ordered, Label("CategoryRepositoryインターフェイスメソッドのテスト"), func() {
	var repository categories.CategoryRepository
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		// Contextの生成
		ctx = context.Background()
		// サービスのインスタンス生成
		container = fx.New(
			gorm.RepDepend,
			fx.Populate(&repository),
		)
		// fxを起動し、起動時にエラーがないことを確認する
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// 後処理
	AfterAll(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})
	// List()メソッドのテスト
	Context("商品カテゴリリストの取得", Label("List"), func() {
		It("商品カテゴリのスライスを返す", func() {
			categories, _ := repository.List(ctx)
			Expect(len(categories) > 0).To(Equal(true))
			for _, category := range categories {
				log.Println(category)
			}
		})
	})
	// FindByCategoryId()メソッドのテスト
	Context("カテゴリID(UUID)で該当するカテゴリを取得する", Label("FindByCategoryId"), func() {
		It("存在するカテゴリIDで問合せする", func() {
			category, _ := repository.FindByCategoryId(ctx, "b1524011-b6af-417e-8bf2-f449dd58b5c0")
			Expect(category.Id()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			Expect(category.Name()).To(Equal("文房具"))
		})
	})
})
