package mvc

import (
	"fmt"
	"net/mail"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/rivo/tview"
)

var (
	formRegistrationUser = tview.NewForm()
)

func addNewUser(cu *ControllerUI) {
	user := models.UserDto{}
	var password string
	formRegistrationUser.AddInputField("Username", "", 40,
		nil,
		func(text string) { user.Username = text })
	formRegistrationUser.AddPasswordField("Password", "", 40, rune(42),
		func(text string) { user.Password = text })
	formRegistrationUser.AddPasswordField("Confirm password", "", 40, rune(42),
		func(text string) { password = text })

	formRegistrationUser.AddInputField("Email", "", 40,
		nil,
		func(text string) { user.Email = text })

	formRegistrationUser.AddButton("Save", func() {
		if validateUser(user, password) {
			//TODO add register func
			cu.AddItemInfoList(fmt.Sprintf("The data of the user(%s) has been sent for registration. Pleas wait", user.Username))
			pagesMenu.SwitchToPage("Menu")
		}
	})

	formRegistrationUser.AddButton("Back", func() {
		pagesMenu.SwitchToPage("Menu")
	})
	formRegistrationUser.SetBorder(true).SetTitle("Registration").SetTitleAlign(tview.AlignLeft)
}

func validateUser(user models.UserDto, password string) bool {
	modalError.
		ClearButtons().
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				pagesMenu.SwitchToPage("Registration User")
			}
		}).SetTitle("Error")

	if user.Password != password {
		modalError.
			SetText("The passwords not equal")
		pagesMenu.SwitchToPage("Error")
	} else if len(user.Password) < 8 {
		modalError.
			SetText("Password must be at least 8 characters")
		pagesMenu.SwitchToPage("Error")
	} else if _, err := mail.ParseAddress(user.Email); err != nil {
		modalError.
			SetText("Email address not valid")
		pagesMenu.SwitchToPage("Error")
	} else if len(user.Username) < 5 {
		modalError.
			SetText("Username must be at least 5 characters")
		pagesMenu.SwitchToPage("Error")
	} else {
		return true
	}
	return false
}
