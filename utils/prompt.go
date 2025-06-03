package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	p "github.com/manifoldco/promptui"
)

var URLProtocols = []string{"http://", "https://"}
var URLTopDomain = []string{".com", ".org", ".net", ".io", ".dev", ".ai", ".app", ".co", ".info", ".tech", ".gov"}
var WWWPrefix string = "www."

func validateURL() func(string) error {
	return func(input string) error {
		//! check for protocol
		if !ContainsAllValues(input, URLProtocols) {
			return errors.New("Needs to STARTS with one of the following: " + strings.Join(URLProtocols, ", "))
		}

		//! check for top-level domain
		if !CheckTopLevelDomain(input, URLTopDomain) {
			return errors.New("Needs to ENDS with one of the following: " + strings.Join(URLTopDomain, ", "))
		}

		//! check if there's a path/subdirectory
		if strings.Count(input, "/") != 2 {
			return errors.New("Please, enter up to the top-level domain of the url (e.g.: 'https://google.com' ) ")
		}

		return nil
	}

}

func PromptInput() (string, error) {
	prompt := &p.Prompt{
		Label:       "Enter the url you want to get all its URL's paths",
		Default:     "https://google.com",
		AllowEdit:   true,
		HideEntered: false,
		Validate:    validateURL(),
	}
	res, err := prompt.Run()

	if err != nil {
		return "", err
	}
	PrintColor(YELLOW, fmt.Sprintf("You choose %q\n", res))

	return res, nil
}

func PromptConfirm(MaxCount int) bool {
	prompt := &p.Prompt{
		Label:     "Reached " + strconv.Itoa(MaxCount) + " requests. Wanna keep searching? (making too much requests would probably getting you timed-out or ip banned)",
		AllowEdit: false,
		IsConfirm: true,
	}
	res, err := prompt.Run()

	// if its anything except 'y' or 'Y' it just throws an error
	if err != nil {
		return false
	}

	// its not necessary but whatever
	return res == "Y" || res == "y"
}
