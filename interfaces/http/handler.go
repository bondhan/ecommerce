package handler

import (
	presentercashier "github.com/bondhan/ecommerce/modules/cashier/presenter"
	querycashier "github.com/bondhan/ecommerce/modules/cashier/query"
	usecasecashier "github.com/bondhan/ecommerce/modules/cashier/usecase"

	presentercategory "github.com/bondhan/ecommerce/modules/category/presenter"
	querycategory "github.com/bondhan/ecommerce/modules/category/query"
	usecasecategory "github.com/bondhan/ecommerce/modules/category/usecase"

	presenterpayment "github.com/bondhan/ecommerce/modules/payment/presenter"
	querypayment "github.com/bondhan/ecommerce/modules/payment/query"
	usecasepayment "github.com/bondhan/ecommerce/modules/payment/usecase"

	presenterproduct "github.com/bondhan/ecommerce/modules/product/presenter"
	queryproduct "github.com/bondhan/ecommerce/modules/product/query"
	usecaseproduct "github.com/bondhan/ecommerce/modules/product/usecase"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Handler struct {
	cashier  presentercashier.ICashierP
	category presentercategory.ICategoryP
	payment  presenterpayment.IPaymentP
	product  presenterproduct.IProductP
}

func NewHandler(logger *logrus.Logger, db *gorm.DB) *Handler {
	cashierQ := querycashier.NewCashierQ(logger, db)
	cashierUC := usecasecashier.NewCashierUC(logger, cashierQ)
	cashierP := presentercashier.NewCashierP(cashierUC)

	categoryQ := querycategory.NewCategoryQ(logger, db)
	categoryUC := usecasecategory.NewCategoryUC(logger, categoryQ)
	categoryP := presentercategory.NewCategoryP(categoryUC)

	paymentQ := querypayment.NewPaymentQ(logger, db)
	paymentUC := usecasepayment.NewPaymentUC(logger, paymentQ)
	paymentP := presenterpayment.NewPaymentP(paymentUC)

	productQ := queryproduct.NewProductQ(logger, db)
	productUC := usecaseproduct.NewProductUC(logger, productQ)
	productP := presenterproduct.NewProductP(productUC)

	return &Handler{
		cashier:  cashierP,
		category: categoryP,
		payment:  paymentP,
		product:  productP,
	}
}
