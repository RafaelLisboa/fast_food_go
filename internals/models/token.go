package models

type Token struct {
	AcessToken string `json:"access_token"`
	ExpiresIn  uint32	`json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	UserId string
	ExpiresIn uint32
	Token string
}