package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

var response *widget.Label

func main() {
	fontpaths := findfont.List()
	for _, path := range fontpaths {
		if strings.Contains(path, "STSONG.TTF") {
			os.Setenv("FYNE_FONT", path)
		}
	}
	a := app.New()
	win := a.NewWindow("HttpAPITest")
	response = widget.NewLabel("Response")
	win.Resize(fyne.NewSize(1000, 50))
	win.SetContent(APIUI())
	win.ShowAndRun()
}

func APIUI() *fyne.Container {
	editURLBox := widget.NewEntry()
	editURLBox.SetPlaceHolder("Please Enter URL")
	APISelect := widget.NewSelect([]string{"GET", "POST"}, ChooseAPISelect)
	APISelect.SetSelected("GET")
	SendButton := widget.NewButton("Send", func() { SendButton(editURLBox.Text) })
	SendButton.Importance = widget.HighImportance
	HBox := container.NewGridWithColumns(3, APISelect, editURLBox, SendButton, response)
	return HBox
}

func ChooseAPISelect(API string) {
	fmt.Println(API)
}

func SendButton(URL string) {
	body := httpGet(URL)
	if body == nil {
		return
	}
	response.SetText(string(body))
	fmt.Println(string(body))
}

func httpGet(url string) []byte {
	rsps, err := http.Get(url)
	if err != nil {
		response.SetText(err.Error())
		fmt.Println(err.Error())
		return nil
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		response.SetText(err.Error())
		fmt.Println("read error", err)
		return nil
	}
	return body
}
