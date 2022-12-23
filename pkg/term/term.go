package term

import (
	"fmt"
	"os"
	"strconv"
	"regexp"
	"os/exec"
	"os/signal"
	"syscall"
)

const (
	AttrReset     = "\x1b[0m"
	AttrBold      = "\x1b[1m"
	AttrItalics   = "\x1b[3m"
	AttrUnderline = "\x1b[4m"
	AttrBlink     = "\x1b[5m"

	AttrBlack   = "\x1b[30m"
	AttrRed     = "\x1b[31m"
	AttrGreen   = "\x1b[32m"
	AttrYellow  = "\x1b[33m"
	AttrBlue    = "\x1b[34m"
	AttrMagenta = "\x1b[35m"
	AttrCyan    = "\x1b[36m"
	AttrWhite   = "\x1b[37m"

	AttrGrey          = "\x1b[90m"
	AttrBrightRed     = "\x1b[91m"
	AttrBrightGreen   = "\x1b[92m"
	AttrBrightYellow  = "\x1b[93m"
	AttrBrightBlue    = "\x1b[94m"
	AttrBrightMagenta = "\x1b[95m"
	AttrBrightCyan    = "\x1b[96m"
	AttrBrightWhite   = "\x1b[97m"
)

var mode string

// 'stty size' output format is '<HEIGHT> <WIDTH>'
var sizeRegex = regexp.MustCompile("([0-9]*)\\s([0-9]*)\\s*")

func GetSize() (int, int) {
	bytes, err := exec.Command("stty", "-F", "/dev/tty", "size").Output()
	if err != nil {
		panic(err)
	}

	// Match the output format
	info := sizeRegex.FindStringSubmatch(string(bytes))

	var h, w int

	// Parse the strings
	h, err = strconv.Atoi(info[1])
	if err != nil {
		panic(err)
	}

	w, err = strconv.Atoi(info[2])
	if err != nil {
		panic(err)
	}

	return w, h
}

func Init(cbreak func(), resized func()) error {
	// Save the previous terminal attributes
	bytes, err := exec.Command("stty", "-F", "/dev/tty", "-g").Output()
	if err != nil {
		return err
	}

	mode = string(bytes[:len(bytes) - 1])

	// Ignore CTRL+C
	c1 := make(chan os.Signal, 1)
	signal.Notify(c1, os.Interrupt)

	go func() {
		for {
			<-c1
			cbreak()
		}
	}()

	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, syscall.SIGWINCH)

	go func() {
		for {
			<-c2
			resized()
		}
	}()

	//                                     No input echo      No buffered input
	exec.Command("stty", "-F", "/dev/tty", "-echo", "-echok", "-icanon", "min", "0").Run()

	return nil
}

func Restore() {
	exec.Command("stty", "-F", "/dev/tty", mode).Run()
}

func GetKey() byte {
	in := make([]byte, 1)
	os.Stdin.Read(in)

	return in[0]
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func MoveCursorUp(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vB", -by)
	} else {
		fmt.Printf("\x1b[%vA", by)
	}
}

func MoveCursorDown(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vA", -by)
	} else {
		fmt.Printf("\x1b[%vB", by)
	}
}

func MoveCursorRight(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vD", -by)
	} else {
		fmt.Printf("\x1b[%vC", by)
	}
}

func MoveCursorLeft(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vC", -by)
	} else {
		fmt.Printf("\x1b[%vD", by)
	}
}
