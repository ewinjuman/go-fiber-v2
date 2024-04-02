package repository

import (
	"github.com/go-sql-driver/mysql"
	Error "gitlab.pede.id/otto-library/golang/share-pkg/error"
	"gorm.io/gorm"
)

// const for code
const (
	SuccessCode        = 200
	ContinueCode       = 100
	UndefinedCode      = 500
	BadRequestCode     = 400
	NotFoundCode       = 404
	UnauthorizedCode   = 401
	PendingCode        = 451
	DuplicateEntryCode = 462
)

// const for Status
const (
	SuccessStatus   = "SUCCESS"
	PendingStatus   = "PENDING"
	FailedStatus    = "FAILED"
	UndefinedStatus = "FAILED"
	ContinueStatus  = "CONTINUE"
	ErrorStatus     = "ERROR"
)

var (
	PendingErr        = Error.NewError(PendingCode, PendingStatus, "Transaksi Sedang Diproses. Jika transaksi gagal dana Anda akan dikembalikan ke saldo OttoCash")
	UndefinedErr      = Error.NewError(UndefinedCode, ErrorStatus, "Terjadi Kesalahan Pada Server")
	ContinueErr       = Error.NewError(ContinueCode, ContinueStatus, "Silahkan Lanjutkan ke Tahap Berikutnya")
	UnauthorizedErr   = Error.NewError(UnauthorizedCode, FailedStatus, "Sesi Anda Telah Habis")
	NotFoundErr       = Error.NewError(NotFoundCode, FailedStatus, "Data Tidak Ditemukan")
	BadRequestErr     = Error.NewError(BadRequestCode, FailedStatus, "Format Request Salah")
	DuplicateEntryErr = Error.NewError(DuplicateEntryCode, FailedStatus, "Data Sudah ada")
)
var listError = []error{
	PendingCode:        PendingErr,
	ContinueCode:       ContinueErr,
	UndefinedCode:      UndefinedErr,
	UnauthorizedCode:   UnauthorizedErr,
	NotFoundCode:       NotFoundErr,
	BadRequestCode:     BadRequestErr,
	DuplicateEntryCode: DuplicateEntryErr,
}

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
		return Error.NewError(UndefinedCode, ErrorStatus, m)
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
		return Error.NewError(UndefinedCode, ErrorStatus, m)
	}
}

// Convert and Mapping  MySql Error
func HandleMysqlError(err error) error {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Error.NewError(NotFoundCode, FailedStatus)
		}
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return MapMysqlError(mysqlErr)
		} else {
			return err
		}
	}
	return nil
}

// Mapping  MySql Error
func MapMysqlError(mySqlErr *mysql.MySQLError) (err error) {
	switch mySqlErr.Number {
	case 1062: // MySQL code for duplicate entry
		return DuplicateEntryErr
	default:
		return UndefinedErr
	}
}
