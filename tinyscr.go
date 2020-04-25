package tinyscr

import (
	"golang.org/x/crypto/ssh/terminal"
	"strconv"
	"bufio"
	"fmt"
	"os"
)

func Getch() rune {
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	in := bufio.NewReader(os.Stdin)
	r, _, err := in.ReadRune()

	if err != nil {
		log.Println("STDIN:", err)
	}

	return r
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
