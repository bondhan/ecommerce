package model

import (
	"encoding/json"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	modelcashier "github.com/bondhan/ecommerce/modules/cashier/model"
	modelpayment "github.com/bondhan/ecommerce/modules/payment/model"
	"github.com/bondhan/ecommerce/modules/product/model"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strconv"
)

type CreateOrderResp struct {
	Order
	PassCode  string `json:"passcode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateOrderReq struct {
	Name     string `json:"name"`
	PassCode string `json:"passcode"`
}

type Order struct {
	OrderID uint   `json:"orderId"`
	Name    string `json:"name"`
}

type CreateOrderUpdate struct {
	ID uint
	CreateOrderReq
}

func (c CreateOrderReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.PassCode, validation.Required, is.Digit,
			validation.Length(6, 6)),
	)
}

type OrderPaginatedReq struct {
	Skip  string `json:"skip"`
	Limit string `json:"limit"`
}

func NewOrderPaginatedReq(r *http.Request) (OrderPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)

	orders := OrderPaginatedReq{
		Skip:  skip,
		Limit: limit,
	}

	err := validation.ValidateStruct(&orders,
		validation.Field(&orders.Limit, is.Digit),
		validation.Field(&orders.Skip, is.Digit),
	)

	var l, s = 0, 0

	if len(orders.Limit) > 0 {
		l, err = strconv.Atoi(orders.Limit)
		if err != nil {
			return OrderPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(orders.Skip) > 0 {
		s, err = strconv.Atoi(orders.Skip)
		if err != nil {
			return OrderPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	vp := OrderPaginated{
		Skip:  s,
		Limit: l,
	}

	return vp, nil
}

type OrderPaginated struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

type SubTotalReq struct {
	ProductID uint  `json:"productId"`
	Qty       int64 `json:"qty"`
}

func (c SubTotalReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ProductID, validation.Required),
		validation.Field(&c.Qty, validation.Required),
	)
}

func NewSubtotalOrder(r *http.Request) ([]SubTotalReq, error) {
	var req []SubTotalReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	if len(req) < 1 {
		return req, ecommerceerror.ErrEmptyBody
	}

	for _, v := range req {
		err := v.Validate()
		if err != nil {
			return req, err
		}
	}

	return req, nil
}

type ProductSubTotal struct {
	ProductID        uint                  `json:"productId"`
	Name             string                `json:"name"`
	Price            int64                 `json:"price"`
	Discount         *model.DiscountDetail `json:"discount"`
	Qty              int64                 `json:"qty"`
	TotalNormalPrice float64               `json:"totalNormalPrice"`
	TotalFinalPrice  float64               `json:"totalFinalPrice"`
}

type SubTotal struct {
	Subtotal float64           `json:"subtotal"`
	Products []ProductSubTotal `json:"products"`
}

type OrderReq struct {
	CashierID uint          `json:"cashierId"`
	PaymentID uint          `json:"paymentId"`
	TotalPaid float64       `json:"totalPaid"`
	Products  []SubTotalReq `json:"products"`
}

func NewSOrder(r *http.Request) (OrderReq, error) {
	var req OrderReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	if len(req.Products) < 1 {
		return OrderReq{}, ecommerceerror.ErrEmptyProduct
	}

	for _, v := range req.Products {
		err := v.Validate()
		if err != nil {
			return req, err
		}
	}

	claims, ok := r.Context().Value(params.JWTClaims).(*basemodel.Claims)
	if !ok {
		return req, ecommerceerror.ErrUnauthorized
	}
	ID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return req, err
	}
	req.CashierID = uint(ID)
	return req, nil

}

type OrderTotal struct {
	OrderID        uint    `json:"orderId"`
	CashiersID     string  `json:"cashiersId"`
	PaymentTypesID uint    `json:"paymentTypesId"`
	TotalPrice     float64 `json:"totalPrice"`
	TotalPaid      float64 `json:"totalPaid"`
	TotalReturn    float64 `json:"totalReturn"`
	ReceiptID      string  `json:"receiptId"`
	UpdatedAt      string  `json:"updatedAt"`
	CreatedAt      string  `json:"createdAt"`
}

type OrderTotalResp struct {
	Order    OrderTotal        `json:"order"`
	Products []ProductSubTotal `json:"products"`
}

type OrderDetailDB struct {
	OrderID        uint      `gorm:"column:id"`
	CashiersID     uint      `gorm:"column:cashier_id"`
	PaymentTypesID uint      `gorm:"column:payment_type_id"`
	TotalPrice     int64     `gorm:"column:total_price"`
	TotalPaid      int64     `gorm:"column:total_paid"`
	TotalReturn    int64     `gorm:"column:total_return"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	CashierID      uint      `gorm:"column:cashier_id"`
	CashierName    string    `gorm:"column:cashier_name"`
	PaymentTypeID  uint      `gorm:"column:payment_type_id"`
	PaymentName    string    `gorm:"column:payment_name"`
	PaymentType    string    `gorm:"column:payment_type"`
	PaymentLogo    *string   `gorm:"column:payment_logo"`
}

type OrderDetail struct {
	OrderID        uint                 `json:"orderId"`
	CashiersID     uint                 `json:"cashiersId"`
	PaymentTypesID uint                 `json:"paymentTypesId"`
	TotalPrice     int64                `json:"totalPrice"`
	TotalPaid      int64                `json:"totalPaid"`
	TotalReturn    int64                `json:"totalReturn"`
	ReceiptID      string               `json:"receiptId"`
	UpdatedAt      string               `json:"updatedAt"`
	CreatedAt      string               `json:"createdAt"`
	Cashier        modelcashier.Cashier `json:"cashier"`
	Payment        modelpayment.Payment `json:"payment"`
}

type ListOrderResponse struct {
	OrderDetails []OrderDetail  `json:"orders"`
	Meta         basemodel.Meta `json:"meta"`
}

func GetOrderID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrOrderNotFound
	}

	return ID, nil
}

type OrderProductDetail struct {
	OrderID        uint                 `json:"orderId"`
	CashiersID     uint                 `json:"cashiersId"`
	PaymentTypesID uint                 `json:"paymentTypesId"`
	TotalPrice     int64                `json:"totalPrice"`
	TotalPaid      int64                `json:"totalPaid"`
	TotalReturn    int64                `json:"totalReturn"`
	ReceiptID      string               `json:"receiptId"`
	CreatedAt      string               `json:"createdAt"`
	Cashier        modelcashier.Cashier `json:"cashier"`
	Payment        modelpayment.Payment `json:"paymentType"`
}

type OrderProductSubTotal struct {
	ProductID        uint                  `json:"productId"`
	Name             string                `json:"name"`
	Price            int64                 `json:"price"`
	DiscountID       *uint                 `json:"discountsId"`
	Discount         *model.DiscountDetail `json:"discount"`
	Qty              int64                 `json:"qty"`
	TotalNormalPrice float64               `json:"totalNormalPrice"`
	TotalFinalPrice  float64               `json:"totalFinalPrice"`
}

type DetailOrderProductResponse struct {
	OrderDetail OrderProductDetail     `json:"order"`
	Products    []OrderProductSubTotal `json:"products"`
}

type DownloadStatus struct {
	IsDownloaded bool `json:"isDownload"`
}

type RevenueByPaymentType struct {
	PaymentTypeID uint    `gorm:"column:payment_type_id"  json:"paymentTypeId"`
	Name          string  `gorm:"column:name"  json:"name"`
	Logo          *string `gorm:"column:logo"  json:"logo"`
	TotalAmount   float64 `gorm:"column:total_amount"  json:"totalAmount"`
}

type Revenue struct {
	TotalRevenue          float64                `json:"totalRevenue"`
	RevenueByPaymentTypes []RevenueByPaymentType `json:"paymentTypes"`
}

type Solds struct {
	ProductID   uint    `gorm:"column:product_id" json:"productId"`
	Name        string  `gorm:"column:name" json:"name"`
	TotalQty    int64   `gorm:"column:total_qty" json:"totalQty"`
	TotalAmount float64 `gorm:"column:total_amount" json:"totalAmount"`
}
