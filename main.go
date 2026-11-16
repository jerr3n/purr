// SPDX-License-Identifier: MPL-2.0
package main

import (
	"os"

	"github.com/jerr3n/purr/internal/kitty"
)

func main() {
	data, _ := os.ReadFile("./test/kitty.png")
	kitty.Image(data, os.Stdout)
}
