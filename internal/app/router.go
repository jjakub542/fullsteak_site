package app

import (
	"fullsteak/internal/admin"
	"fullsteak/internal/public"
	"fullsteak/internal/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) Router() http.Handler {
	e := echo.New()
	e.Use(user.SessionMiddleware(s.store))
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//h := handlers.Handler{Repository: s.repository}
	publicHandler := public.Handler{User: s.repository.User, Article: s.repository.Article}
	adminHandler := admin.Handler{User: s.repository.User, Article: s.repository.Article}
	userHandler := user.Handler{User: s.repository.User}

	e.GET("/", publicHandler.HomePage)
	e.GET("/portfolio", publicHandler.PortfolioPage)
	e.GET("/blog", publicHandler.BlogPage)
	e.GET("/blog/:article_id", publicHandler.ArticleView)

	adminGroup := e.Group("/admin")
	adminGroup.GET("/statistics", user.AdminAuth(adminHandler.AdminStatsPage))
	adminGroup.GET("/articles", user.AdminAuth(adminHandler.AdminHomePage))
	adminGroup.POST("/articles/create", user.AdminAuth(adminHandler.ArticleCreate))
	adminGroup.POST("/articles/:article_id/delete", user.AdminAuth(adminHandler.ArticleDelete))
	adminGroup.POST("/articles/:article_id/update", user.AdminAuth(adminHandler.ArticleUpdate))
	adminGroup.POST("/articles/:article_id/attach-image", user.AdminAuth(adminHandler.ArticleAttachImage))
	adminGroup.POST("/articles/delete-image", user.AdminAuth(adminHandler.ArticleDeleteImage))
	adminGroup.GET("/articles/:article_id/edit", user.AdminAuth(adminHandler.ArticleEditPage))
	adminGroup.Any("/login", userHandler.LoginPage)
	adminGroup.Any("/logout", userHandler.LogoutPage)

	return e
}
