package usecase

import "github.com/bondhan/ecommerce/modules/order/model"

type IOrderUC interface {
	SubTotal(req []model.SubTotalReq) (model.SubTotal, error)
	Create(req model.OrderReq) (model.OrderTotalResp, error)
	List(req model.OrderPaginated) (model.ListOrderResponse, error)
	Detail(id uint) (model.DetailOrderProductResponse, error)
	CheckDownload(id uint) (model.DownloadStatus, error)
	Download(id uint) ([]byte, error)
}
