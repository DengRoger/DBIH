package args

import (
	"flag"
)

type CmdArgs struct {
	DB    string
	TABLE string
}

var Cmd CmdArgs

func ParseCmd() {
	flag.StringVar(&Cmd.DB, "db", "", "db operation")
	flag.StringVar(&Cmd.TABLE, "table", "", "table name")
	flag.Parse()
}