package wp

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

var posts = &cobra.Command{
	Use: "posts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List posts")
		os.Exit(0)
	},
}

var run = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {

		pass, _ := cmd.Flags().GetString("password")

		fmt.Println("Starting with password:", pass)

		mux := chi.NewMux()
		Load(mux, _defaultTheme)
		http.ListenAndServe(":8085", mux)
	},
}

var wp = &cobra.Command{
	Use: "wp",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var _defaultTheme *Theme

func SetDefaultTheme(theme *Theme) {
	_defaultTheme = theme
}

func SetupCLI() {

	wp.AddCommand(posts)
	wp.AddCommand(run)

	run.PersistentFlags().String("password", "", "Admin password")

	if err := wp.Execute(); err != nil {
		log.Default().Println(err)
	}
}
