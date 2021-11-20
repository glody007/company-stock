package main

import (
  "math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/glody007/eth-contract/api"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/my-stock-qte", func(c echo.Context) error {
		reply, err := conn.StockQte(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(200, reply)
	})

  e.GET("/stock-qte-of/:address", func(c echo.Context) error {
    address := c.Param("address")
    reply, err := conn.StockQteOf(&bind.CallOpts{}, common.HexToAddress(address))
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

  e.GET("/semd-stock-to/:address/:amount", func(c echo.Context) error {
    address := c.Param("address")
    amountString := c.Param("amount")
    amount := new(big.Int)
    amount, ok := amount.SetString(amountString, 10)
    if !ok {
      return  c.JSON(400, "amount error")
    }
    reply, err := conn.SendStockTo(&bind.TransactOpts{}, common.HexToAddress(address), amount)
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

  e.GET("/total-supply", func(c echo.Context) error {
    reply, err := conn.TotalSupply(&bind.CallOpts{})
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

  e.GET("/nbr-owners", func(c echo.Context) error {
    reply, err := conn.NbrOwners(&bind.CallOpts{})
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

  e.GET("/owners", func(c echo.Context) error {
    reply, err := conn.GetOwners(&bind.CallOpts{})
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

  e.GET("/share-dividents", func(c echo.Context) error {
    reply, err := conn.Distribute(&bind.TransactOpts{})
    if err != nil {
      return err
    }
    return c.JSON(200, reply)
  })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
