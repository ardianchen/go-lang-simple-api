package repository

import (
	"log"

	"github.com/ardianchen/simple-api/src/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.DB
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser() ([]User, error) {
	var users []User
	db := database.InitDb()
	if err := db.Find(&users).Error; err != nil {
		log.Fatalln(err)
	}
	return users, nil
}

// // memperbaharui data user
// func (u *user) updateUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("UPDATE users SET name='%s' WHERE id=%d", u.name, u.id)
// 	_, err := db.Exec(statement)
// 	return err
// }

// // menghapus data user
// func (u *user) deleteUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.id)
// 	_, err := db.Exec(statement)
// 	return err
// }

// // menyimpan user baru
// func (u *user) createUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("INSERT INTO users(name) VALUES('%s')", u.name)
// 	_, err := db.Exec(statement)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // menampilkan banyak user
// func getUsers(db *sql.DB, start, count int) ([]user, error) {
// 	statement := fmt.Sprintf("SELECT id, name FROM users LIMIT %d OFFSET %d", count, start)
// 	rows, err := db.Query(statement)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	users := []user{}
// 	for rows.Next() {
// 		var u user
// 		if err := rows.Scan(&u.id, &u.name); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}
// 	return users, nil
// }