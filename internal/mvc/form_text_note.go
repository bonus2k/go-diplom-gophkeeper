package mvc

import (
	"fmt"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

var (
	formTextNote = tview.NewForm()
)

func addTextNote(cu *ControllerUI) {
	note := models.TextNote{}
	note.Id, _ = uuid.NewUUID()
	note.Type = models.TEXT
	note.Created = time.Now().Unix()
	var metaInfo string
	var textArea string
	formTextNote.AddTextArea("Text data", "", 40, 0, 0,
		func(text string) { textArea = text })
	formTextNote.AddTextArea("Additional information", "", 40, 0, 0,
		func(text string) { metaInfo = text })
	formTextNote.AddInputField("Save as", "", 40,
		nil,
		func(text string) { note.NameRecord = text })

	formTextNote.AddButton("Save", func() {
		note.MetaInfo = append(note.MetaInfo, metaInfo)
		note.Text = textArea
		cu.AddItemInfoList(fmt.Sprintf("The note has been saved with the text data:  %s", note.NameRecord))
		cu.AddNote(&note)
		pagesMenu.SwitchToPage("Menu")
	})

	formTextNote.AddButton("Back", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formTextNote.SetBorder(true).SetTitle("New text note").SetTitleAlign(tview.AlignLeft)
}
