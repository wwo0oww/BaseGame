package error

type XError struct {
	code int32
	str  string
}

func (self *XError) Error() string {
	return self.str
}

func (self *XError) Code() int32 {
	return self.code
}

func New(str string, code int32) *XError {
	err := &XError{code: code, str: str}
	return err
}
