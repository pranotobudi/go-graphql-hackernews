package links

import (
	"fmt"
	"log"

	"github.com/pranotobudi/go-graphql-hackernews/database"
	"github.com/pranotobudi/go-graphql-hackernews/internal/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

//#2
func (link Link) Save() int64 {
	//#3
	postgres := database.InitDB()
	// 	query := fmt.Sprintf(`
	// 	INSERT INTO Users(Username,Password)
	// 	VALUES
	// 	('%s', '%s')
	// 	RETURNING ID;
	// `, user.Username, hashedPassword)
	fmt.Println("title: ", link.Title, "address: ", link.Address)
	query := fmt.Sprintf(`
		INSERT INTO links( title,address)
		VALUES
		('%s', '%s')
		RETURNING ID;
	`, link.Title, link.Address)
	log.Println("query: ", query)

	// res, err := postgres.DB.ExecContext(context.Background(), query)
	var id int64
	err := postgres.DB.QueryRow(query).Scan(&id)
	if err != nil {
		log.Println("SQL statement execution failed: ", err)
		return -1
	}
	log.Println("data inserted")

	// id, _ := res.LastInsertId()
	return id
}

func GetAll() []Link {
	postgres := database.InitDB()
	stmt, err := postgres.DB.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username) // changed
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
