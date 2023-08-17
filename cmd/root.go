package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/aryayunanta-ralali/shorty/cmd/http"
	"github.com/aryayunanta-ralali/shorty/cmd/migration"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/generator"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
    logger.SetJSONFormatter()
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

    migrateCmd := &cobra.Command{
        Use:   "db:migrate",
        Short: "database migration",
        Run: func(c *cobra.Command, args []string) {
            migration.MigrateDatabase()
        },
    }

    migrateCmd.Flags().BoolP("version", "", false, "print version")
    migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
    migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
    migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
    migrateCmd.Flags().BoolP("guide", "", false, "print help")

	cmd := []*cobra.Command{
		{
			Use:   "serve-http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start(ctx)
			},
		},
		{
			Use:   "generate-template",
			Short: "Run Generate Template",
			Run: func(cmd *cobra.Command, args []string) {
				generator.BuildProject()
			},
		},
		migrateCmd,
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
