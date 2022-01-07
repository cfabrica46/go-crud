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
		rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
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

func GetIDByUsername(username string) (id int, err error) {
	row := db.QueryRow("SELECT users.id FROM users WHERE users.username = $1", username)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
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

func InsertUser(username, password, email string) (err error) {
	stmt, err := db.Prepare("INSERT INTO users(username,password,email) VALUES ($1,$2,$3)")
	if err != nil {
		return
	}

	_, err = stmt.Exec(username, password, email)
	if err != nil {
		return
	}
	return
}

func DeleteUserbByID(id int) (count int, err error) {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return
	}

	r, _ := stmt.Exec(id)
	countAux, _ := r.RowsAffected()
	count = int(countAux)
	return
}

func DeleteUserbByUsername(username string) (count int, err error) {
	stmt, err := db.Prepare("DELETE FROM users WHERE username = $1")
	if err != nil {
		return
	}

	r, _ := stmt.Exec(username)
	countAux, _ := r.RowsAffected()
	count = int(countAux)
	return
}
