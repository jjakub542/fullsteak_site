package handlers

import (
	"fullsteak/internal/repository"
)

type Handler struct {
	Repository *repository.Repository
}
