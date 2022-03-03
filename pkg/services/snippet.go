package services

import (
	"github.com/aitumik/snippetbox/pkg/models"
	"github.com/elastic/go-elasticsearch"
)

type SnippetService struct {
	*elasticsearch.Client
}

func (s *SnippetService) CreateSnippet(snippet models.Snippet) error {
	return nil
}

func (s *SnippetService) ReadSnippet(id int) (*models.Snippet, error) {
	return nil, nil
}

func (s *SnippetService) CreateManySnippets(sn []*models.Snippet) error {
	return nil
}

func (s *SnippetService) SearchSnippet(query string) error {
	return nil
}

func NewSnippetService() *SnippetService {
	return &SnippetService{}
}
