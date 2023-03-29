package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"gorm.io/gorm"
)

var (
	config = Config{
		dsn: os.Getenv("DATABASE_URL") + "?sslmode=require",
	}
	db = func() *gorm.DB {
		
		db, err := connect(config)
		if err != nil {
			panic(err)
		}
		if err := migrateAll(db); err != nil {
			panic(err)
		}
		return db
	}()
	app      = fiber.New()
	validate = validator.New()
	
)
func init() {
    // if port := os.Getenv("PORT"); port != "" {
    //     app.Listen(":" + port)
    // } else {
    //     app.Listen(":8080")
    // }
}






func main() {
	app.Use(cors.New())
	fmt.Println(fmt.Sprintln("Connected to", "server", "as", "hamid"))

	app.Get("/", getDefault)

	sessions := app.Group("/auth")
	sessions.Post("/signup", signUpUser)
	sessions.Post("/login", loginUser)

	user := app.Group("/user")
	user.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))

	user.Get("/", getUser)
	user.Put("/", updateUser)
	user.Delete("/", deleteUser)

	notebooks := app.Group("/notebooks")
	notebooks.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))

	notebooks.Get("/", getNotebook)
	notebooks.Get("/all", getAllNotebooks)
	notebooks.Post("/", createNotebook)
	notebooks.Put("/", updateNotebook)
	notebooks.Delete("/", deleteNotebook)

	notes := app.Group("/notes")
	notes.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))

	notes.Post("/", getNote)
	notes.Get("/all", getAllNotes)
	notes.Post("/create", createNote)
	notes.Put("/", updateNote)
	notes.Delete("/", deleteNote)

	public := app.Group("/public")
	public.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "public"})
	})
	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
        panic(err)
    }
}
