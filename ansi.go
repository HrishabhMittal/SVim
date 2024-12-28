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
const BLINKING_BLOCK_CURSOR = "\033[1 q"
const STEADY_BLOCK_CURSOR = "\033[2 q"
const BLINKING_UNDERLINE_CURSOR = "\033[3 q"
const STEADY_UNDERLINE_CURSOR = "\033[4 q"
const BLINKING_THIN_CURSOR = "\033[5 q"
const STEADY_THIN_CURSOR = "\033[6 q"
const RESET_COLOR = "\033[0m"
func SetFGColor(num int) {
    fmt.Printf("\033[38;5;%dm",num)
}

func SetBGColor(num int) {
    fmt.Printf("\033[48;5;%dm",num)
}
func moveTo(i int,j int) {
    fmt.Printf("\033[%d;%dH",i+1,j+1)
}
