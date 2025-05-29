package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Player struct {
	Username  string
	Initials  string
	CreatedAt time.Time
}

func GetUserByID(db *sql.DB, user_id int32) error {
	var newPlayer Player
	dbQuery := "SELECT username, initials, created_at FROM users WHERE user_id = $1"
	dbRows := db.QueryRow(dbQuery, user_id)

	// .Scan needs pointers to write into the memory instead of copies
	err := dbRows.Scan(&newPlayer.Username, &newPlayer.Initials, &newPlayer.CreatedAt)
	if err != nil {
		return fmt.Errorf("Couldn't scan database: %s", err.Error())
	}

	fmt.Println(newPlayer)
	return nil
}

func SearchUsersByField(db *sql.DB, searchField string, query string) ([]string, error) {
	var sqlQuery string
	var usernames []string

	switch searchField {
	case "username":
		sqlQuery = `SELECT username FROM users 
                    WHERE username ILIKE '%' || $1 || '%' 
                    LIMIT 50`
	case "initials":
		sqlQuery = `SELECT username FROM users 
                    WHERE initials ILIKE '%' || $1 || '%' 
                    LIMIT 50`
	default:
		return nil, fmt.Errorf("invalid search field")
	}

	rows, err := db.Query(sqlQuery, query)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		usernames = append(usernames, username)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return usernames, nil
}
