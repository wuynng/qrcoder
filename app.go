package qrcoder

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/skip2/go-qrcode"
	"path"

	"os"
)

func init() {
	// fyne still not support unicode
	// see https://github.com/fyne-io/fyne/issues/73
	os.Setenv("FYNE_FONT", "/System/Library/Fonts/STHeiti Light.ttc")
}

var w fyne.Window
var input *widget.Entry
var cancelBtn, confirmBtn, saveBtn *widget.Button
var img *canvas.Image

var text = ""

func Start() {
	myApp := app.New()
	w = myApp.NewWindow("QR Coder")
	w.SetMaster()

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		leftContent(), rightContent()))

	w.Resize(fyne.NewSize(600, 300))
	w.ShowAndRun()

	tidy()
}

func rightContent() *fyne.Container {
	img = canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal

	saveBtn = widget.NewButtonWithIcon("Save To Disk", theme.DocumentSaveIcon(), func() {
		save()
	})
	saveBtn.Hide()

	return fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, saveBtn, nil, nil),
		img, saveBtn)
}


func leftContent() *fyne.Container {
	input = widget.NewMultiLineEntry()
	input.SetPlaceHolder("Paste your text...")

	cancelBtn = widget.NewButtonWithIcon("Reset", theme.ViewRefreshIcon(), func() {
		reset()
	})
	confirmBtn = widget.NewButtonWithIcon("Generate", theme.ConfirmIcon(), func() {
		confirm()
	})

	toolbar := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		cancelBtn, confirmBtn)

	form := widget.NewScrollContainer(input)

	left := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, toolbar, nil, nil),
		form, toolbar,
	)
	return left
}

// Save save QR code to disk.
func save() {
	filepath := path.Join(os.ExpandEnv("$HOME"), "qrcode.png")
	err := qrcode.WriteFile(input.Text, qrcode.Medium, 256,  filepath)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	dialog.ShowInformation("Success", "Saved the '" + filepath + "' to your desktop.", w)
}

func tidy() {
	os.Unsetenv("FYNE_FONT")
}

func reset() {
	if input.Text == "" {
		return
	}
	input.SetText("")
	input.Refresh()

	img.Resource = theme.FyneLogo()
	img.Refresh()

	saveBtn.Hide()
}

func confirm() {
	//time.Sleep(3*time.Second)
	val := input.Text
	if val == "" || val == text {
		return
	}
	text = val

	buf, _ := qrcode.Encode(text, qrcode.Medium, 256)
	//img, _ := png.Decode(bytes.NewReader(buf))
	//right.Image = img
	img.Resource = fyne.NewStaticResource("qrcode", buf)
	//right.Refresh()
	img.Refresh()

	saveBtn.Show()
}
