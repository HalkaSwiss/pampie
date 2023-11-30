package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/spf13/cobra"
	"log"
	"pampie/populate"

	_ "pampie/migrations"
)

func main() {
	app := pocketbase.New()
	app.RootCmd.AddCommand(&cobra.Command{
		Use: "populate",
		Run: func(cmd *cobra.Command, args []string) {
			err := populate.CreateClients(app)
			if err != nil {
				panic(err)
			}
		},
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
