package app

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"fullsteak/internal/article"
	"fullsteak/internal/contact"
	"fullsteak/internal/user"
)

type Repository struct {
	User    user.Repository
	Article article.Repository
	Contact contact.Repository
}

type Server struct {
	port       int
	db         *pgxpool.Pool
	store      *user.SessionStore
	repository *Repository
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := PostgresClient()
	store := user.NewSessionStore()

	NewServer := &Server{
		port:  port,
		db:    db,
		store: store,
		repository: &Repository{
			User:    user.NewRepository(db),
			Article: article.NewRepository(db),
			Contact: contact.NewRepository(db),
		},
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
