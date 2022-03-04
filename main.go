package main

import (
	"log"
	"net/http"

	"box/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/sajanjswl/ethereum-notes/api"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x2B3d254D7958161990Cc17df581D423a33135515"), client)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, reply)
		return nil
	})
	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, reply) // Hello World
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
