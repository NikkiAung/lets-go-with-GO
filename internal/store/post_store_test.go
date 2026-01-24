package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func setupTestDB(t *testing.T) *sql.DB {
	connectionString := "host=localhost user=user password=pass dbname=postgres port=5434 sslmode=disable"

	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		t.Fatalf("error on db connection %v", err)
	}

	err = Migrate(db, "../../migrations/")

	if err != nil {
		t.Fatalf("error on db migration %v", err)
	}

	return db
}

func TestCreatePost(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresPostStore(db)

	tests := []struct {
		name string
		post *Post 
		wantErr bool
	} {
		{
			name : "valid post",
			post : &Post {
				Title : "Testing post",
				Content: "Testing post content",
			},
			wantErr: false, // <- Expect success 
		},
		{
			name: "empty title should fail",
			post: &Post{Title: "", Content: "World"},
			wantErr: true,   // <- Expect error  
		}
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			createdPost, err := store.CreatePost(tt.post)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t,err)
			assert.Equal(t, tt.post.Title, createdPost.Title)
			assert.Equal(t, tt.post.Content, createdPost.Content)

			post, err := store.GetPostById(createdPost.ID)

			require.NoError(t, err)
			
			assert.Equal(t, createdPost.ID, post.ID)
			assert.Equal(t, createdPost.Title, post.Title)
			assert.Equal(t, createdPost.Content, post.Content)
		})
	}
}