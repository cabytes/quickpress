package main

import (
	"cabytes/zine"
	"cabytes/zine/themes/light"
	"fmt"
	"log"
	"net/http"
	"os"

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

type AdminUser struct {
	User     string
	Pass     string
	UserName string
}

func (au *AdminUser) ID() string       { return "1" }
func (au *AdminUser) Username() string { return au.User }
func (au *AdminUser) Password() string { return au.Pass }
func (au *AdminUser) Name() string     { return au.UserName }

var zineCmd = &cobra.Command{
	Use:   "zine",
	Short: "Run blog",
	Run: func(cmd *cobra.Command, args []string) {

		user, _ := cmd.Flags().GetString("username")
		pass, _ := cmd.Flags().GetString("username")
		port, _ := cmd.Flags().GetString("listen")
		data, _ := cmd.Flags().GetString("data")

		app, err := zine.New(

			// Set data path for storage
			zine.DataPath(data),

			zine.LoadTheme("../themes/light/", light.Files),

			// Hook users
			zine.AuthHook(func(username, password string) zine.User {
				if username == user && password == pass {
					return &AdminUser{
						User:     user,
						Pass:     pass,
						UserName: "admin",
					}
				}
				return nil
			}),
		)

		if err != nil {
			panic(err)
		}

		if err := http.ListenAndServe(port, app); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {

	zineCmd.CompletionOptions.HiddenDefaultCmd = true
	zineCmd.AddCommand(postsCmd)
	zineCmd.AddCommand(themeCmd)
	zineCmd.AddCommand(versionCmd)

	zineCmd.PersistentFlags().String("listen", "", "Listen address")
	zineCmd.PersistentFlags().String("username", "", "Default username")
	zineCmd.PersistentFlags().String("password", "", "Default username password")
	zineCmd.PersistentFlags().String("data", "", "Set data folder path")

	if err := zineCmd.Execute(); err != nil {
		log.Default().Println(err)
	}
}
