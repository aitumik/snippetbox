package sqlite

import (
	"database/sql"
	"github.com/aitumik/snippetbox/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name,email,password string) error {
	HasedPassword, err := bcrypt.GenerateFromPassword([]byte(password),12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users(name,email,hashed_password,created) VALUES(?,?,?,TIME())`
	// now you have the has
	_, err = m.DB.Exec(stmt, name, email, string(HasedPassword))
	if err != nil {
		return err
	}
	return err
}

func (m *UserModel) Authenticate(email,password string) (int,error) {
	return 0,nil
}

func (m *UserModel) Get(id int) (*models.User,error) {
	return nil,nil
}




