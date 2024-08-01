package mvc

import (
	"fmt"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/model"
	"github.com/rivo/tview"
)

var (
	formCredentialNote = tview.NewForm()
)

func addCredentialNote() {
	note := model.CredentialNote{}
	note.Type = model.CREDENTIAL
	note.Created = time.Now().Unix()
	var metaInfo string
	formCredentialNote.AddInputField("Username", "", 40, nil, func(text string) { note.Username = text })
	formCredentialNote.AddInputField("Password", "", 40, nil, func(text string) { note.Password = text })
	formCredentialNote.AddTextArea("Additional information", "", 40, 0, 0, func(text string) { metaInfo = text })
	formCredentialNote.AddInputField("Save as", "", 40, nil, func(text string) { note.NameRecord = text })

	formCredentialNote.AddButton("Save", func() {
		note.MetaInfo = append(note.MetaInfo, metaInfo)
		addInfo(fmt.Sprintf("add credential note as %s", note.NameRecord))
		save(&note)
	})

	formCredentialNote.AddButton("Back", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formCredentialNote.SetBorder(true).SetTitle("New credential note").SetTitleAlign(tview.AlignLeft)
}
