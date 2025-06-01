package utils

import (
	"fmt"
	"os"
)

type Colors string

const (
	RED    Colors = "\033[0;31m"
	GREEN  Colors = "\033[0;32m"
	YELLOW Colors = "\033[0;33m"
	BLUE   Colors = "\033[0;34m"
	PURPLE Colors = "\033[0;35m"
	CYAN   Colors = "\033[0;36m"
	WHITE  Colors = "\033[0;37m"
	NONE   Colors = "\033[0m"
)

func PrintColor(color Colors, message string) {
	fmt.Fprintf(os.Stdout, "%s", string(color)+message+string(NONE))
}
