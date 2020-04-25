package tinyscr

import (
	"golang.org/x/crypto/ssh/terminal"
	"strconv"
	"bufio"
	"fmt"
	"os"
)

var ERR error = nil

func Getch() (rune, error) {
	state, err := terminal.MakeRaw(0)

	if err != nil {
		return ' ', err
	}

	defer func() {
		ERR = terminal.Restore(0, state)
	}()

	if ERR != nil {
		return ' ', ERR
	}

	in := bufio.NewReader(os.Stdin)
	r, _, err := in.ReadRune()

	if err != nil {
		return ' ', err
	}

	return r, nil
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?121[?25h")
}

func Up(n int) {
	fmt.Print("\033[" + strconv.Itoa(n) + "A")
}

func Down(n int) {
	fmt.Print("\033[" + strconv.Itoa(n) + "B")
}

func Right(n int) {
	fmt.Print("\033[" + strconv.Itoa(n) + "C")
}

func Left(n int) {
	fmt.Print("\033[" + strconv.Itoa(n) + "D")
}

func Cursor(row, column int) {
	fmt.Print("\033[" + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H")
}
