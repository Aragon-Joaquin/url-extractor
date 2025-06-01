package utils

import (
	"errors"
	"fmt"
	"strings"

	p "github.com/manifoldco/promptui"
)

func validateURL() func(string) error {
	return func(input string) error {
		//! check for protocol
		checkOne := []string{"http://", "https://"}
		if !ContainsAllValues(input, checkOne) {
			return errors.New("Needs to STARTS with one of the following: " + strings.Join(checkOne, ", "))
		}

		//! check for top-level domain
		checkTwo := []string{".com", ".org", ".net", ".io", ".dev", ".ai", ".app", ".co", ".info", ".tech", ".gov"}
		if !CheckTopLevelDomain(input, checkTwo) {
			return errors.New("Needs to ENDS with one of the following: " + strings.Join(checkTwo, ", "))
		}

		//! check if there's a path/subdirectory
		if strings.Count(input, "/") != 2 {
			return errors.New("Please, enter up to the top-level domain of the url (e.g.: 'https://google.com' ) ")
		}

		return nil
	}

}

func PromptInput() {
	prompt := &p.Prompt{
		Label:       "Enter the url you want to get all its URL's paths",
		Default:     "https://google.com",
		AllowEdit:   true,
		HideEntered: false,
		Validate:    validateURL(),
	}
	res, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	PrintColor(YELLOW, fmt.Sprintf("You choose %q\n", res))
}
