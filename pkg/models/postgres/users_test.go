package postgres

import (
	"reflect"
	"testing"
	"time"

	"github.com/aitumik/snippetbox/pkg/models"
)

func TestUserModelGet(t *testing.T) {
	if t.Short() {
		t.Skip("postgres: skipping integration tests")
	}

	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Indiana Jones",
				Email:   "indiana@example.com",
				Created: time.Date(2022, 2, 21, 8, 45, 0, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "Non Existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		// Create the runner
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			// Create new instance of the UserModel
			m := UserModel{db}

			user, err := m.Get(tt.userID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}