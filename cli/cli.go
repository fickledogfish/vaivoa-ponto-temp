package cli

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/fickledogfish/vaivoa-timesheets-cli/net"
)

type App struct {
	kongKong *kong.Kong
	context  *kong.Context
}

func NewApp() (App, error) {
	var cmd cli

	k, err := kong.New(
		&cmd,
		kong.Name("vaivoa-timesheets"),
		kong.Description("Register work times in your terminal"),
	)
	if err != nil {
		return App{}, err
	}

	return App{
		kongKong: k,
	}, nil
}

func (app *App) Parse(args []string) error {
	context, err := app.kongKong.Parse(args)
	if err != nil {
		return err
	}

	app.context = context
	return nil
}

func (app *App) Configure() error {
	if app.context == nil {
		log.Fatal("Call App.Parse(args []string) before App.Configure()")
	}

	injectDependencies(
		app.context,
		[]dependency{
			{net.DefaultProvider{}, (*net.Provider)(nil)},
		},
	)

	return nil
}

func (app App) Run() error {
	if err := app.context.Run(); err != nil {
		return err
	}

	return nil
}

type cli struct {
	Report   ReportCommand        `cmd aliases:"rep,g" help:"Generate the report of the timesheet"`
	Register RegisterEventCommand `cmd aliases:"reg,r" help:"Record an event in the timesheet"`
}

type dependencyProvider interface{}
type dependencyInterface interface{}

type dependency struct {
	Provider  dependencyProvider
	Interface dependencyInterface
}

func injectDependencies(context *kong.Context, dependencies []dependency) {
	for _, dependency := range dependencies {
		context.BindTo(dependency.Provider, dependency.Interface)
	}
}
