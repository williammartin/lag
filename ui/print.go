package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/tscolari/lag/parser"
)

func printEntries(v *gocui.View, entries parser.Entries) {
	for _, entry := range entries {
		printEntry(v, entry)
	}
}

func printEntry(v *gocui.View, entry *parser.Entry) {
	date := entry.Data.Timestamp.Format(time.RFC3339)
	message := entry.Data.Message
	if entry.Errored {
		message = redText(entry.Data.Message)
	}

	fmt.Fprintf(v, "%s [%s] (%d) %s \n",
		blueText(date),
		yellowText(entry.Data.Session),
		len(entry.Children),
		message)
}
