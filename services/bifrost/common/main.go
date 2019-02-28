package common

import (
	"github.com/hcnet/go/support/log"
)

const HcnetAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
