package initilazers

import "github.com/VaheMuradyan/Tasks/go-gorm-jwt/models"

func init() {
	LoadEnvVariables()
	ConnectToDb()
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
