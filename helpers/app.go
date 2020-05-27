package helpers

import (
	"fmt"
	"net/http"
	"time"
)

type App struct {
	Name string
	Version float32
	Creator string
	Time string
	Path string
}

func Information(r *http.Request) *App {
	currentTime := time.Now()
	formatted := fmt.Sprintf("%02d:%02d:%02d %02d.%02d.%d", currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Day(), currentTime.Month(), currentTime.Year())
	path := fmt.Sprintf("%v", r.URL.Path)
	return &App{Name: "DNA Sequence Converter", Version: 0.1, Creator: "Matan Budimir", Time: formatted, Path: path}
}