package app

func (a *App) setupDb() {
	a.db.AutoMigrate(&Link{})
}
