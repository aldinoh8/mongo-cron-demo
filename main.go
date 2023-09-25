package main

import (
	"example/config"
	"example/handler"
	"example/repository"
	"example/schedule"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func main() {
	db := config.InitDb()
	e := echo.New()
	cronJob := cron.New()

	paymentRepository := repository.NewPayment(db)
	paymentHandler := handler.NewPayment(paymentRepository)
	paymentSchedulerHandler := schedule.NewPayment(paymentRepository)

	cronJob.AddFunc("@every 10s", paymentSchedulerHandler.UpdateInvoiceStatus)
	cronJob.Start()

	e.POST("/payment", paymentHandler.Create)
	e.PUT("/invoice/:id", paymentHandler.UpdateInvoiceStatus)

	e.Logger.Fatal(e.Start(":8080"))
}
