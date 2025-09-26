package handler

import (
	"net/http"

	"github.com/Komilov31/l0/internal/model"
	"github.com/gin-gonic/gin"
	json "github.com/goccy/go-json"
	"github.com/google/uuid"
)

type Handler struct {
	orderService model.OrderService
}

func New(service model.OrderService) *Handler {
	return &Handler{orderService: service}
}

func (h *Handler) GetOrderById(c *gin.Context) {
	uid := c.Param("order_uid")
	orderUid, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid uid: "+err.Error())
		return
	}

	order, err := h.orderService.GetOrderById(c, orderUid)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "not order with provided uid"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	data, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "could not unmarshal order"})
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(data)
}

func (h *Handler) GetMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
