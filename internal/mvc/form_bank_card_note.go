package mvc

import (
	"fmt"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

var (
	formCardBankNote = tview.NewForm()
)

func addBankCardNote(cu *ControllerUI) {
	note := models.BankCardNote{}
	note.Id, _ = uuid.NewUUID()
	note.Type = models.CARD
	note.Created = time.Now().Unix()
	var metaInfo string
	var cardNumber string
	formCardBankNote.AddInputField("Bank name", "", 40,
		nil,
		func(text string) { note.Bank = text })
	formCardBankNote.AddInputField("Card number", "", 40,
		func(textToCheck string, lastChar rune) bool { return lastChar >= '0' && lastChar <= '9' },
		func(text string) { cardNumber = text })
	formCardBankNote.AddInputField("Expiration", "", 40,
		nil,
		func(text string) { note.Expiration = text })
	formCardBankNote.AddInputField("Cardholder name", "", 40,
		nil,
		func(text string) { note.Cardholder = text })
	formCardBankNote.AddInputField("Security code", "", 40,
		func(textToCheck string, lastChar rune) bool { return len(textToCheck) <= 3 },
		func(text string) { note.SecurityCode = text })
	formCardBankNote.AddTextArea("Additional information", "", 40, 0, 0,
		func(text string) { metaInfo = text })
	formCardBankNote.AddInputField("Save as", "", 40,
		nil,
		func(text string) { note.NameRecord = text })

	formCardBankNote.AddButton("Save", func() {
		note.MetaInfo = append(note.MetaInfo, metaInfo)
		note.Number = formatCardNumber(cardNumber)
		cu.AddItemInfoList(fmt.Sprintf("The note has been saved with the card data of the bank:%s", note.NameRecord))
		cu.AddNote(&note)
		pagesMenu.SwitchToPage("Menu")
	})

	formCardBankNote.AddButton("Back", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formCardBankNote.SetBorder(true).SetTitle("New bank card note").SetTitleAlign(tview.AlignLeft)
}

func formatCardNumber(text string) string {
	str := ""
	i := 1
	for index, char := range text {
		str += string(char)
		if i%4 == 0 && index+1 != len(text) {
			str += "-"
		}
		i++
	}
	return str
}
