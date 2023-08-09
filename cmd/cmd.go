package cmd

import (
	"clean_architecture_with_ddd/config"
	"clean_architecture_with_ddd/database"
	"clean_architecture_with_ddd/internal/controller/api"
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/infrastructure/repository"
	"clean_architecture_with_ddd/internal/usecase"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

func Run() *echo.Echo {

	// Read the configuration file
	cfg, err := parseYaml("config.yml")
	panicWithErr(err, fmt.Sprintf("Failed to open the yaml file: %v", err))

	db, err := connectDB(cfg)
	panicWithErr(err, fmt.Sprintf("Failed to connect DB: %v", err))

	e := NewRouter(buildServer(cfg, db))

	return e
}

func panicWithErr(err error, message string) {
	if err != nil {
		panic(message)
	}
}

func parseYaml(path string) (*config.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find the yaml config file")
	}
	decoder := yaml.NewDecoder(file)
	var cfg *config.Config
	if err = decoder.Decode(&cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to parse the yaml file into struct")
	}
	return cfg, nil
}

func connectDB(cfg *config.Config) (*sqlx.DB, error) {
	connectDbStr := mysql.Config{
		DBName:               cfg.DB.DatabaseName,
		User:                 cfg.DB.Username,
		Passwd:               cfg.DB.Password,
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Open("mysql", connectDbStr.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect Database")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping to DB is failed")
	}

	database.CreateTable(db)
	database.SeedTable(db)

	return db, nil
}

func buildServer(_ *config.Config, db *sqlx.DB) (api.AuthHandler, api.CartHandler, api.CartItemHandler, api.OrderHandler, api.PaymentHandler, api.ProductHandler, api.UserHandler) {

	// Build external devices
	repo := repository.NewRepo(db)

	// Put them in each usecase instances
	authUsecase := usecase.NewAuthService(repo)
	cartUsecase := usecase.NewCartService(repo)
	cartItemUsecase := usecase.NewCartItemService(repo)
	orderUsecase := usecase.NewOrderService(repo)
	paymentUsecase := usecase.NewPaymentService(repo)
	productUsecase := usecase.NewProductService(repo)
	userUsecase := usecase.NewUserService(repo)

	// controller middleware instances
	middlewareAuthUsecase := middleware.NewAuth(repo)

	// Create handlers for APIs
	authHandler := api.NewAuthHandler(authUsecase)
	cartHandler := api.NewCartHandler(cartUsecase, middlewareAuthUsecase)
	cartItemHandler := api.NewCartItemHandler(cartItemUsecase, middlewareAuthUsecase)
	orderHandler := api.NewOrderHandler(orderUsecase, middlewareAuthUsecase)
	paymentHandler := api.NewPaymentHandler(paymentUsecase, middlewareAuthUsecase)
	productHandler := api.NewProductHandler(productUsecase)
	userHandler := api.NewUserHandler(userUsecase)

	return authHandler, cartHandler, cartItemHandler, orderHandler, paymentHandler, productHandler, userHandler
}
