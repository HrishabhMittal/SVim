package main

import (
	"fmt"
)
const ARROW_UP = 65
const ARROW_DOWN = 66
const ARROW_RIGHT = 67
const ARROW_LEFT = 68
const HOME = "\033[H"
const CLEAR = "\033[2J"
const INVIS_CURSOR = "\033[?25l"
const VIS_CURSOR = "\033[?25h"
const OPEN_ALT_BUFFER = "\033[?1049h"
const CLOSE_ALT_BUFFER = "\033[?1049l"
func moveTo(i int,j int) {
    fmt.Printf("\033[%d;%dH",i+1,j+1)
}
