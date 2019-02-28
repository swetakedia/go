package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/hcnet/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("hcnet_address", govalidator.CustomTypeValidator(isHcnetAddress))
}

func isHcnetAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
