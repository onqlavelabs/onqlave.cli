package errors

func NewErrGenerateID(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_GENERATE_ID,
		baseErr: errBase,
		mode:    0,
	}
}
