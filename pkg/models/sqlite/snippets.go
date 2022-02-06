package sqlite

import (
	"database/sql"
	"github.com/aitumik/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert This will insert a new snippet into the database
func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get this will return specific snippet based on id
func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id,title,content,created,expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// use the query row to execute the SQL statement
	row := s.DB.QueryRow(stmt, id)

	// initialize a pointer to a new zerod struct
	m := &models.Snippet{}

	err := row.Scan(&m.ID, &m.Title, &m.Content, &m.Created, &m.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	// If everything went okay then return the snippet object
	return m, nil
}

// Latest This will return the top 10 most recently created snippets
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id,title,content,created,expires FROM snippets ORDER BY DESC LIMIT 10`

	rows,err := s.DB.Query(stmt)
	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		m := &models.Snippet{}

		err = rows.Scan(&m.ID,&m.Title,&m.Content,&m.Created,&m.Expires)
		if err != nil {
			return nil,err
		}
		snippets = append(snippets,m)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}

	return snippets,nil
}
