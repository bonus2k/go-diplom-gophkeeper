package mvc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/note_service"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	cu   *ControllerUI
	sn   *note_service.ServiceNote
	once sync.Once
	log  *logger.Logger
	app  = tview.NewApplication()

	pagesMenu     = tview.NewPages()
	notesList     = tview.NewList().ShowSecondaryText(false)
	flexMain      = tview.NewFlex()
	modalViewNote = tview.NewModal()
	modalError    = tview.NewModal()
	textInfo      = tview.NewTextView()
)

type ControllerUI struct {
	infoList []string
}

func NewControllerUI(logger *logger.Logger, serviceNote *note_service.ServiceNote) *ControllerUI {
	once.Do(func() {
		log = logger
		sn = serviceNote
		log.Infof("controller UI initializing")
		log.Infof("create controller UI")
		cu = &ControllerUI{infoList: make([]string, 0)}
		log.Info("create menu")
		createMainMenu()
		log.Info("create flex")
		creteMainFlex()
		log.Info("setup input")
		setInput(cu)
	})
	return cu
}

func (cu *ControllerUI) Run() error {
	log.Infof("controller UI running")
	return app.SetRoot(pagesMenu, true).EnableMouse(true).Run()
}

func (cu *ControllerUI) AddItemInfoList(msg string) {
	cu.infoList = append(cu.infoList, msg)
	textInfo.Clear()
	textInfo.SetText(strings.Join(cu.infoList, "\n"))
	textInfo.SetTextColor(tcell.ColorYellowGreen).SetBorder(true).SetTitle("Info").SetBorderColor(tcell.ColorYellowGreen).SetTitleColor(tcell.ColorYellowGreen)
	textInfo.ScrollToEnd()
}

func (cu *ControllerUI) AddNote(note models.Noteable) {
	storage := sn.AddNote(note)
	createNotesList(storage)
}

func setInput(cu *ControllerUI) *tview.Box {
	return flexMain.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 113:
			app.Stop()
		case 108:
			cu.AddItemInfoList("try load notes")
		case 100:
			cu.AddItemInfoList("try decrypt note")
		case 98:
			formCardBankNote.Clear(true)
			addBankCardNote(cu)
			pagesMenu.SwitchToPage("Add Bank Card Note")
		case 99:
			formCredentialNote.Clear(true)
			addCredentialNote(cu)
			pagesMenu.SwitchToPage("Add Credential")
		case 116:
			formTextNote.Clear(true)
			addTextNote(cu)
			pagesMenu.SwitchToPage("Add Text Note")
		case 105:
			formBinaryNote.Clear(true)
			addBinaryNote(cu)
			pagesMenu.SwitchToPage("Add Binary Note")
		case 114:
			formRegistrationUser.Clear(true)
			addNewUser(cu)
			pagesMenu.SwitchToPage("Registration User")
		case 115:
			formAuthorization.Clear(true)
			authorizationUser(cu)
			pagesMenu.SwitchToPage("Sign in")
		}
		return event
	})
}

func createMainMenu() {
	pagesMenu.AddPage("Menu", flexMain, true, true)
	pagesMenu.AddPage("Add Bank Card Note", createModalForm(formCardBankNote, 70, 23), true, false)
	pagesMenu.AddPage("Add Credential", createModalForm(formCredentialNote, 70, 17), true, false)
	pagesMenu.AddPage("Add Text Note", createModalForm(formTextNote, 70, 19), true, false)
	pagesMenu.AddPage("Add Binary Note", createModalForm(formBinaryNote, 70, 19), true, false)
	pagesMenu.AddPage("Registration User", createModalForm(formRegistrationUser, 70, 13), true, false)
	pagesMenu.AddPage("Show Note", modalViewNote, true, false)
	pagesMenu.AddPage("Error", modalError, true, false)
	pagesMenu.AddPage("Sign in", createModalForm(formAuthorization, 55, 10), true, false)
}

func creteMainFlex() *tview.Flex {
	textMenu1 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(q) quit \n(l) load notes \n(d) decrypt notes")
	textMenu2 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(b) add bank card \n(c) add credential \n(t) add text \n(i) add binary")
	textMenu3 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(r) register an account \n(s) sign in")

	flexMain.SetBorder(true).SetTitle("Welcome to gopher keeper").SetTitleAlign(tview.AlignLeft)
	return flexMain.
		AddItem(notesList, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(false).SetTitle(""), 0, 3, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(textMenu1, 0, 1, false).
				AddItem(textMenu2, 0, 1, false).
				AddItem(textMenu3, 0, 1, false), 4, 1, false), 0, 2, false).
		AddItem(textInfo, 0, 1, false)
}

func createNotesList(storage []models.Noteable) {
	notesList.Clear()

	notesList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		createModalNote(storage[i])
		pagesMenu.SwitchToPage("Show Note")
	})

	for i, note := range storage {
		item := fmt.Sprintf("[\r%s] %s", strings.ToUpper(note.GetType().String()), note.GetName())
		notesList.AddItem(item, "", rune(49+i), nil)
	}
}

func createModalNote(note models.Noteable) {
	modalViewNote.ClearButtons()
	modalViewNote.
		SetText(note.Print()).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				pagesMenu.SwitchToPage("Menu")
			}
		})

}

func createModalForm(p tview.Primitive, width, height int) tview.Primitive {
	flex := tview.NewFlex()
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pagesMenu.SwitchToPage("Menu")
		}
		return event
	})
	return flex.
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}
