package restapp

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// RefreshContent is a struct containing the decoded refresh token
type RefreshContent struct {
	UserID   string
	TokenID  string
	UserMail string
	jwt.RegisteredClaims
}

func (r *RefreshContent) ToDomain(m map[string]interface{}) {
	r.UserID = m["UserID"].(string)
	r.TokenID = m["TokenID"].(string)
	r.UserMail = m["UserMail"].(string)
	return
}

func (r *RefreshContent) SetClaims(claims jwt.RegisteredClaims) {
	r.RegisteredClaims = claims
}

var RefreshTokenDuration = time.Hour * 24 * 30 * 2 // Two month

// AccessContent is a struct containing the decoded access token
type AccessContent struct {
	UserMail string
	UserName string
	jwt.RegisteredClaims
}

func (a *AccessContent) ToDomain(m map[string]interface{}) {
	a.UserMail = m["UserMail"].(string)
	a.UserName = m["UserName"].(string)
	return
}

func (a *AccessContent) SetClaims(claims jwt.RegisteredClaims) {
	a.RegisteredClaims = claims
}

var AccessTokenDuration = time.Hour * 24 // One day
