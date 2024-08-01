package mvc

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app = tview.NewApplication()

	pagesMenu     = tview.NewPages()
	notesList     = tview.NewList().ShowSecondaryText(false)
	flexMain      = tview.NewFlex()
	modalViewNote = tview.NewModal()
	textInfo      = tview.NewTextView()
	storage       = make([]model.Noteable, 0)
	infoList      = make([]string, 0)
)

func Init() {
	testData()
	createMainMenu()
	creteMainFlex()
	setInput()
	addInfo("application loaded successful")
	if err := app.SetRoot(pagesMenu, true).EnableMouse(true).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addInfo(msg string) {
	infoList = append(infoList, msg)
	textInfo.Clear()
	textInfo.SetText(strings.Join(infoList, "\n"))
	textInfo.SetTextColor(tcell.ColorYellowGreen).SetBorder(true).SetTitle("Info").SetBorderColor(tcell.ColorYellowGreen).SetTitleColor(tcell.ColorYellowGreen)
	textInfo.ScrollToEnd()
}

func addNotesList() {
	notesList.Clear()
	notesList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		viewNote(storage[i])
		pagesMenu.SwitchToPage("Show Note")
	})
	for i, note := range storage {
		notesList.AddItem(fmt.Sprintf("[%s] %s", strings.ToUpper(note.GetType().String()), note.GetName()), "", rune(49+i), nil)
	}
}

func createMainMenu() {
	pagesMenu.AddPage("Menu", flexMain, true, true)
	pagesMenu.AddPage("Add Bank Card Note", formCardBankNote, true, false)
	pagesMenu.AddPage("Add Credential", formCredentialNote, true, false)
	pagesMenu.AddPage("Add Text Note", formTextNote, true, false)
	pagesMenu.AddPage("Add Binary Note", formBinaryNote, true, false)
	pagesMenu.AddPage("Show Note", modalViewNote, true, false)
}

func creteMainFlex() *tview.Flex {
	textMenu1 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(q) quit \n(l) load note \n(d) decrypt note")
	textMenu2 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(b) add bank card \n(c) add credential \n(t) add text \n(i) add binary")
	textMenu3 := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(a) create account \n(s) sign in")

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

func setInput() *tview.Box {
	return flexMain.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 113: //q
			app.Stop()
		case 108: //l
			addInfo("try load notes")
		case 100: //d
			addInfo("try decrypt note")
		case 98: //b
			formCardBankNote.Clear(true)
			addBankCardNote()
			pagesMenu.SwitchToPage("Add Bank Card Note")
		case 99: //c
			formCredentialNote.Clear(true)
			addCredentialNote()
			pagesMenu.SwitchToPage("Add Credential")
		case 116: //t
			formTextNote.Clear(true)
			addTextNote()
			pagesMenu.SwitchToPage("Add Text Note")
		case 105: //i
			formBinaryNote.Clear(true)
			addBinaryNote()
			pagesMenu.SwitchToPage("Add Binary Note")
		case 97: //a
			addInfo("create account")
		//TODO create account form
		case 115: //s
			addInfo("sign in")
		//TODO create sig in form
		default:
			addInfo(fmt.Sprintf("Invalid input %v", event.Rune()))
		}
		return event
	})
}

func save(note model.Noteable) {
	storage = append(storage, note)
	addNotesList()
	pagesMenu.SwitchToPage("Menu")
}

func viewNote(note model.Noteable) {
	modalViewNote.ClearButtons()
	modalViewNote.
		SetText(note.String()).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				pagesMenu.SwitchToPage("Menu")
			}
		})

}

func testData() {
	note1 := model.BankCardNote{
		Bank:         "Bank",
		Number:       "5556 4655 4655 4655 4655",
		Expiration:   "2022-04-01",
		Cardholder:   "John Smith",
		SecurityCode: "123",
	}
	note1.BaseNote = model.BaseNote{
		NameRecord: "Note1",
		Created:    time.Now().Unix(),
		Type:       model.CARD,
		MetaInfo:   []string{"test1 = test1\n", "test2 = test2\n", "сайт: www.test.com\n"},
	}

	note2 := model.BankCardNote{
		Bank:         "Bank",
		Number:       "5556 4611 4655 4655 4655",
		Expiration:   "2022-05-01",
		Cardholder:   "John Smith1",
		SecurityCode: "124",
	}
	note2.BaseNote = model.BaseNote{
		NameRecord: "Note2",
		Created:    time.Now().Unix(),
		Type:       model.CARD,
		MetaInfo:   []string{"test2 = test2\n", "test1 = test1\n", "test3 = test3\n"},
	}
	storage = append(storage, note1)
	storage = append(storage, note2)
}

//func viewCardBankNote(note model.BankCardNote) {
//	modalViewNote.
//		SetText(note.Number).
//		AddButtons([]string{"OK"}).
//		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "OK" {
//				pagesMenu.SwitchToPage("Menu")
//			}
//		})
//}
