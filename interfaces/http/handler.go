package handler

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Handler struct {
}

func NewHandler(logger *logrus.Logger, db *gorm.DB) *Handler {
	return &Handler{}
}
