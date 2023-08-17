// Package consts
package consts

const (

	// ResponseAuthenticationFailure response var
	ResponseAuthenticationFailure = `AUTHENTICATION_FAILURE`

	// ResponseSignatureFailure response var
	ResponseSignatureFailure = `SIGNATURE_FAILURE`

	// MiddlewarePassed response var for middleware http request passed
	MiddlewarePassed = `MIDDLEWARE_PASSED`

	// VALIDATION_INVALID_CUSTOMER_ID response var
	ResponseValidationInvalidCustomerId = `VALIDATION_INVALID_CUSTOMER_ID`

	// ResponseValidationFailure response var for general validation error or not pass
	ResponseValidationFailure = `VALIDATION_FAILURE`

	// ResponseValidationInvalidEmail response var
	ResponseValidationInvalidEmail = `VALIDATION_INVALID_EMAIL`

	// ResponseDataNotFound response var for general not found data in our system or third party services
	ResponseDataNotFound = `DATA_NOT_FOUND`

	// CUSTOMER_NOT_FOUND response var
	ResponseCustomerNotFound = `CUSTOMER_NOT_FOUND`

	// ResponseProductNotFound response var
	ResponseProductNotFound = `PRODUCT_NOT_FOUND`

	// ResponseAlreadyPaid response var for the billing already paid status
	ResponseAlreadyPaid = `ALREADY_PAID`

	// ResponseTransactionProcessing response var
	ResponseTransactionProcessing = `TRANSACTION_PROCESSING`

	// ResponseTransactionNotFound response var for transaction is not found
	ResponseTransactionNotFound = `TRANSACTION_NOT_FOUND`

	// ResponseTransactionRejected response var for transaction is rejected
	ResponseTransactionRejected = `TRANSACTION_REJECTED`

	// ResponseTransactionInvalid response name var
	ResponseTransactionInvalid = `TRANSACTION_INVALID`

	// ResponseTransactionDuplicate response var for duplicate transaction
	ResponseTransactionDuplicate = `TRANSACTION_DUPLICATE`

	// ResponseTransactionInsufficientBalance response var for insufficient balance
	ResponseTransactionInsufficientBalance = `TRANSACTION_INSUFFICIENT_BALANCE`

	// ResponseTransactionNoBills response var for when if haven't billing
	ResponseTransactionNoBills = `TRANSACTION_NO_BILLS`

	// ResponseTransactionFailed response var when  transaction failed
	ResponseTransactionFailed = `TRANSACTION_FAILED`

	// ResponseInquiryNotAvailable response var for inquiry is not available
	ResponseInquiryNotAvailable = `INQUIRY_NOT_AVAILABLE`

	// ResponseInquiryFailure response var
	ResponseInquiryFailure = `INQUIRY_FAILURE`

	// ResponseSuccess response var for general success
	ResponseSuccess = "SUCCESS"

	// ResponseInternalFailure response var for internal server error or like something went wrong in system
	ResponseInternalFailure = `INTERNAL_FAILURE`

	// ResponseVendorIdInvalid response var for when invalid provider vendor id
	ResponseVendorIdInvalid = `VENDOR_ID_INVALID`

	// ResponseVendorError response var for when vendor error
	ResponseVendorError = `VENDOR_ERROR`

	// ResponseUnprocessableEntity response var for general wen we cannot continue process and can not retry
	ResponseUnprocessableEntity = `UNPROCESSABLE_ENTITY`

	// ResponseServiceUnavailable response var when tah service is not available
	ResponseServiceUnavailable = `SERVICE_UNAVAILABLE`

	// ResponseUnknownError response var for unknown error
	ResponseUnknownError = `UNKNOWN_ERROR`

	// ResponseForbidden response var for forbidden access
	ResponseForbidden = `FORBIDDEN`

	// ResponseOrderNotifyFailed response var
	ResponseOrderNotifyFailed = `ORDER_NOTIFY_FAILED`

	// ResponseRequestTimeout response var
	ResponseRequestTimeout = `REQUEST_TIMEOUT`
)
