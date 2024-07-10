package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/akposieyefa/open-ai/config"
	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/core/types"
	"github.com/akposieyefa/open-ai/internals"
	"github.com/akposieyefa/open-ai/pkg/hash"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(config.Envs.JWT_SECRET)

// sigin in user with email and password
func GetAuthToken(email, password string) (string, error) {
	user := &models.User{}

	if err := internals.DB.Where("Email = ?", email).First(user).Error; err != nil {
		return "", fmt.Errorf("email address not found: %v", err)
	}

	check, err := hash.CheckPassword(user.Password, password)
	if err != nil {
		return "", fmt.Errorf("error checking password: %v", err)
	}
	if !check {
		return "", errors.New("invalid password")
	}

	ttl, _ := strconv.Atoi(config.Envs.JWT_TTL)

	expirationTime := time.Now().Add(time.Minute * time.Duration(ttl)).Unix()
	tk := &models.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(config.Envs.JWT_ALGO), tk)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %v", err)
	}

	return tokenString, nil
}

// get auth user
func AuthUser(request *http.Request) (types.UserResponse, error) {

	tokenString := request.Header.Get("Authorization") //get token string

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return types.UserResponse{}, errors.New("invalid token signature")
		}
		return types.UserResponse{}, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return types.UserResponse{}, errors.New("token is not valid")
	}

	user, err := getUserByEmail(claims.Email)
	userResponse := types.UserResponse{
		ID:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		PhoneNumber:     user.PhoneNumber,
		RoleId:          user.RoleId,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
	}
	if err != nil {
		return types.UserResponse{}, fmt.Errorf("error retrieving user: %v", err)
	}

	return userResponse, nil
}

// get auth user by email
func getUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := internals.DB.Where("Email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
