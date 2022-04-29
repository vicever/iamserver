/**
 * @Author: zhangguangsheng
 * @Description:
 * @File: app
 * @Version: 1.0.0
 * @Date: 2022/4/28 23:04
 */

package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type RunFunc func(basename string) error
type Option func(*App)

type App struct {
	name     string
	basename string
	discript string
	runfunc  RunFunc
	args     cobra.PositionalArgs
	cmd      *cobra.Command
}

func NewApp(name string, basename string, opt ...Option) *App {
	fmt.Println("NewApp as: ", name)
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, optfunc := range opt {
		optfunc(a)
	}
	a.buildCommand()

	return a
}

func (app *App) buildCommand() {

	cmd := cobra.Command{
		Use:   app.basename,
		Short: app.name,
		Long:  app.discript,
		Args:  app.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	if app.runfunc != nil {
		cmd.RunE = app.runfunc(app.basename)
	}

	//  flagsets

	app.cmd = &cmd
}

func (app *App) Run() {
	if err := app.cmd.Execute(); err != nil {
		fmt.Printf("App Run Error: %v \n", err)
		os.Exit(1)
	}

}

func (app *App) WithOptions() {

}
func (app *App) WithDescription(cmdDesc string) {

}
func (app *App) WithDefautArgs() {

}
func (app *App) WithRunFunc(runFunc RunFunc) {

}
