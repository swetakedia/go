package handlers

import (
	"github.com/hcnet/go/clients/federation"
	"github.com/hcnet/go/clients/aurora"
	"github.com/hcnet/go/clients/hcnettoml"
	"github.com/hcnet/go/services/bridge/internal/config"
	"github.com/hcnet/go/services/bridge/internal/db"
	"github.com/hcnet/go/services/bridge/internal/listener"
	"github.com/hcnet/go/services/bridge/internal/submitter"
	"github.com/hcnet/go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Aurora              aurora.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	HcnetTomlResolver  hcnettoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
