package users

import (
	"fmt"
	"log"

	"github.com/pranotobudi/go-graphql-hackernews/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

func (user *User) Create() {
	// statement, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	// print(statement)

	hashedPassword, err := HashPassword(user.Password)
	postgres := database.InitDB()
	query := fmt.Sprintf(`
		INSERT INTO Users(Username,Password)
		VALUES
		('%s', '%s')
		RETURNING ID;
	`, user.Username, hashedPassword)
	log.Println("query: ", query)

	// res, err := postgres.DB.ExecContext(context.Background(), query)
	var id int64
	err = postgres.DB.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

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

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	// statement, err := postgres.DB.Prepare("select ID from Users WHERE Username = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// row := statement.QueryRow(username)
	log.Println("GetUserIdByUsername")
	postgres := database.InitDB()
	query := fmt.Sprintf(`
	SELECT ID from users WHERE username = '%s';
	`, username)
	log.Println("query: ", query)
	var id int
	err := postgres.DB.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (user *User) Authenticate() bool {
	// statement, err := database.Db.Prepare("select Password from Users WHERE Username = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// row := statement.QueryRow(user.Username)
	log.Println("Authenticate")
	postgres := database.InitDB()
	query := fmt.Sprintf(`
	select Password from Users WHERE Username = '%s';
	`, user.Username)
	log.Println("query: ", query)
	var hashedPassword string
	err := postgres.DB.QueryRow(query).Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	// err = row.Scan(&hashedPassword)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return false
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

	return CheckPasswordHash(user.Password, hashedPassword)
}
