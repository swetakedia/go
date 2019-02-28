package bridge

import (
	"github.com/hcnet/go/amount"
	b "github.com/hcnet/go/build"
	shared "github.com/hcnet/go/services/internal/bridge-compliance-shared"
	"github.com/hcnet/go/services/internal/bridge-compliance-shared/http/helpers"
	"github.com/hcnet/go/services/internal/bridge-compliance-shared/protocols"
)

// ChangeTrustOperationBody represents change_trust operation
type ChangeTrustOperationBody struct {
	Source *string
	Asset  protocols.Asset
	// nil means max limit
	Limit *string
}

// ToTransactionMutator returns go-hcnet-base TransactionMutator
func (op ChangeTrustOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		op.Asset.ToBaseAsset(),
	}

	if op.Limit == nil {
		// Set MaxLimit
		mutators = append(mutators, b.MaxLimit)
	} else {
		mutators = append(mutators, b.Limit(*op.Limit))
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.ChangeTrust(mutators...)
}

// Validate validates if operation body is valid.
func (op ChangeTrustOperationBody) Validate() error {
	err := op.Asset.Validate()
	if err != nil {
		return helpers.NewInvalidParameterError("asset", err.Error())
	}

	if op.Limit != nil {
		_, err := amount.Parse(*op.Limit)
		if err != nil {
			return helpers.NewInvalidParameterError("limit", "Limit is not a valid amount.")
		}
	}

	if op.Source != nil && !shared.IsValidAccountID(*op.Source) {
		return helpers.NewInvalidParameterError("source", "Source must be a public key (starting with `G`).")
	}

	return nil
}
