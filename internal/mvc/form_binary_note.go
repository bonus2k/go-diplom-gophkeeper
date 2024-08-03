package mvc

import (
	"fmt"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

var (
	formBinaryNote = tview.NewForm()
)

func addBinaryNote(cu *ControllerUI) {
	note := models.BinaryNote{}
	note.Id, _ = uuid.NewUUID()
	note.Type = models.BINARY
	note.Created = time.Now().Unix()
	var metaInfo string
	var textArea string
	formBinaryNote.AddTextArea("Binary data", "", 40, 0, 0,
		func(text string) { textArea = text })
	formBinaryNote.AddTextArea("Additional information", "", 40, 0, 0,
		func(text string) { metaInfo = text })
	formBinaryNote.AddInputField("Save as", "", 40,
		nil,
		func(text string) { note.NameRecord = text })

	formBinaryNote.AddButton("Save", func() {
		note.MetaInfo = append(note.MetaInfo, metaInfo)
		note.Binary = []byte(textArea)
		cu.AddItemInfoList(fmt.Sprintf("The note has been saved with the binary data: %s", note.NameRecord))
		cu.AddNote(&note)
		pagesMenu.SwitchToPage("Menu")
	})

	formBinaryNote.AddButton("Back", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formBinaryNote.SetBorder(true).SetTitle("New binary note").SetTitleAlign(tview.AlignLeft)
}
