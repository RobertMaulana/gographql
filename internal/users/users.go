package users

import (
	"database/sql"
	"github.com/glyphack/graphlq-golang/internal/pkg/db/postgre"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := database.Client.Prepare("INSERT INTO users(username,password) VALUES($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Authenticate() bool {
	statement, err := database.Client.Prepare("select password from users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Client.Prepare("select id from users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

//GetUserByID check if a user exists in database and return the user object.
func GetUsernameById(userId int) (User, error) {
	statement, err := database.Client.Prepare("select username from users WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(userId)

	var username string
	err = row.Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return User{}, err
	}

	return User{ID: userId, Username: username}, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
