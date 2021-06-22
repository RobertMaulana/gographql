package links

import (
	"github.com/RobertMaulana/graphql-go/internal/pkg/db/model/users"
	database "github.com/RobertMaulana/graphql-go/internal/pkg/db/postgre"
	"log"
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
	statement, err := database.Client.Prepare("INSERT INTO links(title,address, user_id) VALUES($1,$2, $3) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()
	//#5
	var lastInsertId int
	err = statement.QueryRow(link.Title, link.Address, link.User.ID).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Row inserted!")
	return int64(lastInsertId)
}

func GetAll() []Link {
	stmt, err := database.Client.Prepare("select L.id, L.title, L.address, L.user_id, U.username from links L inner join users U on L.user_id = U.id")
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
	var id int
	var username string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
