package internal

import (
	"context"
	"flux-notifier/pkg/notifier"
	"fmt"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for events",
	RunE: func(cmd *cobra.Command, _ []string) error {
		port, err := cmd.Flags().GetInt("svc-port")
		if err != nil {
			return fmt.Errorf("failed to get %v flag value. %w", "port", err)
		}
		return notifier.Listen(context.Background(), port)
	},
}

func init() {
	listenCmd.PersistentFlags().IntP("svc-port", "p", 8080, "service port")
	cobra.OnInitialize(func() {
		fillWithEnvVars(listenCmd.Flags())
	})
	rootCmd.AddCommand(listenCmd)
}
