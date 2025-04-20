package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userId int, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetDataByToken(token string) (int, string, error)
}

type jwtUserClaim struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "tugas-rpl-cuy",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "tugas-rpl-cuy"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userId int, role string) string {
	claims := jwtUserClaim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetDataByToken(token string) (int, string, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return 0, "", err
	}

	claims, ok := t_Token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("cannot parse claims")
	}

	idFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("user_id not found or not valid")
	}
	userID := int(idFloat)

	// Ambil role
	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", fmt.Errorf("role not found or not valid")
	}

	return userID, role, nil
}
