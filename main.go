package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("HttpAPITest")
	win.Resize(fyne.NewSize(1000, 50))
	win.SetContent(APIUI())
	win.ShowAndRun()
}

func APIUI() *fyne.Container {
	editURLBox := widget.NewEntry()
	editURLBox.SetPlaceHolder("Please Enter URL")
	APISelect := widget.NewSelect([]string{"GET", "POST"}, ChooseAPISelect)
	SendButton := widget.NewButton("Send", func() { SendButton(editURLBox.Text) })
	SendButton.Importance = widget.HighImportance
	HBox := container.NewGridWithColumns(3, APISelect, editURLBox, SendButton)
	return HBox
}

func ChooseAPISelect(API string) {
	fmt.Println(API)
}

func SendButton(URL string) {
	body := httpGet(URL)
	fmt.Println(string(body))
}

func httpGet(url string) []byte {
	rsps, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
		return nil
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("read error", err)
		return nil
	}
	return body
}
