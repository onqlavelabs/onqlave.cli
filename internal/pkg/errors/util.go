package errors

func NewErrGenerateID(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_GENERATE_ID,
		baseErr: errBase,
		mode:    0,
	}
}

func NewErrInvalidURL(errBase errorI) error {
	return Error{
		service: "Authenticator",
		message: ERR_EXPIRED_TOKEN,
		baseErr: errBase,
		mode:    0,
	}
}

func NewErrExpiredToken(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_INVALID_URL,
		baseErr: errBase,
		mode:    0,
	}
}

func NewErrLoadConfig(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_LOAD_CONFIG,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrEncrypt(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_ENCRYPT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrDecrypt(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_DECRYPT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrCORS(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_CORS,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}
