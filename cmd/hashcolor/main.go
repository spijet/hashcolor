package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/spijet/hashcolor/pkg/utils"
)

func hex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func main() {
	c := utils.New(strings.Join(os.Args[1:], " "))
	t, s := utils.Tint(c), utils.Shade(c)
	term := os.Getenv("TERM_PROGRAM")
	if (term == "iTerm.app") {
		fmt.Printf("\033]1337;SetColors=fg=rgb:%s\033\\\033]1337;SetColors=bg=rgb:%s\033\\", hex(t), hex(s))
	} else {
		fmt.Printf("\033]10;#%s\007\033]11;#%s\007", hex(t), hex(s))
	}
}
