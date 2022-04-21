package usecase

import (
	"fmt"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	modelcashier "github.com/bondhan/ecommerce/modules/cashier/model"
	"github.com/bondhan/ecommerce/modules/order/model"
	"github.com/bondhan/ecommerce/modules/order/query"
	modelpayment "github.com/bondhan/ecommerce/modules/payment/model"
	queryproduct "github.com/bondhan/ecommerce/modules/product/query"
	usecaseproduct "github.com/bondhan/ecommerce/modules/product/usecase"
	"github.com/spf13/cast"
	"time"

	"github.com/sirupsen/logrus"
)

type orderUC struct {
	logger    *logrus.Logger
	orderQ    query.IOrderQ
	productQ  queryproduct.IProductQ
	productUC usecaseproduct.IProductUC
}

func NewOrderUC(logger *logrus.Logger, orderQ query.IOrderQ,
	productQ queryproduct.IProductQ, productUC usecaseproduct.IProductUC) IOrderUC {
	return &orderUC{
		logger:    logger,
		orderQ:    orderQ,
		productQ:  productQ,
		productUC: productUC,
	}
}
func (c orderUC) Create(req model.OrderReq) (model.OrderTotalResp, error) {

	data, err := c.SubTotal(req.Products)
	if err != nil {
		return model.OrderTotalResp{}, err
	}

	order := model.OrderTotal{
		CashiersID:     cast.ToString(req.CashierID),
		PaymentTypesID: req.PaymentID,
		TotalPrice:     data.Subtotal,
		TotalPaid:      req.TotalPaid,
		TotalReturn:    req.TotalPaid - data.Subtotal,
	}

	res, err := c.orderQ.Insert(order, data.Products)
	if err != nil {
		return model.OrderTotalResp{}, err
	}
	resp := model.OrderTotalResp{
		Order: model.OrderTotal{
			OrderID:        res.ID,
			CashiersID:     cast.ToString(res.CashierID),
			PaymentTypesID: res.PaymentTypeID,
			TotalPrice:     res.TotalPrice,
			TotalPaid:      res.TotalPaid,
			TotalReturn:    res.TotalReturn,
			ReceiptID:      fmt.Sprintf("ID%03d", res.ID),
			UpdatedAt:      res.UpdatedAt.UTC().Format(time.RFC3339),
			CreatedAt:      res.CreatedAt.UTC().Format(time.RFC3339),
		},
		Products: data.Products,
	}

	return resp, nil
}

func (c orderUC) SubTotal(req []model.SubTotalReq) (model.SubTotal, error) {
	var total float64
	pst := []model.ProductSubTotal{}
	for _, v := range req {
		prod, err := c.productUC.Detail(v.ProductID)
		if err != nil {
			return model.SubTotal{}, err
		}

		if prod.Stock <= 0 {
			return model.SubTotal{}, ecommerceerror.ErrOutOfStock
		}

		prd := model.ProductSubTotal{
			ProductID:        prod.ProductID,
			Name:             prod.Name,
			Price:            prod.Price,
			Discount:         prod.Discount,
			Qty:              v.Qty,
			TotalNormalPrice: float64(v.Qty * prod.Price),
			TotalFinalPrice:  float64(v.Qty * prod.Price),
		}

		if prod.Discount != nil {
			if prod.Discount.Type == params.BuyN {
				multiplier := v.Qty / prod.Discount.Qty
				remains := v.Qty % prod.Discount.Qty
				//prd.TotalFinalPrice = int64(math.Round(float64(multiplier*prod.Discount.Result + remains*prod.Price)))
				prd.TotalFinalPrice = float64(multiplier*prod.Discount.Result) + float64(remains*prod.Price)
			} else if prod.Discount.Type == params.Percentage {
				//prd.TotalFinalPrice = int64(math.Round(float64(v.Qty * (prod.Price - prod.Price*prod.Discount.Result/100))))
				prd.TotalFinalPrice = float64(v.Qty) * (float64(prod.Price) - float64(prod.Price*prod.Discount.Result)/100)
			}
		}

		total += prd.TotalFinalPrice
		pst = append(pst, prd)
	}

	res := model.SubTotal{}
	res.Products = pst
	res.Subtotal = total

	return res, nil
}

func (c orderUC) List(req model.OrderPaginated) (model.ListOrderResponse, error) {
	data, count, err := c.orderQ.List(req)
	if err != nil {
		return model.ListOrderResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	orders := []model.OrderDetail{}
	for _, v := range data {

		cashier := modelcashier.Cashier{
			CashierID: v.CashiersID,
			Name:      v.CashierName,
		}
		payment := modelpayment.Payment{
			PaymentID: v.PaymentTypesID,
			Name:      v.PaymentName,
			Type:      v.PaymentType,
			Logo:      v.PaymentLogo,
		}
		vv := model.OrderDetail{
			OrderID:     v.OrderID,
			TotalPrice:  v.TotalPrice,
			TotalPaid:   v.TotalPaid,
			TotalReturn: v.TotalReturn,
			ReceiptID:   fmt.Sprintf("ID%03d", v.OrderID),
			UpdatedAt:   v.UpdatedAt.UTC().Format(time.RFC3339),
			CreatedAt:   v.CreatedAt.UTC().Format(time.RFC3339),
		}
		vv.Cashier = cashier
		vv.Payment = payment
		vv.CashiersID = cashier.CashierID
		vv.PaymentTypesID = payment.PaymentID

		orders = append(orders, vv)
	}

	res := model.ListOrderResponse{
		OrderDetails: orders,
		Meta:         meta,
	}

	return res, nil
}
func (c orderUC) Detail(id uint) (model.DetailOrderProductResponse, error) {
	// get order
	data, err := c.orderQ.DetailOrdersById(id)
	if err != nil {
		return model.DetailOrderProductResponse{}, err
	}

	cashier := modelcashier.Cashier{
		CashierID: data.CashiersID,
		Name:      data.CashierName,
	}
	payment := modelpayment.Payment{
		PaymentID: data.PaymentTypesID,
		Name:      data.PaymentName,
		Type:      data.PaymentType,
		Logo:      data.PaymentLogo,
	}
	od := model.OrderProductDetail{
		OrderID:     data.OrderID,
		TotalPrice:  data.TotalPrice,
		TotalPaid:   data.TotalPaid,
		TotalReturn: data.TotalReturn,
		ReceiptID:   fmt.Sprintf("ID%03d", data.OrderID),
		CreatedAt:   data.CreatedAt.UTC().Format(time.RFC3339),
	}
	od.Cashier = cashier
	od.Payment = payment
	od.CashiersID = cashier.CashierID
	od.PaymentTypesID = payment.PaymentID

	// get order details
	listOfOrderDetails, err := c.orderQ.ListOrderDetailsByOrderId(data.OrderID)
	if err != nil {
		return model.DetailOrderProductResponse{}, err
	}

	var prds []model.OrderProductSubTotal
	for _, v := range listOfOrderDetails {
		detail, err := c.productUC.Detail(v.ProductID)
		if err != nil {
			return model.DetailOrderProductResponse{}, err
		}

		prd := model.OrderProductSubTotal{
			ProductID:        detail.ProductID,
			Name:             v.Name,
			Price:            v.Price,
			DiscountID:       v.DiscountID,
			Discount:         detail.Discount,
			Qty:              v.Qty,
			TotalNormalPrice: v.TotalNormalPrice,
			TotalFinalPrice:  v.TotalFinalPrice,
		}

		prds = append(prds, prd)
	}

	orderProductDetail := model.DetailOrderProductResponse{
		OrderDetail: od,
		Products:    prds,
	}

	return orderProductDetail, nil
}

func (c orderUC) CheckDownload(id uint) (model.DownloadStatus, error) {
	order, err := c.orderQ.Detail(id)
	if err != nil {
		return model.DownloadStatus{}, err
	}

	status := model.DownloadStatus{IsDownloaded: false}
	if order.Downloaded != 0 {
		status.IsDownloaded = true
	}
	return status, nil
}

func (c orderUC) Download(id uint) ([]byte, error) {
	bytes, err := c.orderQ.SetDownload(id, 1)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
