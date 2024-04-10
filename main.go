package main

import (
	"embed"

	"github.com/Daanooo/taski/internal/controller"
	"github.com/Daanooo/taski/internal/data"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	db, dberr := data.GetSqlite("data.db")
	if dberr != nil {
		panic(dberr)
	}
	defer db.Close()

	//Register repositories
	tasks := data.NewTaskRepository(db)

	// Register controllers
	tasksController := controller.NewTaskController(tasks)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "taski",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			tasksController,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
