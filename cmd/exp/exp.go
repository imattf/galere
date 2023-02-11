// Fun with postgres...

package main

import (
	"fmt"

	"github.com/imattf/galere/models"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

type Order struct {
	ID          int
	UserID      int
	Amount      int
	Description string
}

func main() {
	// cfg := PostgresConfig{
	// 	Host:     "localhost",
	// 	Port:     "5432",
	// 	User:     "baloo",
	// 	Password: "junglebook",
	// 	Database: "lenslocked",
	// 	SSLMode:  "disable",
	// }

	// db, err := sql.Open("pgx", "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable")

	// db, err := sql.Open("pgx", cfg.String())

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("bob4@bob.com", "bob123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// Create Table...
	// _, err = db.Exec(`
	//   CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT NOT NULL
	//   );

	//   CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	//   );
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables checked/created.")

	// name := "Jon Calhoun"
	// email := "jon@calhoun.io"
	// _, err = db.Exec(`
	// INSERT INTO users(name, email)
	// VALUES($1, $2);`, name, email)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created.")

	// name := "Bob Aol"
	// email := "bob@aol.com"
	// row := db.QueryRow(`
	//   INSERT INTO users(name, email)
	//   VALUES($1, $2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id=", id)

	// id := 22 //id that does not exist
	// row := db.QueryRow(`
	//   SELECT name, email
	//   FROM users
	//   WHERE id=$1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("No records found!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User info: name=%s email=%s\n", name, email)

	// userID := 2
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 	  INSERT INTO orders(user_id, amount, description)
	// 	  VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	//TODO Uncomment from here down...
	// 	var orders []Order

	// 	// read all the records for a user into rows
	// 	userID := 2
	// 	rows, err := db.Query(`
	// 	  SELECT id, amount, description
	// 	  FROM orders
	// 	  WHERE user_id=$1`, userID)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer rows.Close()

	// 	// add all the read records into a order struct & print
	// 	for rows.Next() {
	// 		var order Order
	// 		order.UserID = userID
	// 		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		orders = append(orders, order)
	// 	}
	// 	err = rows.Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Orders:", orders)
}
