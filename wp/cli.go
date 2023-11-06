package wp

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

var Version = "0.0.0"

var postsCmd = &cobra.Command{
	Use: "posts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List posts")
		os.Exit(0)
	},
}

var themeCmd = &cobra.Command{
	Use: "theme",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Theme")
		os.Exit(0)
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
		os.Exit(0)
	},
}

var wpCmd = &cobra.Command{
	Use:   "wp",
	Short: "Run blog",
	Run: func(cmd *cobra.Command, args []string) {

		listen, err := cmd.Flags().GetString("listen")

		if err != nil {
			panic(err)
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			panic(err)
		}

		if listen == "" {
			listen = ":8585"
		}

		log.Default().Println("Set admin password:", password)
		log.Default().Println("Listen:", listen)

		mux := chi.NewMux()
		LoadRoutes(mux, _defaultTheme)
		http.ListenAndServe(listen, mux)
	},
}

var _defaultTheme *Theme

func SetDefaultTheme(theme *Theme) {
	_defaultTheme = theme
}

var _adminFS embed.FS

func SetAdminFS(adminFS embed.FS) {
	_adminFS = adminFS
}

func SetupCLI() {

	wpCmd.CompletionOptions.HiddenDefaultCmd = true
	wpCmd.AddCommand(postsCmd)
	wpCmd.AddCommand(themeCmd)
	wpCmd.AddCommand(versionCmd)

	wpCmd.PersistentFlags().String("listen", "", "Listen address")
	wpCmd.PersistentFlags().String("theme", "", "Run with specific theme")
	wpCmd.PersistentFlags().String("password", "", "Admin password")

	if err := wpCmd.Execute(); err != nil {
		log.Default().Println(err)
	}
}
