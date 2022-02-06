package elasticsearch

import (
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/elastic/go-elasticsearch"
)

type UserModel struct {
	*elasticsearch.Client
}

func (m *UserModel) Insert(name,email,password string) error {
	return nil
}

func (m *UserModel) Get(id int) (*models.User,error) {
	return nil,nil
}

func (m *UserModel) Authenticate(email,password string) (int,error) {
	return 0,nil
}
