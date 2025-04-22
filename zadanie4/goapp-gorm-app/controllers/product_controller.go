package controllers

import (
    "net/http"
    "strconv"

    "goapp-gorm-app/config"
    "goapp-gorm-app/models"

    echo "github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
    var products []models.Product
    if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, products)
}
func GetProduct(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    if err := config.DB.Preload("Category").First(&product, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie został znaleziony"})
    }
    return c.JSON(http.StatusOK, product)
}
func CreateProduct(c echo.Context) error {
    p := new(models.Product)
    if err := c.Bind(p); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := config.DB.Create(p).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, p)
}
func UpdateProduct(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Product
    if err := config.DB.First(&p, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie został znaleziony"})
    }
    if err := c.Bind(&p); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    config.DB.Save(&p)
    return c.JSON(http.StatusOK, p)
}
func DeleteProduct(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Product
    if err := config.DB.First(&p, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie został znaleziony"})
    }
    config.DB.Delete(&p)
    return c.JSON(http.StatusOK, map[string]string{"message": "Produkt został usunięty"})
}
