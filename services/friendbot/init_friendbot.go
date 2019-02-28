package main

import (
	"net/http"

	"github.com/hcnet/go/clients/aurora"
	"github.com/hcnet/go/services/friendbot/internal"
	"github.com/hcnet/go/strkey"
)

func initFriendbot(
	friendbotSecret string,
	networkPassphrase string,
	auroraURL string,
	startingBalance string,
) *internal.Bot {

	if friendbotSecret == "" || networkPassphrase == "" || auroraURL == "" || startingBalance == "" {
		return nil
	}

	// ensure its a seed if its not blank
	strkey.MustDecode(strkey.VersionByteSeed, friendbotSecret)

	return &internal.Bot{
		Secret: friendbotSecret,
		Aurora: &aurora.Client{
			URL:  auroraURL,
			HTTP: http.DefaultClient,
		},
		Network:           networkPassphrase,
		StartingBalance:   startingBalance,
		SubmitTransaction: internal.AsyncSubmitTransaction,
	}
}
