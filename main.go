package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/KyleSmith19091/sqlite/metacommand"
	"github.com/xwb1989/sqlparser"
)

const (
	VERSION = "0.1"
)

func prologue() {
	fmt.Printf("KSQLite version %s %s\n", VERSION, string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Enter .help for usage hints")
}

func prompt() (string, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read input from stdin")
	}

	line, _ = strings.CutSuffix(line, "\n")
	return line, nil
}

func main() {
	prologue()
	for ;; {
		out, err := prompt()
		if err != nil {
			os.Exit(1)
		}
	
		if len(out) == 0 {
			continue
		}

		if out[0] == metacommand.META_COMMAND_PREFIX {
			if err := metacommand.Handle(out); err != nil {
				slog.Error("could not handle meta command", "err", err.Error())
			}
		}

		stmt, err := sqlparser.Parse(out)
		if err != nil {
			slog.Error("could not parse the SQL statement", "err", err.Error())
		}
		
	}
}
