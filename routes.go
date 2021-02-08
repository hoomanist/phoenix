package main

func (app *App) SetupRoutes() {
	app.router.HandleFunc("/ping", app.HandlePing()).Methods("GET")
	app.router.HandleFunc("/register", app.HandleRegister()).Methods("POST")
	app.router.HandleFunc("/login", app.HandleLogin()).Methods("POST")
	app.router.HandleFunc("/submit/music", app.HandleMusicUpload()).Methods("POST")

}
