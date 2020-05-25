package tinyscr

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"fmt"
)

func Getch() (rune, error) {
	state, err := terminal.MakeRaw(0)

	if err != nil {
		return ' ', err
	}

	defer terminal.Restore(0, state)

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)

	return rune(b), nil
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
