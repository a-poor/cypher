package main

import (
	"os"

	"github.com/alecthomas/participle/v2/ebnf"
	log "github.com/sirupsen/logrus"
)

const filePath = "cypher.ebnf"

func main() {
	log.Info("Starting")

	log.Info("Opening file")
	f, err := os.Open(filePath)
	if err != nil {
		log.WithField("error", err).Panic("Error opening file")
		return
	}
	defer f.Close()
	log.Info("Successfully opened file")

	log.Info("Parsing file")
	e, err := ebnf.Parse(f)
	if err != nil {
		log.WithField("error", err).Error("Error parsing file as EBNF")
		return
	}

	log.WithField("nProductions", len(e.Productions)).Info("Parsed Cypher EBNF")
}
