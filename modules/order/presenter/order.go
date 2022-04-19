package presenter

import (
	"net/http"
)

type IOrderP interface {
	SubTotal(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Download(w http.ResponseWriter, r *http.Request)
	DownloadStatus(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}
