package repository_test

import (
	"context"
	"log"
	"queryservice/domain/models/products"
	"queryservice/infra/gorm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("ProductRepositoryGORM構造体", Ordered, Label("ProductRepositoryインターフェイスメソッドのテスト"), func() {
	var repository products.ProductRepository
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
	Context("すべての商品を取得する", Label("List"), func() {
		It("すべての商品を取得する", func() {
			products, _ := repository.List(ctx)
			for _, product := range products {
				log.Println(product)
			}
			Expect(len(products) > 0).To(BeTrue())
		})
	})
	// FindByProductId()メソッドのテスト
	Context("指定キーワードを含む商品の問合せ", Label("FindByProductId"), func() {
		It("存在する商品の商品IDを指定する", func() {
			product, err := repository.FindByProductId(ctx, "8f81a72a-58ef-422b-b472-d982e8665292")
			if err != nil {
				log.Println(err)
				Expect(err).Error()
				return
			}
			Expect(product.Id()).To(Equal("8f81a72a-58ef-422b-b472-d982e8665292"))
			Expect(product.Name).To(Equal("水性ボールペン(赤)"))
			Expect(product.Price).To(Equal(uint32(120)))
			Expect(product.Category().Id()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			Expect(product.Category().Name()).To(Equal("文房具"))
		})
		It("存在しない商品の商品IDを指定する", func() {
			_, err := repository.FindByProductId(ctx, "8f81a72a-58ef-422b-b472-d982e866529a")
			Expect(err.Error()).To(Equal("商品ID:8f81a72a-58ef-422b-b472-d982e866529aは存在しません。"))
		})
	})
	// FindByProductNameLike()メソッドのテスト
	Context("指定キーワードを含む商品の問合せ", Label("FindByProductNameLike"), func() {
		It("キーワード:ペンを指定した結果を評価する", func() {
			products, _ := repository.FindByProductNameLike(ctx, "ペン")
			for _, product := range products {
				log.Println(product)
			}
			Expect(len(products) > 0).To(BeTrue())
		})
	})

})
