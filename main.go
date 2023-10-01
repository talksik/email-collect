package main

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/gofiber/fiber"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type users struct {
	mu sync.Mutex
	db *sql.DB
}

const file string = "users.db"
const createTable string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY autoincrement,
  time DATETIME NOT NULL,
  email TEXT
  );`

func initializeDb() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		logrus.Fatal(err)
	}

	if _, err := db.Exec(createTable); err != nil {
		logrus.Fatal(err)
	}
	return db
}

func main() {
	app := fiber.New()

	db := initializeDb()

	app.Post("/email", func(c *fiber.Ctx) {
		email := c.FormValue("email")
		if email == "" {
			c.Send("Please provide an email address")
			c.Status(400)
			c.Redirect("https://flowy.live?error=true")
      logrus.Error("No email provided")
			return
		}

		logrus.Info(fmt.Sprintf("Received email: %s", email))

		db.Exec("INSERT INTO users (time, email) VALUES (datetime('now'), ?)", email)
		c.Send("Welcome to the newsletter!")
		c.Redirect("https://flowy.live?success=true")
		return
	})

	logrus.Fatal(app.Listen(":3000"))
}
