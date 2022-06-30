package log

import (
	"git.medlinker.com/service/common/cmodel"
)

var Logger cmodel.Logger = nil

func SetLogger(log cmodel.Logger) {
	Logger = log
}
