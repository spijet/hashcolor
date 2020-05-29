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
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func main() {
	c := utils.New(strings.Join(os.Args[1:], " "))
	t, s := utils.Tint(c), utils.Shade(c)
	fmt.Printf("\033]10;%s\007\033]11;%s\007", hex(t), hex(s))
}
