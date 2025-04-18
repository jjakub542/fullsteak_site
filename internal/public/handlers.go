package public

import (
	"net/http"
	"strconv"

	"fullsteak/internal/article"
	"fullsteak/internal/user"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	User    user.Repository
	Article article.Repository
}

func (h *Handler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (h *Handler) PortfolioPage(c echo.Context) error {
	return c.Render(http.StatusOK, "portfolio.html", nil)
}

func (h *Handler) BlogPage(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit := 8
	offset := (page - 1) * limit
	articles, err := h.Article.GetAllPublicBetween(limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	count, err := h.Article.GetCount()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	pagesCount := count/limit + 1

	return c.Render(http.StatusOK, "blog.html", map[string]interface{}{
		"articles":      articles,
		"articlesCount": count,
		"pagesCount":    pagesCount,
		"page":          page,
		"previous":      page - 1,
		"next":          page + 1,
	})
}

func (h *Handler) ArticleView(c echo.Context) error {
	id := c.Param("article_id")
	article, err := h.Article.GetOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "article.html", article)
}
