package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pei0804/go-api-clean-architecture-example/account"
)

type Controller struct {
	AccountUsecase account.Usecase
}

func NewAccountController(router *chi.Mux, accountUsecase account.Usecase) {
	c := &Controller{
		AccountUsecase: accountUsecase,
	}
	router.Route("/accounts", func(r chi.Router) {
		r.Get("/", c.GetList)
	})
}

func (a *Controller) GetList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
