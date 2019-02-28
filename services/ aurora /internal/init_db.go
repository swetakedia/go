package aurora

import (
	"github.com/hcnet/go/services/aurora/internal/db2/core"
	"github.com/hcnet/go/services/aurora/internal/db2/history"
	"github.com/hcnet/go/support/db"
	"github.com/hcnet/go/support/log"
)

func initAuroraDb(app *App) {
	session, err := db.Open("postgres", app.config.DatabaseURL)
	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down Aurora.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)
	app.historyQ = &history.Q{session}
}

func initCoreDb(app *App) {
	session, err := db.Open("postgres", app.config.HcnetCoreDatabaseURL)
	if err != nil {
		log.Panic(err)
	}

	// Make sure MaxIdleConns is equal MaxOpenConns. In case of high variance
	// in number of requests closing and opening connections may slow down Aurora.
	session.DB.SetMaxIdleConns(app.config.MaxDBConnections)
	session.DB.SetMaxOpenConns(app.config.MaxDBConnections)
	app.coreQ = &core.Q{session}
}

func init() {
	appInit.Add("aurora-db", initAuroraDb, "app-context", "log")
	appInit.Add("core-db", initCoreDb, "app-context", "log")
}
