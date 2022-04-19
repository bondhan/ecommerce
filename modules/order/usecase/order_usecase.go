package usecase

import (
	"fmt"
	"github.com/bondhan/ecommerce/constants/params"
	"github.com/bondhan/ecommerce/modules/order/model"
	"github.com/bondhan/ecommerce/modules/order/query"
	queryproduct "github.com/bondhan/ecommerce/modules/product/query"
	usecaseproduct "github.com/bondhan/ecommerce/modules/product/usecase"
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
		CashiersID:     uint(1), //ToDo: get from jwt
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
			CashiersID:     res.CashierID,
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
	var total int64
	pst := []model.ProductSubTotal{}
	for _, v := range req {
		prod, err := c.productUC.Detail(v.ProductID)
		if err != nil {
			return model.SubTotal{}, err
		}

		prd := model.ProductSubTotal{
			ProductID:        prod.ProductID,
			Name:             prod.Name,
			Stock:            prod.Stock,
			Price:            prod.Price,
			Image:            prod.Image,
			Discount:         prod.Discount,
			Qty:              v.Qty,
			TotalNormalPrice: v.Qty * prod.Price,
			TotalFinalPrice:  v.Qty * prod.Price,
		}

		if prod.Discount != nil {
			if prod.Discount.Type == params.BuyN {
				multiplier := v.Qty / prod.Discount.Qty
				remains := v.Qty % prod.Discount.Qty
				prd.TotalFinalPrice = multiplier*prod.Discount.Result + remains*prod.Price
			} else if prod.Discount.Type == params.Percentage {
				prd.TotalFinalPrice = v.Qty * (prod.Price - prod.Price*prod.Discount.Result/100)
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

func (c orderUC) Delete(id uint) error {
	err := c.orderQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c orderUC) List(req model.OrderPaginated) (model.ListResponse, error) {
	//data, count, err := c.orderQ.List(req)
	//if err != nil {
	//	return model.ListResponse{}, err
	//}
	//
	//meta := basemodel.Meta{
	//	Total: count,
	//	Skip:  req.Skip,
	//	Limit: req.Limit,
	//}
	//
	//orders := []model.Order{}
	//for _, v := range data {
	//	vv := model.Order{
	//		Name:    v.TotalPrice,
	//		OrderID: v.ID,
	//	}
	//	orders = append(orders, vv)
	//}
	//
	//res := model.ListResponse{
	//	Orders: orders,
	//	Meta:   meta,
	//}

	return model.ListResponse{}, nil
}
func (c orderUC) Detail(id uint) (model.Order, error) {
	data, err := c.orderQ.Detail(id)
	if err != nil {
		return model.Order{}, err
	}

	order := model.Order{
		OrderID: data.ID,
	}

	return order, nil
}
