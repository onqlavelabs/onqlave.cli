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
	KeyHandlerPayloadErr = "handler_payload_err"
	KeyHandlerDataErr    = "handler_data_err"
	KeyHandlerError      = "handler_logic_err"
)

const (
	KeyServiceErr                       = "service_logic_err"
	KeyServiceRevokeEmailErr            = "service_revoke_email_err"
	KeyServiceRegisterCompletedErr      = "service_register_completed_err"
	KeyServiceRegisterNotActionedErr    = "service_register_not_actioned_err"
	KeyServiceEncryptErr                = "service_encrypt_err"
	KeyServiceDecryptErr                = "service_decrypt_err"
	KeyServicePublishEventErr           = "service_publish_event_err"
	KeyServiceUnauthorizedErr           = "service_unauthorized_err"
	KeyServiceDuplicateIDErr            = "service_duplicate_id_err"
	KeyServiceDefaultClusterNotExistErr = "service_default_cluster_not_exist"
)

const (
	KeyDbContextErr   = "db_context_err"
	KeyDbQueryErr     = "db_query_err"
	KeyDbInsertErr    = "db_insert_err"
	KeyDbUpdateErr    = "db_update_err"
	KeyDbDeleteErr    = "db_delete_err"
	KeyDbEmptyErr     = "db_empty_err"
	KeyDbDuplicateErr = "db_duplicate_err"
	KeyDbNotExistErr  = "db_not_exist_err"
	KeyDbCountErr     = "key_count_err"
)

const (
	ScopeHandler    = "handler"
	ScopeService    = "service"
	ScopeRepository = "repository"
	ScopeInternal   = "internal"
)
