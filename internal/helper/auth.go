package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


type Auth struct {
	Secret string
}


func (a *Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return "", errors.New("password hashing failed")
	}

	return string(hashedPassword), nil
}


func (a *Auth) GenerateToken(id uint, email, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs are missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("error signing token")
	}

	return tokenStr, nil
}


func (a *Auth) VerifyPassword(password, hashedPassword string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}


func (a *Auth) VerifyToken(authHeader string) (dto.UserLocals, error) {
	tokenSlice:= strings.Fields(authHeader)

	if len(tokenSlice) < 2 || tokenSlice[0] != "Bearer" {
		return dto.UserLocals{}, errors.New("malformed auth header")
	}

	token, err := jwt.Parse(
		tokenSlice[1],
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", t.Header)
			}
			return []byte(a.Secret), nil
		},
	)

	if err != nil {
		return dto.UserLocals{}, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		if exp, ok := claims["exp"].(float64); ok && exp < float64(time.Now().Unix()) {
			return dto.UserLocals{}, errors.New("token expired")
		}

		user := dto.UserLocals{
			ID: uint(claims["user_id"].(float64)),
			Email: claims["email"].(string),
			UserType: claims["role"].(string),
		}
		return user, nil
	}

	return dto.UserLocals{}, errors.New("invalid token")
}


func (a *Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "missing auth header",
		})
	}

	user, err := a.VerifyToken(authHeader)

	if err != nil || user.ID < 1 {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason": err.Error(),
		})
	}

	ctx.Locals("user", user)

	return ctx.Next()
}


func (a *Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")

	return user.(domain.User)
}


func (a *Auth) GenerateCode() (int, error) {

	return RandomNumbers(6)
}


func (a *Auth) AuthorizeSeller(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "missing auth header",
		})
	}

	user, err := a.VerifyToken(authHeader)

	if err != nil || user.ID < 1 {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason": err.Error(),
		})
	}

	if user.UserType != domain.SELLER {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "unauthorized user",
			"reason": "you are not a seller",
		})
	}

	ctx.Locals("user", user)

	return ctx.Next()
}