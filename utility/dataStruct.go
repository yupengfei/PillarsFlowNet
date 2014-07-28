package utility

type User struct{
	UserName string
	Password string
}
type Error struct {
	ErrorCode string
	ErrorMessage string
}
type OutMessage struct {
	Error Error
	Result string
}
type InMessage struct {
	Command string
	Parameter string
}

type LoginInMessage struct {
	Command string
	Parameter User
}