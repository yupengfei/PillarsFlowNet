package utility

type User struct{
	UserName string
	Password string
}
type Error struct {
	errorCode string
	errorMessage string
}
type OutMessage struct {
	error Error
	result string
}
type InMessage struct {
	Command string
	Parameter interface{}
}

type LoginInMessage struct {
	Command string
	Parameter User
}