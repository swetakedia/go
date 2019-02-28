package aurora

import (
	"github.com/hcnet/go/services/aurora/internal/reap"
)

func initReaper(app *App) {
	app.reaper = reap.New(app.config.HistoryRetentionCount, app.AuroraSession(nil))
}

func init() {
	appInit.Add("reaper", initReaper, "app-context", "log", "aurora-db")
}
