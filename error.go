package goerror

type Error struct {
	errno string
	msg   string
}

func New(errno, msg string) *Error {
	return &Error{
		errno: errno,
		msg:   msg,
	}
}

func (e *Error) Error() string {
	result := "errno: " + e.errno + ", "
	result += "msg: " + e.msg

	return result
}

func (e *Error) Errno() string {
	return e.errno
}

func (e *Error) Msg() string {
	return e.msg
}
