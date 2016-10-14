package utils

//Error Wapper error
type _error struct {
	err error
	msg string
}

func (e *_error) Error() string {
	if e.err != nil {
		return e.err.Error() + "->" + e.msg
	}
	return e.msg
}

//WrapError implements wrapper error given with msg info for context notice
func WrapError(e error, m string) error {
	return &_error{err: e, msg: m}
}
