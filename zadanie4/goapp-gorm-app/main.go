package main

import (
    "log"
    "net/http"

    "goapp-gorm-app/config"
    "goapp-gorm-app/controllers"
    "goapp-gorm-app/models"

    echo "github.com/labstack/echo/v4"
)

func main() {
    db := config.InitDB()
    db.AutoMigrate(
        &models.Product{},
        &models.Category{},
        &models.Customer{},
        &models.Order{},
        &models.OrderItem{},
    )

    e := echo.New()
    e.GET("/products", controllers.GetProducts)
    e.GET("/products/:id", controllers.GetProduct)
    e.POST("/products", controllers.CreateProduct)
    e.PUT("/products/:id", controllers.UpdateProduct)
    e.DELETE("/products/:id", controllers.DeleteProduct)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Aplikacja działa poprawnie!")
    })

    log.Fatal(e.Start("0.0.0.0:8080"))
}