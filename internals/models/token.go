package models

type Token struct {
	AcessToken string `json:"access_token"`
	ExpiresIn  uint32	`json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
