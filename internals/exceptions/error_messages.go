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
	EMPTY_REQUIRED_FIELD
	LOGIN_FAILED
)

var errorsMessages = [...]ErrorMessage{
	{id: USER_ALREADY_EXISTS, message: "User already exists"},
	{id: TOKEN_ERROR, message: "Your token is invalid"},
	{id: INVALID_USER, message: "Missed data on user creation body request"},
	{id: INTERNAL_ERROR, message: "An unexpected error has ocurred"},
	{id: EMPTY_REQUIRED_FIELD, message: "The required field %s is empty"},
	{id: LOGIN_FAILED, message: "Your email or passoword is incorrent"},
}
