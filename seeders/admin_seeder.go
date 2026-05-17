package seeders

import (
    "log"

    "github.com/binojmadhuranga/BrainGrid-go/constants"
    "github.com/binojmadhuranga/BrainGrid-go/models"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {

    adminEmail := "admin@braingrid.com"
    adminPassword := "admin123"

    var existingAdmin models.User

    err := db.Where("email = ?", adminEmail).First(&existingAdmin).Error

    if err == nil {
        log.Println("Admin already exists")
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword(
        []byte(adminPassword),
        bcrypt.DefaultCost,
    )

    admin := models.User{
        Name:     "BrainGrid Admin",
        Email:    adminEmail,
        Password: string(hashedPassword),
        Role:     constants.ROLE_ADMIN,
    }

    if err := db.Create(&admin).Error; err != nil {
        log.Println("Failed to seed admin:", err)
        return
    }

    log.Println("Admin account seeded successfully")
}