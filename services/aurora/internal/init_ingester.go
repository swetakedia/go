package aurora

import (
	"log"

	"github.com/hcnet/go/services/aurora/internal/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.config.NetworkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with hcnet-core.")
	}

	app.ingester = ingest.New(
		app.config.NetworkPassphrase,
		app.config.HcnetCoreURL,
		app.CoreSession(nil),
		app.AuroraSession(nil),
		ingest.Config{
			EnableAssetStats:         app.config.EnableAssetStats,
			IngestFailedTransactions: app.config.IngestFailedTransactions,
		},
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
	app.ingester.HistoryRetentionCount = app.config.HistoryRetentionCount
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "aurora-db", "core-db", "hcnetCoreInfo")
}
