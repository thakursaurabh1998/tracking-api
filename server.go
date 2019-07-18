package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"tracking-api/controllers"
	"tracking-api/database"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	appPort := flag.String("port", os.Getenv("SERVER_PORT"), "App running on a port")
	flag.Parse()

	db := InitDB()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	defer db.Close()

	database.InitService{db}.Init()
	// database.Test()

	e := echo.New()

	controllers.HomeController{db}.Init(e.Group("/"))

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowCredentials: true,
	// 	AllowMethods:     []string{http.MethodPost},
	// 	AllowOrigins:     []string{"https://localhost:8000"},
	// 	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXCSRFToken, "X-SC-Broker", "X-SC-JWT", "Cache-Control"},
	// }))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf("%s %d %s\n", c.Request().Method, c.Response().Status, c.Request().URL)
	}))

	e.Debug = true

	if err := e.Start(":" + *appPort); err != nil {
		log.Fatal(err)
	}
}

// InitDB initializes database
func InitDB() *sql.DB {
	dbUser, dbName, host, password := os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", dbUser, password, host, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
