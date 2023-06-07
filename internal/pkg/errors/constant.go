package errors

type Mode int64

const (
	DEFAULT Mode = iota
	DEBUG
)

// Database
const (
	ERROR_CONNECT_TO_DB = "Connect to database failed"
	ERROR_MIGRATION     = "Unable to migrate"
	ERROR_SEEDING       = "Unable to seeding database"
	ERROR_INVALID_QUERY = "Invalid Query"
	ERROR_DUPLICATE_ID  = "Duplicate ID"
	ERROR_NO_RESULT     = "No record in database"

	ERROR_COMMIT_TRANSACTION   = "Commit database transaction failed"
	ERROR_ROLLBACK_TRANSACTION = "Unable to commit transaction"
)

// service
const (
	ERROR_INIT_SERVICE    = "Init service failed"
	ERROR_PUBLISH_EVENT   = "Unable to publish event"
	ERROR_PUBLISH_COMMAND = "Unable to publish command"
	ERROR_CREATE_COMMAND  = "Unable to create command"

	ERROR_INVALID_PAYLOAD = "Invalid payload"
	ERROR_EXTRACT_DATA    = "Unable to extract data"
)

const (
	ERROR_ACTION_NOT_ALLOW = "You are unauthorized to perform this action"
)

// keyGen service
const (
	ERROR_CREATE_AUTH_TOKEN = "Create auth token failed"
)

// firebase service
const (
	ERROR_FIREBASE_INIT                   = "Unable to init firebase client"
	ERROR_FIREBASE_ADD_CLAIM              = "Unable to add claim to user"
	ERROR_FIREBASE_UPDATE_USER            = "Unable to update user"
	ERROR_FIREBASE_DELETE_USER            = "Unable to delete user"
	ERROR_FIREBASE_VERIFY_USER            = "Unable to verify user"
	ERROR_FIREBASE_CREATE_USER            = "Unable to create user"
	ERROR_FIREBASE_USER_NOT_FOUND         = "User not exist"
	ERROR_FIREBASE_EMAIL_ALREADY_VERIFIED = "User email already verified"
	ERROR_FIREBASE_VERIFY_EMAIL_LINK      = "Unable to get verify email link"
	ERROR_FIREBASE_RESET_PASSWORD_LINK    = "Unable to get reset password link"
	ERROR_FIREBASE_INVITATION_LINK        = "Unable to get invitation link"
)

// payment service
const (
	ERROR_PAYMENT_UPDATE_CUSTOMER       = "Unable to update customer"
	ERROR_PAYMENT_ATTACH_PAYMENT_METHOD = "Unable to attach payment method"
	ERROR_PAYMENT_DETACH_PAYMENT_METHOD = "Unable to detach payment method"
)

// Communication service
const (
	ERROR_OPERATION_NOT_SUPPORTED = "Operation not supported"
	ERROR_SEND_MAIL               = "Unable to send email"
	ERROR_SEND_MAIL_TIMEOUT       = "Already sent a mail to your email"
)

// Tenant service
const (
	ERROR_USER_DISABLE   = "Unable to disable user"
	ERROR_USER_ENABLE    = "Unable to enable user"
	ERROR_TENANT_DISABLE = "Unable to enable tenant"
	ERROR_USER_INACTIVE  = "Unable to take action since user is inactive"
)

// Registration service
const (
	ERROR_USER_NOT_FOUND        = "User is not allowed to login "
	ERROR_REGISTER              = "Process signup or login on provided email failed"
	ERROR_LOGIN_TIMEOUT         = "Please wait some time before login again"
	ERROR_REGISTER_COMPLETED    = "The request cannot be verified as it is already completed"
	ERROR_REGISTER_NOT_ACTIONED = "The request cannot be verified as it is not actioned by user yet"
	ERROR_REVOKE_EMAIL          = "There is already an operation on going"
)

// Billing service
const (
	ERROR_MAXIMUM_PAYMENT_METHOD = "Maximum number of payment methods reached"
)

// Cluster service
const (
	ERR_UNABLE_SET_TO_DEFAULT     = "Setup cluster to default failed"
	ERR_CLUSTER_DEFAULT_NOT_FOUND = "You do not have default cluster"
)

// Application service
const (
	ERR_APPLICATION_NOT_DISABLED     = "Application have not been disabled yet"
	ERR_APPLICATION_ALREADY_DISABLED = "Application already disabled"
	ERR_APPLICATION_ALREADY_ARCHIVED = "Application already archived"
)

// authentication & authorization
const (
	ERR_EXPIRED_TOKEN = "Token is expired"
	ERR_LOAD_CONFIG   = "Unable to load config"
	ERR_ENCRYPT       = "Unable to encrypt token"
	ERR_DECRYPT       = "Unable to decrypt token"
	ERR_INVALID_URL   = "URL not valid"
	ERR_CORS          = "Invalid Origin"
	ERR_GENERATE_ID   = "Unable to generate ID"
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
	KeyInternalAuthAddClaimErr       = "internal_auth_add_claim_err"
	KeyInternalAuthUpdateUserErr     = "internal_auth_update_user_err"
	KeyInternalAuthDeleteUserErr     = "internal_auth_delete_user_err"
	KeyInternalAuthVerifyUserErr     = "internal_auth_verify_user_err"
	KeyInternalAuthCreateUserErr     = "internal_auth_create_user_err"
	KeyInternalAuthUserNotFoundErr   = "internal_auth_user_not_found_err"
	KeyInternalAuthEmailVerifiedErr  = "internal_auth_email_verified_err"
	KeyInternalAuthInvitationLinkErr = "internal_auth_invitation_link_err"
	KeyInternalAuthCreateTokenErr    = "internal_auth_create_token_err"

	KeyInternalDaprInitErr          = "internal_dapr_init_err"
	KeyInternalDaprPublishErr       = "internal_dapr_publish_err"
	KeyInternalDaprGetStateErr      = "internal_dapr_get_state_err"
	KeyInternalDaprSaveStateErr     = "internal_dapr_save_state_err"
	KeyInternalDaprStateNotFoundErr = "internal_dapr_state_not_found_err"
	KeyInternalDaprDeleteStateErr   = "internal_dapr_delete_state_err"
	KeyInternalDaprLockErr          = "internal_dapr_lock_err"
	KeyInternalDaprUnLockErr        = "internal_dapr_unlock_err"

	KeyInternalQuotaCountErr                      = "internal_quota_count_err"
	KeyInternalQuotaGetSubscriptionErr            = "internal_quota_get_subscription_err"
	KeyInternalQuotaGetSubscriptionTypeProductErr = "internal_quota_get_subscription_type_product_err"
	KeyInternalQuotaActionNotAllowedErr           = "internal_quota_action_not_allowed"

	KeyInternalPaymentCreateMethodErr       = "internal_payment_create_method_err"
	KeyInternalPaymentDetachMethodErr       = "internal_payment_detach_method_err"
	KeyInternalPaymentAttachMethodErr       = "internal_payment_attach_method_err"
	KeyInternalPaymentCreateSubscriptionErr = "internal_payment_create_subscription_err"
	KeyInternalPaymentCancelSubscriptionErr = "internal_payment_cancel_subscription_err"
	KeyInternalPaymentCreateCustomerErr     = "internal_payment_create_customer_err"
	KeyInternalPaymentDeleteCustomerErr     = "internal_payment_delete_customer_err"
	KeyInternalPaymentUpdateCustomerErr     = "internal_payment_update_customer_err"
	KeyInternalPaymentCreateUsageRecordErr  = "internal_payment_create_usage_record_err"
	KeyInternalPaymentCreateSetupIntentErr  = "internal_payment_create_setup_intent_err"

	KeyInternalAuthzErr = "internal_authz_err"

	KeyInternalLockHoldByAnother = "internal_lock_hold_by_another"
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

const (
	KeyRemoteDbStateStoreGetErr  = "remote_db_state_store_get_err"
	KeyRemoteDbStateStoreSaveErr = "remote_db_state_store_save_err"
	KeyRemoteDbUnmarshalErr      = "remote_db_unmarshal_err"
	KeyRemoteDbMarshalErr        = "remote_db_marshal_err"
	KeyRemoteDbExternalAPIErr    = "remote_db_external_api_err"
)
