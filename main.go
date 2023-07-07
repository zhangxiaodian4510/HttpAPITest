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

var HeaderList []Header

type Header struct {
	HBox  *fyne.Container
	Key   *widget.Entry
	Value *widget.Entry
}

func main() {
	fontpaths := findfont.List()
	for _, path := range fontpaths {
		if strings.Contains(path, "STSONG.TTF") {
			os.Setenv("FYNE_FONT", path)
		}
	}
	a := app.New()
	win := a.NewWindow("HttpAPITest")
	win.Resize(fyne.NewSize(1000, 500))
	win.SetContent(APIUI())
	win.ShowAndRun()
}

func APIUI() *fyne.Container {

	response = widget.NewLabel("Response")
	response.Resize(fyne.NewSize(900, 200))
	response.Move(fyne.NewPos(20, 300))

	editURLBox := widget.NewEntry()
	editURLBox.SetPlaceHolder("Please Enter URL")
	editURLBox.Resize(fyne.NewSize(650, 35))
	editURLBox.Move(fyne.NewPos(150, 20))

	APISelect := widget.NewSelect([]string{"GET", "POST"}, ChooseAPISelect)
	APISelect.SetSelected("GET")
	APISelect.Resize(fyne.NewSize(85, 35))
	APISelect.Move(fyne.NewPos(50, 20))

	SendButton := widget.NewButton("Send", func() { SendButton(editURLBox.Text) })
	SendButton.Importance = widget.HighImportance
	SendButton.Resize(fyne.NewSize(85, 35))
	SendButton.Move(fyne.NewPos(825, 20))

	HeaderList := NewHeaderList()
	HeaderContainer := container.NewGridWithColumns(2, HeaderList[0].Key, HeaderList[0].Value)
	HeaderContainer.Resize(fyne.NewSize(865, 35))
	HeaderContainer.Move(fyne.NewPos(50, 75))

	MainLayout := container.NewWithoutLayout(APISelect, editURLBox, SendButton, response, HeaderContainer)
	return MainLayout
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
	for _, Header := range HeaderList {
		rsps.Header.Add(Header.Key.Text, Header.Value.Text)
	}
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

func NewHeader() *Header {
	h := &Header{}
	h.Key = widget.NewEntry()
	h.Key.PlaceHolder = "Key"
	h.Value = widget.NewEntry()
	h.Value.PlaceHolder = "Value"
	h.HBox = container.NewGridWithColumns(2, h.Key, h.Value)
	return h
}

func NewHeaderList() []Header {
	h := NewHeader()
	HeaderList = append(HeaderList, *h)
	return HeaderList
}

func AddHeader() {
	h := NewHeader()
	HeaderList = append(HeaderList, *h)
}
