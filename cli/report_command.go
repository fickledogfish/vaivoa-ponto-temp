package cli

import (
	"fmt"
)

type ReportCommand struct{}

func (ReportCommand) Run() error {
	fmt.Println("report")

	return nil
}
