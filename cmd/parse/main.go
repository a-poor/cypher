package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/alecthomas/participle/v2/ebnf"
	log "github.com/sirupsen/logrus"
)

var (
	// Default path to the EBNF file to parse
	filePath = "cypher.ebnf"
)

var (
	// Regex to remove comments
	commentRegex = regexp.MustCompile(`\(\*[.\n]*\*\)`)

	// Regex to replace semi-colons at end of production with period
	termRegex = regexp.MustCompile(`;\s*(\n|$)`)

	// Regex to replace commas

	// Regex to replace single quotes with double quotes
	quoteRegex = regexp.MustCompile(`'`)

	// Regex to replace square brackets
	sqBracketOpenRegex  = regexp.MustCompile(`[^"]\[`)
	sqBracketCloseRegex = regexp.MustCompile(`\[[^"]`)

	// Regex to replace curly braces
	curlyBraceOpenRegex      = regexp.MustCompile(`[^"]\{`)
	curlyBraceCloseDashRegex = regexp.MustCompile(`\}-[^"]`)
	curlyBraceCloseRegex     = regexp.MustCompile(`\}[^"]`)
)

func cleanEBNF(r io.Reader) (io.Reader, error) {
	// Read in the source data
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("unable to read from reader: %w", err)
	}

	// Remove the comments
	b = commentRegex.ReplaceAll(b, []byte(""))

	// Replace semi-colons at end of productions with periods
	b = termRegex.ReplaceAll(b, []byte(".\n\n"))

	// Replace single quotes with double quotes
	b = quoteRegex.ReplaceAll(b, []byte(`"`))

	// Replace square brackets
	b = sqBracketOpenRegex.ReplaceAll(b, []byte(` (`))
	b = sqBracketCloseRegex.ReplaceAll(b, []byte(`)? `))

	// Replace curly braces
	b = curlyBraceOpenRegex.ReplaceAll(b, []byte(` (`))
	b = curlyBraceCloseDashRegex.ReplaceAll(b, []byte(`)+ `))
	b = curlyBraceCloseRegex.ReplaceAll(b, []byte(`)* `))

	return bytes.NewBuffer(b), nil
}

func init() {
	flag.StringVar(&filePath, "file", filePath, "Path to the EBNF file to parse")
	flag.Parse()
}

const stmt = `

SENTENCE = [SP], Statement, [[SP], ";"], [SP], EOI .

WORD = { LETTER }- ;

WS = { SPACE }- ;

SPACE = ' ' ;

LETTER = A
	| B
	| C
	| D
	| E
	| F
	| G
	| H
	| I
	| J
	| K
	| L
	| M
	| N
	| O
	| P
	| Q
	| R
	| S
	| T
	| U
	| V
	| W
	| X
	| Y
	| Z
	;

Dash = '-'
     | '­'
     | '‐'
     | '‑'
     | '‒'
     | '–'
     | '—'
     | '―'
     | '−'
     | '﹘'
     | '﹣'
     | '－'
     ;

A = 'A' | 'a' ;

B = 'B' | 'b' ;

C = 'C' | 'c' ;

D = 'D' | 'd' ;

E = 'E' | 'e' ;

F = 'F' | 'f' ;

G = 'G' | 'g' ;

H = 'H' | 'h' ;

I = 'I' | 'i' ;

K = 'K' | 'k' ;

L = 'L' | 'l' ;

M = 'M' | 'm' ;

N = 'N' | 'n' ;

O = 'O' | 'o' ;

P = 'P' | 'p' ;

Q = 'Q' | 'q' ;

R = 'R' | 'r' ;

S = 'S' | 's' ;

T = 'T' | 't' ;

U = 'U' | 'u' ;

V = 'V' | 'v' ;

W = 'W' | 'w' ;

X = 'X' | 'x' ;

Y = 'Y' | 'y' ;`

func main() {
	log.Info("Starting")

	// log.Info("Opening file")
	// f, err := os.Open(filePath)
	// if err != nil {
	// 	log.WithField("error", err).Panic("Error opening file")
	// 	return
	// }
	// defer f.Close()
	// log.Info("Successfully opened file")

	// log.Info("Removing comments")
	// r, err := removeComments(f)
	// if err != nil {
	// 	log.WithField("error", err).Panic("Error removing comments")
	// 	return
	// }
	// log.Info("Successfully removed comments")

	// r := strings.NewReader(`
	// Cypher = [SP], Statement, [[SP], ";"], [SP], EOI .
	// `)

	f := strings.NewReader(stmt)

	r, err := cleanEBNF(f)
	if err != nil {
		log.WithField("error", err).Panic("Error reformatting EBNF")
		return
	}
	log.Info("Successfully cleaned EBNF")

	b, _ := io.ReadAll(r)
	fmt.Printf("Cleaned EBNF:\n%s\n", string(b))
	r = strings.NewReader(string(b))

	log.Info("Parsing file")
	e, err := ebnf.Parse(r)
	if err != nil {
		log.WithField("error", err).Error("Error parsing file as EBNF")
		return
	}

	log.WithField("nProductions", len(e.Productions)).Info("Parsed Cypher EBNF")
}
