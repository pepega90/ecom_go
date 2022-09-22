package order

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/order/usecases"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

type OrderHandler struct {
	orderUse usecases.OrderUseCase
}

func NewOrderHandler(orderUse usecases.OrderUseCase, r *fiber.App) *OrderHandler {
	handler := &OrderHandler{
		orderUse: orderUse,
	}
	r.Get("/orders", handler.GetAllOrders)
	r.Get("/create-snap-token", handler.CreateSNAP)
	r.Post("/create-order", handler.CreateOrder)
	r.Get("/check-order/:id", handler.CheckOrder)
	r.Delete("/hapus-order/:id", handler.DeleteOrder)
	return handler
}

// TODO: selesaikan handler http order

func (od *OrderHandler) CreateSNAP(c *fiber.Ctx) error {
	s := snap.Client{}
	s.New(midtrans.ServerKey, midtrans.Sandbox)

	// find user
	currentUserID, err := strconv.Atoi(c.GetRespHeaders()["User_id"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User dengan id %d tidak ditemukan. error : %s", currentUserID, err),
		})
	}
	user, _ := od.orderUse.GetPreloadData(uint(currentUserID))

	var items []midtrans.ItemDetails
	total := 0
	for _, v := range user.Carts {
		prod := midtrans.ItemDetails{
			ID:    strconv.Itoa(int(v.CarItemID)),
			Name:  v.CartItems.Product.Title,
			Price: int64(v.CartItems.Product.Price),
			Qty:   int32(v.CartItems.Qty),
		}
		items = append(items, prod)
		total += v.CartItems.Qty * v.CartItems.Product.Price
	}

	// midtrans payment
	req := &snap.Request{TransactionDetails: midtrans.TransactionDetails{
		OrderID:  "MID-GO-ID-" + example.Random(),
		GrossAmt: int64(total),
	},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			LName: user.Name,
			Email: user.Email,
			Phone: strconv.Itoa(user.Phone),
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeIndomaret,
			snap.PaymentTypeAlfamart,
		},
		Items: &items,
	}

	resp, err := s.CreateTransaction(req)

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (od *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	listOrders, err := od.orderUse.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(listOrders)
}

func (od *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var orderDto OrderReq
	if err := c.BodyParser(&orderDto); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	currentUserID, err := strconv.Atoi(c.GetRespHeaders()["User_id"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User dengan id %d tidak ditemukan. error : %s", currentUserID, err),
		})
	}

	order := domain.Order{
		Total:       int(orderDto.GrossAmt),
		TransaksiID: orderDto.OrderID,
		PaymentType: orderDto.PaymentType,
		VaNumber:    orderDto.VaNumber,
		PdfUrl:      orderDto.PdfUrl,
		StatusCode:  orderDto.StatusCode,
		UserID:      uint(currentUserID),
	}
	created, err := od.orderUse.Create(&order)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("order not created! %s", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(created)
}

func (od *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	err = od.orderUse.Delete(uint(id))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus order",
	})
}

func (od *OrderHandler) CheckOrder(c *fiber.Ctx) error {
	s := coreapi.Client{}
	s.New(midtrans.ServerKey, midtrans.Sandbox)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	order, _ := od.orderUse.GetById(uint(id))

	res, _ := s.CheckTransaction(order.TransaksiID)

	if res.StatusCode == "200" {
		od.orderUse.CheckOrder(order)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
