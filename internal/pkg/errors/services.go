package errors

import "fmt"

func NewErrUserNotFound(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_USER_NOT_FOUND,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrApplicationNotDisabled(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_APPLICATION_NOT_DISABLED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrApplicationAlreadyDisable(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_APPLICATION_ALREADY_DISABLED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrApplicationAlreadyArchived(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_APPLICATION_ALREADY_ARCHIVED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrPaymentUpdateCustomer(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_PAYMENT_UPDATE_CUSTOMER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrPaymentAttachPaymentMethod(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_PAYMENT_ATTACH_PAYMENT_METHOD,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrPaymentDetachPaymentMethod(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_PAYMENT_DETACH_PAYMENT_METHOD,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrMaximumPaymentMethod(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_MAXIMUM_PAYMENT_METHOD,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseCreateUser(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_CREATE_USER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseDeleteUser(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_DELETE_USER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseInvitationLink(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_INVITATION_LINK,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseResetPasswordLink(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_RESET_PASSWORD_LINK,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewInactiveUser(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_USER_INACTIVE,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

// Billing Service
func NewDuplicateRecord(entityName, ID string) error {
	return NewErrDuplicateID(fmt.Errorf("there is already existing in %s with ID %s", entityName, ID))
}

func NewErrCreateAuthToken(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_CREATE_AUTH_TOKEN,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrActionNotAllow(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_ACTION_NOT_ALLOW,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrUserDisable(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_USER_DISABLE,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrLoginTimeOut(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_LOGIN_TIMEOUT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrSendMailTimeOut(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_SEND_MAIL_TIMEOUT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrRegister(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_REGISTER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrUserEnable(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_USER_ENABLE,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrMailOperation(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_OPERATION_NOT_SUPPORTED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFirebaseInit(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_INIT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrTenantDisable(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_TENANT_DISABLE,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseVerifyEmail(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_VERIFY_EMAIL_LINK,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFireBaseUserMailAlreadyVerified(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_EMAIL_ALREADY_VERIFIED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFirebaseUserNotFound(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_USER_NOT_FOUND,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFirebaseVerifyUser(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_VERIFY_USER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFirebaseUpdateUser(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_UPDATE_USER,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrFirebaseAddClaim(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_FIREBASE_ADD_CLAIM,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrRegisterCompleted(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_REGISTER_COMPLETED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrRegisterNotActioned(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_REGISTER_NOT_ACTIONED,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrRegisterRevokeEmail(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_REVOKE_EMAIL,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrInitService(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_INIT_SERVICE,
		baseErr: errBase,
		mode:    1,
	}
}

func NewErrPublishEvent(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_PUBLISH_EVENT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrPublishCommand(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_PUBLISH_COMMAND,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrCreateCommand(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_CREATE_COMMAND,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrInvalidPayload(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_INVALID_PAYLOAD,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrExtractData(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_EXTRACT_DATA,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrSendMail(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERROR_SEND_MAIL,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrSetArxDefault(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_UNABLE_SET_TO_DEFAULT,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}

func NewErrDontHaveDefaultCluster(errBase errorI) error {
	return Error{
		service: errConfig.service,
		message: ERR_CLUSTER_DEFAULT_NOT_FOUND,
		baseErr: errBase,
		mode:    errConfig.mode,
	}
}
