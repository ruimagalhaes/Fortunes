package main

import (
	"database/sql"
	"fmt"
	"main/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database, err := startDB()
	if err != nil {
		fmt.Errorf("database: %v", err)
		return
	}

	e := echo.New()
	e.Static("/static", "./static")
	e.Use(middleware.Logger())
	e.Use(databaseMiddleware(database))

	fortunesHandler := handler.FortunesHandler{}
	e.GET("/", fortunesHandler.HandleGetFortuneList)
	e.GET("/fortunes/:id", fortunesHandler.HandleGetFortune)
	e.PUT("/fortunes/:id", fortunesHandler.HandleOpenFortune)
	e.GET("/fortunes/form", fortunesHandler.HandleNewFortunes)
	e.POST("/fortunes", fortunesHandler.HandlePostFortunes)
	e.DELETE("/fortunes/", fortunesHandler.HandleDeleteFortunes)

	e.Logger.Fatal(e.Start(":8082"))
}

func startDB() (*sql.DB, error) {
	var err error

	var db *sql.DB
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASS") + "@/" + os.Getenv("DBNAME") + "?parseTime=true"
	fmt.Printf("Connecting with DSN: %s\n", dsn)
	db, err = sql.Open("mysql", dsn)
	// db, err = sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@/"+os.Getenv("DBNAME")+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}

func databaseMiddleware(database *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("database", database)
			return next(c)
		}
	}
}
