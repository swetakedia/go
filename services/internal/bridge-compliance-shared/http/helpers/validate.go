package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/hcnet/go/address"
	"github.com/hcnet/go/amount"
	"github.com/hcnet/go/strkey"
)

func init() {
	govalidator.CustomTypeTagMap.Set("hcnet_accountid", govalidator.CustomTypeValidator(isHcnetAccountID))
	govalidator.CustomTypeTagMap.Set("hcnet_seed", govalidator.CustomTypeValidator(isHcnetSeed))
	govalidator.CustomTypeTagMap.Set("hcnet_asset_code", govalidator.CustomTypeValidator(isHcnetAssetCode))
	govalidator.CustomTypeTagMap.Set("hcnet_address", govalidator.CustomTypeValidator(isHcnetAddress))
	govalidator.CustomTypeTagMap.Set("hcnet_amount", govalidator.CustomTypeValidator(isHcnetAmount))
	govalidator.CustomTypeTagMap.Set("hcnet_destination", govalidator.CustomTypeValidator(isHcnetDestination))

}

func Validate(request Request, params ...interface{}) error {
	valid, err := govalidator.ValidateStruct(request)

	if !valid {
		fields := govalidator.ErrorsByField(err)
		for field, errorValue := range fields {
			switch {
			case errorValue == "non zero value required":
				return NewMissingParameter(field)
			case strings.HasSuffix(errorValue, "does not validate as hcnet_accountid"):
				return NewInvalidParameterError(field, "Account ID must start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as hcnet_seed"):
				return NewInvalidParameterError(field, "Account secret must start with `S` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as hcnet_asset_code"):
				return NewInvalidParameterError(field, "Asset code must be 1-12 alphanumeric characters.")
			case strings.HasSuffix(errorValue, "does not validate as hcnet_address"):
				return NewInvalidParameterError(field, "Hcnet address must be of form user*domain.com")
			case strings.HasSuffix(errorValue, "does not validate as hcnet_destination"):
				return NewInvalidParameterError(field, "Hcnet destination must be of form user*domain.com or start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as hcnet_amount"):
				return NewInvalidParameterError(field, "Amount must be positive and have up to 7 decimal places.")
			default:
				return NewInvalidParameterError(field, errorValue)
			}
		}
	}

	return request.Validate(params...)
}

// These are copied from support/config. Should we move them to /strkey maybe?
func isHcnetAccountID(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, enc)

	if err == nil {
		return true
	}

	return false
}

func isHcnetSeed(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteSeed, enc)

	if err == nil {
		return true
	}

	return false
}

func isHcnetAssetCode(i interface{}, context interface{}) bool {
	code, ok := i.(string)

	if !ok {
		return false
	}

	if !govalidator.IsByteLength(code, 1, 12) {
		return false
	}

	if !govalidator.IsAlphanumeric(code) {
		return false
	}

	return true
}

func isHcnetAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)
	if err != nil {
		return false
	}

	return true
}

func isHcnetAmount(i interface{}, context interface{}) bool {
	am, ok := i.(string)

	if !ok {
		return false
	}

	_, err := amount.Parse(am)
	if err != nil {
		return false
	}

	return true
}

// isHcnetDestination checks if `i` is either account public key or Hcnet address.
func isHcnetDestination(i interface{}, context interface{}) bool {
	dest, ok := i.(string)

	if !ok {
		return false
	}

	_, err1 := strkey.Decode(strkey.VersionByteAccountID, dest)
	_, _, err2 := address.Split(dest)

	if err1 != nil && err2 != nil {
		return false
	}

	return true
}
