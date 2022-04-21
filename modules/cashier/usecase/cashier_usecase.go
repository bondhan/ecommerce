package usecase

import (
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/status"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/modules/cashier/model"
	"github.com/bondhan/ecommerce/modules/cashier/query"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type cashierUC struct {
	jwtKey   string
	logger   *logrus.Logger
	cashierQ query.ICashierQ
}

func NewCashierUC(logger *logrus.Logger, jwtKey string, cashierQ query.ICashierQ) ICashierUC {
	return &cashierUC{
		logger:   logger,
		jwtKey:   jwtKey,
		cashierQ: cashierQ,
	}
}
func (c cashierUC) Create(req model.CreateCashierReq) (model.CreateCashierResp, error) {
	newCashier, err := c.cashierQ.Insert(req)
	if err != nil {
		return model.CreateCashierResp{}, err
	}

	nCashier := model.CreateCashierResp{
		Cashier: model.Cashier{
			CashierID: newCashier.ID,
			Name:      newCashier.Name,
		},
		PassCode:  newCashier.Passcode,
		CreatedAt: newCashier.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: newCashier.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return nCashier, nil
}

func (c cashierUC) Update(req model.CreateCashierUpdate) error {
	err := c.cashierQ.Update(req)
	if err != nil {
		return err
	}

	return nil
}

func (c cashierUC) Delete(id uint) error {
	err := c.cashierQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c cashierUC) List(req model.CashierPaginated) (model.ListResponse, error) {
	data, count, err := c.cashierQ.List(req)
	if err != nil {
		return model.ListResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	cashiers := []model.Cashier{}
	for _, v := range data {
		vv := model.Cashier{
			Name:      v.Name,
			CashierID: v.ID,
		}
		cashiers = append(cashiers, vv)
	}

	res := model.ListResponse{
		Cashiers: cashiers,
		Meta:     meta,
	}

	return res, nil
}
func (c cashierUC) Detail(id uint) (model.Cashier, error) {
	data, err := c.cashierQ.Detail(id)
	if err != nil {
		return model.Cashier{}, err
	}

	cashier := model.Cashier{
		Name:      data.Name,
		CashierID: data.ID,
	}

	return cashier, nil
}

func (c cashierUC) PassCode(id uint) (model.Passcode, error) {
	data, err := c.cashierQ.Detail(id)
	if err != nil {
		return model.Passcode{}, err
	}

	passcode := model.Passcode{
		Passcode: data.Passcode,
	}

	return passcode, nil
}

func (c cashierUC) Login(id uint, passcode model.CreatePasscodeReq) (model.Token, error) {
	data, err := c.cashierQ.Detail(id)
	if err != nil {
		return model.Token{}, err
	}

	if data.Passcode != passcode.PassCode {
		return model.Token{}, ecommerceerror.ErrPasscodeNotMatch
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	//expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &basemodel.Claims{
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Subject:  strconv.Itoa(int(data.ID)),
			IssuedAt: time.Now().Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(c.jwtKey))
	if err != nil {
		return model.Token{}, err
	}

	err = c.cashierQ.UpdateLogin(id, status.LoggedIn)
	if err != nil {
		return model.Token{}, err
	}

	tkn := model.Token{
		Token: tokenString,
	}

	return tkn, nil
}

func (c cashierUC) Logout(id uint, passcode model.CreatePasscodeReq) error {
	data, err := c.cashierQ.Detail(id)
	if err != nil {
		return err
	}

	if data.Passcode != passcode.PassCode {
		return ecommerceerror.ErrPasscodeNotMatch
	}

	err = c.cashierQ.UpdateLogin(id, status.LoggedOut)
	if err != nil {
		return err
	}

	return nil
}
