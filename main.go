package main

import (
	"Hexagonal/handler"
	"Hexagonal/logs"
	"Hexagonal/repository"
	"Hexagonal/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	db := initDatabase()

	CustomerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	// CustomerRepositoryMock := repository.NewCustomerRepositoryMock()
	CustomerService := service.NewCustomerService(CustomerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(CustomerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	logs.Info("Listening on port " + fmt.Sprintf(":%v", viper.GetInt("app.port")))
	http.ListenAndServe(fmt.Sprintf(":%v",
		viper.GetInt("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)

	return db
}