package mvc

import (
	"fmt"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/rivo/tview"
)

var (
	formAuthorization = tview.NewForm()
)

func authorizationUser(cu *ControllerUI) {
	user := models.UserDto{}
	formAuthorization.AddInputField("Username", "", 40,
		nil,
		func(text string) { user.Username = text })
	formAuthorization.AddPasswordField("Password", "", 40, rune(42),
		func(text string) { user.Password = text })

	formAuthorization.AddButton("Sig in", func() {
		//TODO add registration
		cu.AddItemInfoList(fmt.Sprintf("Welcome back User!"))
		pagesMenu.SwitchToPage("Menu")

	})

	formAuthorization.AddButton("Cancel", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formAuthorization.SetBorder(true).SetTitle("Sig in").SetTitleAlign(tview.AlignLeft)
}
