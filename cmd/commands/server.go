package commands

import (
	"futurisme-api/internal/server"

	"github.com/spf13/cobra"
)

var isDev bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Menjalankan HTTP Server API",
	Long:  `Menjalankan server backend Futurisme API menggunakan Fiber.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Panggil logika server yang sudah kita pindahkan tadi
		server.RunServer(isDev)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Mendaftarkan flag --dev (boolean)
	// Usage: go run main.go start --dev
	startCmd.Flags().BoolVarP(&isDev, "dev", "d", false, "Jalankan server dalam mode development (AutoMigrate on)")
}
