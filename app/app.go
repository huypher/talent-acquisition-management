package app

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type App struct {
	server *http.Server
}

func (a *App) Run() {
	a.addCommands()
	Execute()
}

func (a *App) addCommands() {
	a.serveCmd()
}

func (a *App) serveCmd() {
	// serveCmd represents the serve command
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
			err := a.server.ListenAndServe()
			if err != nil {
				panic(err)
			}
		},
	}

	rootCmd.AddCommand(serveCmd)
}
