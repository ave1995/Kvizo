package handlers

import (
	"kvizo-api/internal/auth"
	"kvizo-api/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.AuthenticationRepository
}

func NewAuthHandler(s auth.AuthenticationRepository) *AuthHandler {
	return &AuthHandler{service: s}
}

// RegisterUserHandler godoc
// @Summary Register a new user
// @Description Create a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.RegisterRequest true "User info"
// @Success 201 {object} auth.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUserHandler(c *gin.Context) {
	var req auth.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	user, err := h.service.RegisterUser(c, req.Email, req.Password)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// LoginUserHandler godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User info"
// @Success 201 {object} auth.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) LoginUserHandler(c *gin.Context) {
	var req auth.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	user, err := h.service.AuthenticateUser(c, req.Email, req.Password)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
