package cli

import (
	"fmt"
)

type RegisterEventCommand struct {
	In  RegisterInEventCommand  `cmd aliases:"i" help:"Register a login event"`
	Out RegisterOutEventCommand `cmd aliases:"o" help:"Register a logout event"`
}

type RegisterEventCommandCommons struct {
	Time string `short:"t" help:"The time to be registered into the timesheet; by default, uses current system time"`
}

type RegisterOutEventCommand struct {
	RegisterEventCommandCommons
}

func (r *RegisterOutEventCommand) Run() error {
	fmt.Println("out")
	fmt.Printf("%#v\n", r.Time)
	return nil
}

type RegisterInEventCommand struct {
	RegisterEventCommandCommons
}

func (*RegisterInEventCommand) Run() error {
	fmt.Println("in")
	return nil
}
