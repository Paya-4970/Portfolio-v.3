package handlers

import (
	"github.com/Paya-4970/Portfolio-v.3/internal/database"
	"github.com/Paya-4970/Portfolio-v.3/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func ProductsHandle(c echo.Context) error {
	var productAll []models.Product
	database.DB.Where(&models.Product{Name: "test1"}).Find(&productAll)

	context := map[string]interface{}{
		"Products": productAll,
	}

	return c.Render(http.StatusOK, "products.html", context)
}

func ShopHandle(c echo.Context) error {
	var shopAll []models.Shop
	database.DB.Find(&shopAll)

	context := map[string]interface{}{
		"Shops": shopAll,
	}
	return c.Render(http.StatusOK, "shop.html", context)
}
