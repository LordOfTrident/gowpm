package main

import (
	"fmt"
	"os"
	"time"
	"math"
	"math/rand"

	"github.com/LordOfTrident/gowpm/pkg/term"
	"github.com/LordOfTrident/gowpm/pkg/wpm"
)

var (
	estimated float64
	words     int
	mistakes  int
	wordsLen  int
	correct   bool

	started = false

	updateWords = true
	updateBar   = true
	barLen int

	testTime = 60
	timeLeft = testTime

	m *wpm.Measurer
)
const (
	correctAttr   = term.AttrReset + term.AttrBold + term.AttrBrightWhite
	incorrectAttr = term.AttrReset + term.AttrBold + term.AttrBrightRed
	untypedAttr   = term.AttrReset + term.AttrWhite
	promptAttr    = term.AttrReset + term.AttrBrightBlue
	numAttr       = term.AttrReset + term.AttrBrightYellow
)

var (
	width  int
	height int
)

func twoDecimalPlaces(num float64) float64 {
	return math.Round(num * 100) / 100
}

func earlyExit() {
	if started {
		fmt.Printf("\n\n\n Estimated WPM: %v%v%v\n", numAttr, estimated, term.AttrReset)
	}

	cleanup()
	os.Exit(0)
}

func cleanup() {
	term.ShowCursor()
	term.Restore()
}

func resized() {
	width, height = term.GetSize()
	width -= 3

	updateWords = true
	updateBar   = true

	if width < 15 {
		fmt.Println("\nWindow is too small")

		cleanup()
		os.Exit(1)
	}
}

func update() {
	for wordsLen + len(m.Words) < width + 1 {
		wordsLen += m.GenWord()
	}

	updateCursor := false
	if updateBar {
		if float64(testTime - timeLeft) > 0 {
			estimated = twoDecimalPlaces(float64(words) / float64(testTime - timeLeft) *
			                             float64(testTime))
		}

		renderBar()
		updateBar    = false
		updateCursor = true
	}

	if updateWords {
		renderWords()
		updateWords  = false
		updateCursor = true
	}

	if updateCursor {
		positionCursor()
	}

	input()
}

func input() {
	key := term.GetKey()
	if key == 0 {
		return
	}

	updateWords = true

	var matched bool
	correct, matched = m.Type(string(key))
	if matched {
		wordsLen -= len(m.Words[0])
		words ++

		updateBar = true
		m.Next()
	}

	if !correct {
		mistakes ++
	}
}

func renderWords() {
	typedAttr := correctAttr
	if !correct {
		typedAttr = incorrectAttr
	}

	fmt.Printf("\r%v> %v%v%v%v%v ", term.AttrReset, promptAttr, typedAttr, m.Input,
	           untypedAttr, m.Words[0][len(m.Input):])

	w := len(m.Words[0])
	for _, word := range m.Words[1:] {
		w += len(word) + 1

		if w >= width {
			word = word[:len(word) - (w - width)]
			fmt.Printf("%v-", word)

			break
		} else {
			fmt.Print(word)

			if w < width {
				fmt.Print(" ")
			}
		}
	}
}

func positionCursor() {
	fmt.Print("\r")
	term.MoveCursorRight(len(m.Input) + 2)
}

func renderBar() {
	fmt.Print("\n\n")
	defer term.MoveCursorUp(2)

	msg := fmt.Sprintf("%v Time left: %v%v%v   Typed words: %v%v%v   Estimated WPM: %v%v%v    ",
	                   term.AttrReset,
	                   numAttr, timeLeft,  term.AttrReset,
	                   numAttr, words,     term.AttrReset,
	                   numAttr, estimated, term.AttrReset)
	if len(msg) - len(term.AttrReset) * 4 - len(numAttr) * 3 > width {
		msg = msg[:width]
	}
	barLen = len(msg)

	fmt.Print(msg)
}

func timer() {
	for range time.Tick(1 * time.Second) {
		timeLeft --

		if timeLeft == 0 {
			break
		}

		updateBar = true
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())

	term.Init(earlyExit, resized)
	resized()
}

func main() {
	defer cleanup()
	m = wpm.NewMeasurer(wpm.DefaultWords, ' ')

	term.HideCursor()
	fmt.Printf("  %vPress any key to start%v", term.AttrBold, term.AttrReset)

	for term.GetKey() == 0 {}
	term.ShowCursor()

	started = true
	go timer()

	for {
		update()

		if timeLeft == 0 {
			renderBar()
			break
		}
	}

	fmt.Printf("\n\n\n WPM: %v%v%v\n", numAttr, words, term.AttrReset)
}
