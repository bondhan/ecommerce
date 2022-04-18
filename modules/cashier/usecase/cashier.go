package usecase

type ICashierUC interface {
	List()
	Detail()
	Create()
	Update()
	Delete()
}
