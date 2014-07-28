package utility

type User struct{
	UserName string
	Password string
}
type Error struct {
	ErrorCode int
	ErrorMessage string
}
type OutMessage struct {
	Error Error
	Command string
	Result string
}

type InMessage struct {
	Command string
	Parameter string
}


type LoginInMessage struct {
	Auth string
	AuthMessage string
}