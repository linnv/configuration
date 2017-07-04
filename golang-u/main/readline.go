package main

import (
	"github.com/chzyer/readline"

	"os"
)

const (
	UPDATE   = "upd"
	DELETE   = "del"
	GENERATE = "gen"
)

//@TODO catch signal c-c c-d and invoke del action
func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		println(line)
		switch line {
		case DELETE:
		case GENERATE:
		case UPDATE:
		default:
			usage()

		}
	}
}

func usage() {
	os.Stdout.Write(append([]byte(
		`
gen:geneate data in memory and write data to db
upd:update data in memory
del: delete data in db
		`), '\n'))
}
