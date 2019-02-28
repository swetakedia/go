package aurora

func initHcnetCoreInfo(app *App) {
	app.UpdateHcnetCoreInfo()
}

func init() {
	appInit.Add("hcnetCoreInfo", initHcnetCoreInfo, "app-context", "log")
}
