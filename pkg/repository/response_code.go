package repository

import Error "gitlab.pede.id/otto-library/golang/share-pkg/error"

// const for code
const (
	SuccessCode      = 200
	ContinueCode     = 100
	UndefinedCode    = 500
	BadRequestCode   = 400
	NotFoundCode     = 404
	UnauthorizedCode = 401
	PendingCode      = 451
)

// const for Status
const (
	SuccessStatus   = "SUCCESS"
	PendingStatus   = "PENDING"
	FailedStatus    = "FAILED"
	UndefinedStatus = "FAILED"
	ContinueStatus  = "CONTINUE"
)

var (
	PendingErr      = Error.NewError(PendingCode, PendingStatus, "Transaksi Sedang Diproses. Jika transaksi gagal dana Anda akan dikembalikan ke saldo OttoCash")
	UndefinedErr    = Error.NewError(UndefinedCode, UndefinedStatus, "Terjadi Kesalahan Pada Server")
	ContinueErr     = Error.NewError(ContinueCode, ContinueStatus, "Silahkan Lanjutkan ke Tahap Berikutnya")
	UnauthorizedErr = Error.NewError(UnauthorizedCode, FailedStatus, "Sesi Anda Telah Habis")
	NotFoundErr     = Error.NewError(NotFoundCode, FailedStatus, "Data Tidak Ditemukan")
	BadRequestErr   = Error.NewError(BadRequestCode, FailedStatus, "Format Request Salah")
)

func SetError(code int, message ...string) error {
	if code == SuccessCode {
		return nil
	}
	defaultMessage := UndefinedErr.Error()
	if code >= len(listError) {
		m := defaultMessage
		if len(message) > 0 && message[0] != "" {
			m = message[0]
		}
		return Error.NewError(UndefinedCode, UndefinedStatus, m)
	}
	errFromList := listError[code]
	if errFromList != nil {
		if len(message) > 0 {
			m := message[0]
			if he, ok := errFromList.(*Error.ApplicationError); ok {
				if m == "" {
					m = he.Message
				}
				return Error.NewError(he.ErrorCode, he.Status, m)
			} else {
				return errFromList
			}
		} else {
			return errFromList
		}
	} else {
		m := defaultMessage
		if len(message) > 0 && message[0] != "" {
			m = message[0]
		}
		return Error.NewError(UndefinedCode, UndefinedStatus, m)
	}
}

var listError = []error{
	PendingCode:      PendingErr,
	ContinueCode:     ContinueErr,
	UndefinedCode:    UndefinedErr,
	UnauthorizedCode: UnauthorizedErr,
	NotFoundCode:     NotFoundErr,
	BadRequestCode:   BadRequestErr,
}
