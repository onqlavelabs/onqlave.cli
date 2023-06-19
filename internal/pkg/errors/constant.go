package errors

type Mode int64

const (
	DEFAULT Mode = iota
	DEBUG
)

const (
	ERROR_ACTION_NOT_ALLOW = "You are unauthorized to perform this action"
)

// authentication & authorization
const (
	ERR_GENERATE_ID = "Unable to generate ID"
)

// ============================================================
// Error Keys
// ============================================================

const (
	KeyServiceDecryptErr = "service_decrypt_err"
)

const (
	ScopeInternal = "internal"
)
