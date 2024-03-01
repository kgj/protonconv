package main

import (
	"fmt"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"os"
	"protonconv/utils"
)

func main() {

	ui := cli

	if len(os.Args[1:]) == 0 {
		ui = gui
	}

	ui()
}

func cli() {
	err := utils.Convert(utils.Params{}.Parse())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gui() {
	app := gtk.NewApplication("ca.joneslabs.protonconv", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	window := gtk.NewWindow()
	window.SetApplication(app)
	window.SetTitle("Proton Pass Export Converter")
	window.SetResizable(false)
	window = buildGUI(window)
	window.Present()
}

func buildGUI(window *gtk.Window) *gtk.Window {
	grid := gtk.NewGrid()
	grid.SetHAlign(gtk.AlignCenter)
	grid.SetVAlign(gtk.AlignCenter)
	grid.SetRowSpacing(15)
	grid.SetColumnSpacing(15)
	grid.SetMarginStart(25)
	grid.SetMarginEnd(25)
	grid.SetMarginTop(25)
	grid.SetMarginBottom(25)

	window.SetChild(grid)

	passFileLbl := gtk.NewLabel("Proton Pass Export")
	passFileLbl.SetHAlign(gtk.AlignStart)
	passFileChooserBtn := gtk.NewButtonFromIconName("folder-new")

	passFileEntry := gtk.NewEntry()
	passFileEntry.SetEditable(false)
	passFileEntry.SetSensitive(false)
	passFileEntry.SetText("Select a ProtonPass export JSON file")

	params := utils.Params{}

	passFileChooserBtn.ConnectClicked(func() {
		chooser := gtk.NewFileChooserNative("Choose File", window, 0, "", "")
		chooser.SetSelectMultiple(false)
		filter := gtk.NewFileFilter()
		filter.SetName("ProtonPass JSON Export")
		filter.AddPattern("*.json")
		chooser.AddFilter(filter)
		chooser.SetModal(true)

		chooser.ConnectResponse(func(responseId int) {
			response := gtk.ResponseType(responseId)

			if response != gtk.ResponseAccept {
				return
			}

			selectedFile := chooser.File()
			filePath := selectedFile.Path()
			passFileEntry.SetText(filePath)
			params.JsonFilename = filePath
		})

		chooser.Show()
	})

	keepassFileLbl := gtk.NewLabel("New Keepass filename")
	keepassFileLbl.SetHAlign(gtk.AlignStart)
	keepassFileBtn := gtk.NewButtonFromIconName("folder-new")

	keepassFileEntry := gtk.NewEntry()
	keepassFileEntry.SetEditable(false)
	keepassFileEntry.SetSensitive(false)
	keepassFileEntry.SetText("KeePass DB Save Location")

	keepassFileBtn.ConnectClicked(func() {
		chooser := gtk.NewFileChooserNative("Choose File", window, 1, "", "")
		chooser.SetSelectMultiple(false)
		chooser.SetModal(true)

		chooser.ConnectResponse(func(responseId int) {
			response := gtk.ResponseType(responseId)

			if response != gtk.ResponseAccept {
				return
			}

			selectedFile := chooser.File()
			filePath := selectedFile.Path()
			keepassFileEntry.SetText(filePath)
			params.DBFileName = selectedFile.Path()
		})

		chooser.Show()
	})

	passwordFileLbl := gtk.NewLabel("New Keepass master password")
	passwordFileLbl.SetHAlign(gtk.AlignStart)
	passwordEntry := gtk.NewPasswordEntry()
	passwordEntry.SetShowPeekIcon(true)

	convertBtn := gtk.NewButtonWithLabel("Convert Proton Pass Export to Keepass File")

	convertBtn.ConnectClicked(func() {

		params.DBFileName = keepassFileEntry.Text()
		params.Password = passwordEntry.Text()

		err := utils.Convert(params)
		if err != nil {
			showGuiErr(err, window)
		} else {
			showGuiSuccess(window)
		}
	})

	grid.Attach(passFileLbl, 0, 2, 1, 1)
	grid.Attach(keepassFileLbl, 0, 3, 1, 1)
	grid.Attach(passFileEntry, 1, 2, 1, 1)
	grid.Attach(passwordFileLbl, 0, 4, 1, 1)

	grid.Attach(passFileChooserBtn, 2, 2, 1, 1)
	grid.Attach(keepassFileBtn, 2, 3, 1, 1)
	grid.Attach(keepassFileEntry, 1, 3, 1, 1)
	grid.Attach(passwordEntry, 1, 4, 1, 1)

	grid.Attach(convertBtn, 1, 5, 1, 1)

	return window
}

func showGuiSuccess(window *gtk.Window) {
	successDialog := gtk.NewMessageDialog(window, gtk.DialogModal, gtk.MessageInfo, gtk.ButtonsOK)
	successDialog.SetMarginStart(15)
	successDialog.SetMarginTop(15)
	successDialog.SetMarginBottom(15)
	successDialog.SetMarginEnd(15)

	successDialog.SetTitle("Success!")
	successDialog.SetName("Success")
	successDialog.SetModal(true)

	successDialog.ConnectResponse(func(responseId int) {
		successDialog.Close()
	})

	successDialog.Show()
}

func showGuiErr(err error, window *gtk.Window) {
	errDialog := gtk.NewMessageDialog(window, gtk.DialogModal, gtk.MessageError, gtk.ButtonsOK)
	errDialog.SetTitle(err.Error())
	errDialog.SetName("Error")
	errDialog.SetModal(true)
	errDialog.SetMarginStart(15)
	errDialog.SetMarginTop(15)
	errDialog.SetMarginBottom(15)
	errDialog.SetMarginEnd(15)

	errDialog.ConnectResponse(func(responseId int) {
		errDialog.Close()
	})

	errDialog.Show()
	fmt.Println(err)
}
