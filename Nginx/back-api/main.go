package main

import (
	"commerce-api/database"
	"commerce-api/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func getProducts(c echo.Context) error {
	reponse, err := database.List()
	if err != nil {
		return err
	}
	return c.JSON(200, reponse)
}

func addProducts(c echo.Context) error {
	var p models.Product
	err := json.NewDecoder(c.Request().Body).Decode(&p)
	if err != nil {
		return err
	}
	err = database.Insert(p)
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func removeProducts(c echo.Context) error {
	req := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&req)

	if err != nil {
		return c.JSON(400, "invalid input")
	}

	switch req["id"].(type) {
	case float64:
		id, ok := req["id"].(float64)
		if !ok {
			return c.JSON(500, "convertion error")
		}
		err = database.Remove(int(id))
		if err != nil {
			return err
		}
		return c.JSON(200, "removed")
	case int:
		id, ok := req["id"].(int)
		if !ok {
			return c.JSON(500, "convertion error")
		}
		err = database.Remove(id)
		if err != nil {
			return err
		}
		return c.JSON(200, "removed")
	default:
		return c.JSON(400, "invalid input")
	}
}

func updateProducts(c echo.Context) error {
	var p models.Product
	err := json.NewDecoder(c.Request().Body).Decode(&p)
	if err != nil {
		return err
	}
	err = database.Update(p)
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func main() {
	database.OpenConnection()
	api := echo.New()

	api.GET("/products", getProducts)
	api.POST("/products/add", addProducts)
	api.DELETE("/products/remove", removeProducts)
	api.POST("/products/update", updateProducts)
	api.Logger.Fatal(api.Start(":8085"))
}
