package response

// Constant
const (
	CommonSuccess           = "common_success"
	CommonBadRequest        = "common_bad_request"
	CommonUnauthorized      = "common_unauthorized"
	CommonNotFound          = "common_not_found"
	CommonInvalidChecksum   = "common_invalid_checksum"
	CommonInvalidPagination = "common_invalid_pagination"
	CommonForbidden         = "common_forbidden"
	CommonInvalidID         = "common_invalid_id"
	CommonTransactionFailed = "common_transaction_failed"
	CommonPageInvalid       = "common_page_invalid"
	CommonLimitInvalid      = "common_limit_invalid"
)

const (
	commonSuccessCode = iota + 1
	commonBadRequestCode
	commonUnauthorizedCode
	commonNotFoundCode
	commonInvalidChecksumCode
	commonInvalidPaginationCode
	commonForbiddenCode
	commonInvalidIDCode
	commonTransactionFailedCode
	commonPageInvalid
	commonLimitInvalid
)

// 1-99
var common = []Code{
	{
		Key:     CommonSuccess,
		Code:    commonSuccessCode,
		Message: "thành công",
	},
	{
		Key:     CommonBadRequest,
		Code:    commonBadRequestCode,
		Message: "dữ liệu không hợp lệ",
	},
	{
		Key:     CommonUnauthorized,
		Code:    commonUnauthorizedCode,
		Message: "bạn không có quyền thực hiện hành động này",
	},
	{
		Key:     CommonNotFound,
		Code:    commonNotFoundCode,
		Message: "dữ liệu không tìm thấy",
	},
	{
		Key:     CommonInvalidChecksum,
		Code:    commonInvalidChecksumCode,
		Message: "xác thực dữ liệu thất bại",
	},
	{
		Key:     CommonInvalidPagination,
		Code:    commonInvalidPaginationCode,
		Message: "phân trang không hợp lệ",
	},
	{
		Key:     CommonForbidden,
		Code:    commonForbiddenCode,
		Message: "không có quyền truy cập",
	},
	{
		Key:     CommonInvalidID,
		Code:    commonInvalidIDCode,
		Message: "id không hợp lệ",
	},
	{
		Key:     CommonTransactionFailed,
		Code:    commonTransactionFailedCode,
		Message: "đã có lỗi xảy ra trong quá trình xử lý, vui lòng thử lại sau giây lát",
	},
	{
		Key:     CommonPageInvalid,
		Code:    commonPageInvalid,
		Message: "page không đúng định dạng",
	},
	{
		Key:     CommonLimitInvalid,
		Code:    commonLimitInvalid,
		Message: "limit không đúng định dạng",
	},
}
