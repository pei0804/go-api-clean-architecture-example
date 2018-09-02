package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	accountController "github.com/pei0804/go-api-clean-architecture-example/account/controller"
	accountRepository "github.com/pei0804/go-api-clean-architecture-example/account/repository"
	accountUsecase "github.com/pei0804/go-api-clean-architecture-example/account/usecase"
)

func main() {
	dbConn, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	r := chi.NewRouter()
	accountRepo := accountRepository.NewAccountRepository(dbConn)
	accountUcase := accountUsecase.NewAccountUsecase(accountRepo)
	accountController.NewAccountController(r, accountUcase)
	http.ListenAndServe(":8080", r)
}
