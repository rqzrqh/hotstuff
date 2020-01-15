package main

import (
	"flag"
	"strings"

	"github.com/relab/hotstuff/pkg/basichotstuff"
)

func main() {
	leader := flag.Bool("leader", false, "Designates this replica as leader")

	addresses := flag.String("nodes", "", "Semicolon separated list of addresses to connect to.")
	flag.Parse()

	nodes := strings.Split(*addresses, ";")

	if *leader {
		l := &basichotstuff.Leader{}
		l.RunHotStuff()
	} else {

	}
}