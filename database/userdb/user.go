package userdb

import (
	"database/sql"

	"github.com/cfabrica46/go-crud/structure"
)

func GetAllUsers() (users []structure.User, err error) {
	rows, err := db.Query("SELECT users.id,users.username,users.email FROM users")
	if err != nil {
		return
	}

	for rows.Next() {
		var userBeta structure.User
		err = rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
		if err != nil {
			return
		}
		users = append(users, userBeta)
	}

	return
}

func GetUserByID(id int) (user *structure.User, err error) {
	row := db.QueryRow("SELECT users.id,users.username,users.password,users.email FROM users WHERE users.id = $1", id)

	var userBeta structure.User
	err = row.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Password, &userBeta.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	user = &userBeta
	return
}

/* func GetUserByUsername(username string) (user *structure.User, err error) {
	var userBeta structure.User

	row := db.QueryRow("SELECT users.id,users.username,users.password,users.email FROM users WHERE users.username = $1", username)
	err = row.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Password, &userBeta.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	user = &userBeta
	return
} */

func GetUserByUsernameAndPassword(username, password string) (user *structure.User, err error) {
	row := db.QueryRow("SELECT users.id, users.email FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	var userBeta structure.User

	err = row.Scan(&userBeta.ID, &userBeta.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	user = &userBeta
	user.Username = username
	user.Password = password
	return
}

func CheckIfUserAlreadyExist(username string) (check bool, err error) {
	row := db.QueryRow("SELECT users.id FROM users WHERE users.username = $1", username)

	var user structure.User

	err = row.Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}

	check = true
	return
}

func InsertUser(username, password, email string) (id int, err error) {
	stmt, err := db.Prepare("INSERT INTO users(username,password,email) VALUES ($1,$2,$3)")
	if err != nil {
		return
	}

	r, err := stmt.Exec(username, password, email)
	if err != nil {
		return
	}

	id64, _ := r.LastInsertId()
	id = int(id64)
	return
}

func DeleteUserbByID(id int) (err error) {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return
	}
	return
}
