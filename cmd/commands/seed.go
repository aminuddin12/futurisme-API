package commands

import (
	"log"

	"futurisme-api/config"
	"futurisme-api/internal/modules/user/entity"
	"futurisme-api/pkg/database"
	"futurisme-api/pkg/utils/security"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Mengisi database dengan data dummy",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		db := database.NewPostgresDatabase(cfg)

		log.Println("🌱 Starting database seeding...")
		runSeeder(db)
		log.Println("✅ Database seeding completed.")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func runSeeder(db *gorm.DB) {
	// Logika seeder yang sama seperti sebelumnya
	defaultPassword := "rahasia123"
	hashedPassword, _ := security.HashPassword(defaultPassword)

	users := []entity.User{
		{
			Name: "Admin Futurisme", Email: "admin@futurisme.com", Password: hashedPassword, Role: "admin", Phone: "08111111111",
		},
		{
			Name: "User Biasa", Email: "user@futurisme.com", Password: hashedPassword, Role: "user", Phone: "08222222222",
		},
	}

	for _, user := range users {
		var existingUser entity.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&user)
				log.Printf("✨ Created user: %s", user.Email)
			}
		}
	}
}
