package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{
		db: d,
	}
}

func (h *Handler) Signup(c *gin.Context) {
	var req CreateUserReq

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Faild tou read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Faild tou hash password"})
		return
	}

	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
	}

	result := h.db.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Faild to create user"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginUserReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faild to read body"})
		return
	}
	var user User
	h.db.First(&user, "email = ?", req.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("alksdjf9182374laksjdfh"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tou create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", tokenString, 3600*24*30, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "all done"})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "logout successful"})
}
