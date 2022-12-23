package wpm

import "math/rand"

const SeparatorNone byte = 0

var DefaultWords = []string{
	"hello",       "love",     "because",     "since",      "operation", "a",
	"meat",        "for",      "as",          "an",         "when",      "while",
	"hint",        "end",      "on",          "laugh",      "poor",      "none",
	"comprehense", "imperial", "fact",        "initialize", "liberal",   "programming",
	"surreal",     "apple",    "peach",       "little",     "amazing",   "wonderful",
	"grip",        "window",   "door",        "sky",        "toxicity",  "cause",
	"anthem",      "computer", "laptop",      "football",   "kick",      "word",
	"background",  "and",      "immediately", "the",        "dog",       "heat",
	"zero",        "one",      "two",         "three",      "four",      "five",
	"six",         "seven",    "eight",       "nine",       "ten",       "size",
	"skill",       "test",     "chat",        "chair",      "first",     "still",
	"am",          "near",     "radio",       "also",       "stand",     "conservative",
	"extra",       "manager",  "old",         "new",        "sand",      "moon",
	"big",         "tiny",     "red",         "orange",     "yellow",    "green",
	"cyan",        "cyan",     "turquoise",   "blue",       "violet",    "purple",
	"magenta",     "pink",     "white",       "grey",       "black",     "burger",
	"pizza",       "monkey",   "list",        "queue",      "board",     "tree",
	"bush",        "grass",    "gravel",      "pebble",     "terminal",  "house",
	"doorway",     "sink",     "bath",        "wardrobe",   "desk",      "table",
	"running",     "walking",  "sleeping",    "eating",     "reading",   "writing",
	"run",         "walk",     "sleep",       "eat",        "read",      "write",
	"public",      "private",  "structure",   "speak",      "speaking",  "free",
	"software",    "hardware", "program",     "web",        "internet",  "developer",
	"example",     "prevent",  "distribute",  "protect",    "right",     "left",
	"clear",       "clearly",  "explain",     "device",     "deny",      "accept",
	"design",      "graphic",  "graphics",    "safe",       "unsafe",    "easy",
	"hard",        "medium",   "freedom",     "slavery",    "rope",      "gore",
	"such",        "is",       "copy",        "move",       "condition", "incompatible",
	"compatible",  "precise",  "exact",       "same",       "term",      "factor",
	"expression",  "add",      "subtract",    "multiply",   "divide",    "modulo",
	"power",       "brackets", "general",     "license",    "law",       "paranthesis",
	"rule",        "work",     "working",     "refer",      "referring", "under",
	"over",        "cap",      "inside",      "outside",    "next",      "previous",
	"propaganda",  "user",     "propagate",   "advertise",  "interface", "appropriate",
	"interactive", "interact", "interacting", "notice",     "legal",     "illegal",
	"label",       "display",  "source",      "code",       "system",    "library",
	"executable",  "execute",  "anything",    "include",    "import",    "other",
	"than",        "then",     "where",       "if",         "package",   "corresponding",
	"correspond",  "need",     "not",         "cute",       "skirt",     "pants",
	"shirt",       "hat",      "tomorrow",    "boot",       "shoe",      "sleeve",
	"suit",        "outfit",   "cloth",       "clothing",   "clothes",   "provide",
	"copyright",   "cover",    "music",       "pop",        "push",      "rap",
	"rock",        "metal",    "jazz",        "classic",    "classical", "opera",
	"convey",      "blues",    "swing",       "instrument", "guitar",    "instrumental",
	"piano",       "violin",   "violent",     "trumpet",    "clarinet",  "finger",
	"hair",        "eye",      "mouth",       "shut",       "cloud",     "clout",
	"hype",        "light",    "darkness",    "dark",       "give",      "take",
	"second",      "whole",    "commercial",  "designated", "formation", "formatting",
	"form",        "access",   "equivalent",  "used",       "from",      "before",
	"after",       "be",       "because",     "of",         "you",       "have",
	"has",         "been",     "bee",         "too",        "go",        "going",
	"gonna",       "try",      "gotta",       "got",        "shadow",    "yesterday",
	"all",         "my",       "trouble",     "seemed",     "so",        "far",
	"away",        "comma",    "now",         "it",         "looks",     "as",
	"though",      "they",     "are",         "here",       "to",        "stay",
	"oh",          "i",        "believe",     "in",         "today",     "dot",
	"hang",        "long",     "sudden",      "why",        "something", "nothing",
	"no",          "yes",      "should",      "what",       "hide",      "show",
	"snow",        "holiday",  "snowman",     "carrot",     "pear",      "grape",
	"leaf",        "leave",    "leaves",      "exit",       "halt",      "entrance",
	"entry",       "enter",    "print",       "measure",    "minute",    "second",
	"hour",        "day",      "millisecond", "micro",      "nano",      "pico",
	"giga",        "tera",     "mega",        "milli",      "month",     "week",
	"weekly",      "year",     "yearly",      "daily",      "monthly",   "summer",
	"winter",      "spring",   "fall",        "january",    "february",  "march",
	"april",       "may",      "june",        "july",       "august",    "september",
	"october",     "november", "december",    "bill",       "above",     "below",
	"beyond",      "creep",    "creepy",      "scary",      "stalk",     "funny",
	"fun",         "funnest",  "funniest",    "agree",      "agreement", "ahead",
	"already",     "ready",    "article",     "ask",        "animal",    "cat",
	"dog",         "mouse",    "mice",        "sheep",      "back",      "beat",
	"authority",   "base",     "benefit",     "lose",       "loose",     "tight",
	"twenty",      "thirty",   "fourty",      "fifty",      "sixty",     "seventy",
	"eighty",      "ninety",   "hundred",     "thousand",   "million",   "billion",
	"trillion",    "negate",   "septillion",  "positive",   "negative",  "hexadecimal",
	"decimal",     "binary",   "third",       "fourth",     "fifth",     "sixth",
	"seventh",     "eighth",   "ninth",       "edit",       "editor",    "text",
	"much",        "lot",      "button",      "press",      "pressing",  "whatever",
	"good",        "bye",      "well",        "this",       "shiny",     "goodbye",
}

type Measurer struct {
	Words     []string
	Input     string
	Separator byte

	knownWords []string
}

func NewMeasurer(wordsList []string, separator byte) *Measurer {
	return &Measurer{knownWords: wordsList, Separator: separator}
}

func (m *Measurer) Type(input string) (correct bool, matched bool) {
	word := m.Words[0]
	if m.Separator != SeparatorNone {
		word += string(m.Separator)
	}

	slice := word[:len(m.Input + input)]
	if slice == m.Input + input {
		m.Input += input

		correct = true
	}

	matched = word == m.Input

	return
}

func (m *Measurer) GenWord() int {
	word   := m.knownWords[rand.Intn(len(m.knownWords))]
	m.Words = append(m.Words, word)

	return len(word)
}

func (m *Measurer) Next() {
	m.Words = m.Words[1:]
	m.Input = ""
}
