package main

import (
    "context"
	"net/http"
    "log"
    "github.com/labstack/echo/v4"
    "github.com/jackc/pgx/v4"
	"GOstarted/routes"
)

func connectToDB() *pgx.Conn {
    conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
    if err != nil {
        log.Fatal("Unable to connect to database:", err)
    }
    return conn
}

func main() {
    e := echo.New()

    // conn := connectToDB()
    // defer conn.Close(context.Background())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Echo and PostgreSQL!")
    })

	routes.SetupRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
