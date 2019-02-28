package resourceadapter

import (
	. "github.com/hcnet/go/protocols/aurora"
	"github.com/hcnet/go/services/aurora/internal/db2/core"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
