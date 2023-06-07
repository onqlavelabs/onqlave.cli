package errors

func NewErrConnectDB(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_CONNECT_TO_DB,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrMigration(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_MIGRATION,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrCommitTransaction(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_COMMIT_TRANSACTION,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrRollBackTransaction(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_ROLLBACK_TRANSACTION,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrSeeding(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_SEEDING,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrInvalidQuery(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_INVALID_QUERY,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrDuplicateID(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_DUPLICATE_ID,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrNoResult(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_NO_RESULT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}
