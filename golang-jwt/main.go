package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model User
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// Model Book
type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	UserID uint   `json:"user_id"`
	// User   User   `json:"user"`
}

var DB *gorm.DB
var jwtSecret = []byte("your_secret_key") // Ganti dengan kunci rahasia yang lebih aman

// Fungsi untuk menghubungkan ke database PostgreSQL
func SetupDatabase() {
	dsn := "host=localhost user=postgres password=admin dbname=golang_jwt port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Migrasi model ke database
	DB.AutoMigrate(&User{}, &Book{})
}

// Fungsi untuk menghasilkan JWT token
func GenerateJWT(user User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Fungsi untuk memverifikasi dan meng-decode JWT
func VerifyJWT(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	// Hapus prefix "Bearer "
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Extract claims (user info)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Set user ID in context for future use (e.g., in handlers)
	c.Locals("userID", claims["sub"])

	return c.Next()
}

// Fungsi untuk mendapatkan semua pengguna (dengan autentikasi)
func GetUsers(c *fiber.Ctx) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Unable to fetch users",
		})
	}
	return c.JSON(users)
}

// Fungsi untuk menambah pengguna baru
func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.Password = string(hash)

	// Simpan user ke database
	if err := DB.Create(user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	return c.Status(201).JSON(user)
}

// Fungsi untuk login dan mendapatkan JWT token
func Login(c *fiber.Ctx) error {
	login := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	if err := c.BodyParser(login); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	// Cari user berdasarkan email
	var user User
	if err := DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := GenerateJWT(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// Fungsi untuk mendapatkan semua buku (hanya jika sudah terautentikasi)
func GetBooks(c *fiber.Ctx) error {
	var books []Book
	if err := DB.Preload("User").Find(&books).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Unable to fetch books",
		})
	}
	return c.JSON(books)
}

// Fungsi untuk menambah buku baru (hanya jika sudah terautentikasi)
func CreateBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	// Ambil ID user dari context
	userID := c.Locals("userID").(float64)

	// Periksa apakah user dengan ID tersebut ada
	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Set UserID dari buku
	book.UserID = uint(userID)

	// Simpan buku ke database
	if err := DB.Create(book).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}
	return c.Status(201).JSON(book)
}

func main() {
	// Setup database
	SetupDatabase()

	// Initialize Fiber app
	app := fiber.New()

	// Middleware logger
	app.Use(logger.New())

	// Route untuk login
	// user: dwiirawan | pass: admin
	app.Post("/login", Login)

	// Middleware autentikasi JWT
	app.Use(VerifyJWT)

	// Routes yang membutuhkan autentikasi
	app.Get("/users", GetUsers)
	app.Post("/users", CreateUser)
	app.Get("/books", GetBooks)
	app.Post("/books", CreateBook)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
