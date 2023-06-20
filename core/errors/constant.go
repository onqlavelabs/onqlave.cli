package errors

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
	ScopeCLI        = "cli"
)

const (
	KeyCLINotLoggedIn          = "cli_not_logged_in"
	KeyCLIMissingRequiredField = "cli_missing_required_field"
	KeyCLIInvalidValue         = "cli_invalid_value"
	KEyCLIEmptyData            = "cli_empty_data"
	KeyCLIRenderDataFailed     = "cli_render_data_failed"
	KeyCLIEnvironmentNotConfig = "cli_environment_not_config"
)

const (
	KeyRemoteDbStateStoreGetErr  = "remote_db_state_store_get_err"
	KeyRemoteDbStateStoreSaveErr = "remote_db_state_store_save_err"
	KeyRemoteDbUnmarshalErr      = "remote_db_unmarshal_err"
	KeyRemoteDbMarshalErr        = "remote_db_marshal_err"
	KeyRemoteDbExternalAPIErr    = "remote_db_external_api_err"
)
