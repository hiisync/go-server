package handlers

import (
	"lunar-server/internal/database"
	"lunar-server/internal/database/models"
	"lunar-server/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

// CreateUser creates a new user based on the JSON data provided in the request.
//
// It takes a gin Context as a parameter, which contains the HTTP request and response.
// It returns a JSON response with the created user data or an error message.
func CreateUser(c *gin.Context) {
	var userRequest UserCreateRequest

	// Bind JSON data to userRequest struct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user model from the request
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Password: hashedPassword,
	}

	// Create user in the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully!",
		"user": map[string]interface{}{
			"name":     user.Name,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func UserLogin(c *gin.Context) {
	// Get email & pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up for requested user
	var user models.User

	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": map[string]interface{}{
			"id":         user.(models.User).ID,
			"username":   user.(models.User).Username,
			"email":      user.(models.User).Email,
			"name":       user.(models.User).Name,
			"created_at": user.(models.User).CreatedAt,
			"updated_at": user.(models.User).UpdatedAt,
			"deleted_at": user.(models.User).DeletedAt,
		},
	})
}
