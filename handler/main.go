package handler

import (
	"example/dto"
	"example/model"
	"example/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	Repository repository.Payment
}

func NewPayment(r repository.Payment) Payment {
	return Payment{Repository: r}
}

func (handler Payment) Create(ctx echo.Context) error {
	body := model.Payment{}

	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "failed to parse request body", Detail: err.Error()})
	}

	invoice, err := handler.Repository.FindByIdInvoice(body.InvoiceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "invoice not found", Detail: err.Error()})
	}

	payment, err := handler.Repository.CreatePayment(invoice.ID, body.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "failed to create payment", Detail: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "success create payment",
		"data":    payment,
	})
}

// invoice/:id
func (handler Payment) UpdateInvoiceStatus(ctx echo.Context) error {
	paramId := ctx.Param("id")

	if paramId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "failed to parse param id"})
	}

	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "failed to parse param id", Detail: err.Error()})
	}

	invoice, err := handler.Repository.FindByIdInvoice(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, dto.ErrorResponse{Message: "invoice not found", Detail: err.Error()})
	}

	_, err = handler.Repository.PaymentFindByInvoiceId(invoice.ID)
	if err != nil {
		// payment gaada yaudah berarti update status blm perlu dilakukan
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "payment not received",
			"invoice": invoice,
		})
	}

	// update status perlu dilakukan
	err = handler.Repository.PaidInvoice(invoice.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Message: "failed to update status", Detail: err.Error()})
	}

	invoice.Status = "PAID"
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "invoice paid",
		"invoice": invoice,
	})
}
