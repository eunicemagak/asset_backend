module gitlab.ci.emalify.com/roamtech/asset_be

go 1.17

require (
	github.com/joho/godotenv v1.4.0
	gorm.io/gorm v1.23.1
)

require github.com/lib/pq v1.10.4 // indirect

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofiber/fiber/v2 v2.27.0
	github.com/google/uuid v1.3.0
	github.com/klauspost/compress v1.14.2 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	gorm.io/driver/postgres v1.3.1
)
