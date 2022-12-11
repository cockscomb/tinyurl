package controller

import "github.com/cockscomb/tinyurl/usecase"

type Controller struct {
	tinyurl *usecase.TinyURLUsecase
}

func NewController(tinyurl *usecase.TinyURLUsecase) *Controller {
	return &Controller{tinyurl: tinyurl}
}
