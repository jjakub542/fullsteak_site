package admin

import (
	"fullsteak/internal/article"
	"fullsteak/internal/contact"
	"fullsteak/internal/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	User    user.Repository
	Article article.Repository
	Contact contact.Repository
}

func (h *Handler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/home.html", nil)
}

func (h *Handler) StatsPage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/statistics.html", nil)
}

func (h *Handler) MessagesPage(c echo.Context) error {
	messages, err := h.Contact.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Render(http.StatusOK, "admin/messages.html", messages)
}
