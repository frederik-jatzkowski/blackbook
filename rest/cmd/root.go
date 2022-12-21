package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/frederik-jatzkowski/blackbook/api"
	"github.com/frederik-jatzkowski/blackbook/database"
	"github.com/frederik-jatzkowski/blackbook/group"
	"github.com/frederik-jatzkowski/blackbook/mail"
	"github.com/frederik-jatzkowski/blackbook/user"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "blackbook-rest",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			db          *gorm.DB
			mailer      *mail.Service
			userService *user.Service
			apiServer   *api.Server
			err         error
		)

		// connect to db
		db, err = database.NewDatabase()
		if err != nil {
			return fmt.Errorf("error while creating database: %s", err)
		}

		// create mail service
		mailer, err = mail.NewService()
		if err != nil {
			return fmt.Errorf("error while creating mail service: %s", err)
		}

		// create user service
		userService, err = user.NewService(db, mailer)
		if err != nil {
			return fmt.Errorf("error while creating user service: %s", err)
		}

		// create group service
		groupService, err := group.NewService(db, userService)
		if err != nil {
			return fmt.Errorf("error while creating group service: %s", err)
		}

		// create api
		apiServer, err = api.NewServer(userService, groupService)
		if err != nil {
			return fmt.Errorf("error while creating server: %s", err)
		}

		log.Println("starting server")

		return apiServer.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
