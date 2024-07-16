package exceptions

type ErrorMessage struct {
	id      int
	message string
}

const (
	ENUM = iota
	USER_ALREADY_EXISTS
	TOKEN_ERROR
	INVALID_USER
	INTERNAL_ERROR
)


var errors = [...]ErrorMessage {
	{id: USER_ALREADY_EXISTS, message: "User already exists"},
	{id: TOKEN_ERROR, message: "Your token is invalid"},
	{id: INVALID_USER,  message: "Missed data on user creation body request"},
	{id: INTERNAL_ERROR, message: "An unexpected error has ocurred"},
}

