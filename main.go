package main

import (
	"io"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getWikiData(url string) string {
	data, err := http.Get("https://wiki.kerbalspaceprogram.com/index.php?title=Kerbin&action=raw")
	check(err)
	defer data.Body.Close()

	byteData, readErr := io.ReadAll(data.Body)
	check(readErr)

	return string(byteData)
}

//func parseWikiData(data string)

func mainWikiPage(url string) *container.TabItem {
	return container.NewTabItemWithIcon("", theme.DocumentIcon(), widget.NewLabel("test"))
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello, world")
	

	w.ShowAndRun()
}
