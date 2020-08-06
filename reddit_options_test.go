package reddit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromEnv(t *testing.T) {
	os.Setenv("GO_REDDIT_CLIENT_ID", "id1")
	defer os.Unsetenv("GO_REDDIT_CLIENT_ID")

	os.Setenv("GO_REDDIT_CLIENT_SECRET", "secret1")
	defer os.Unsetenv("GO_REDDIT_CLIENT_SECRET")

	os.Setenv("GO_REDDIT_CLIENT_USERNAME", "username1")
	defer os.Unsetenv("GO_REDDIT_CLIENT_USERNAME")

	os.Setenv("GO_REDDIT_CLIENT_PASSWORD", "password1")
	defer os.Unsetenv("GO_REDDIT_CLIENT_PASSWORD")

	c, err := NewClient(nil, FromEnv)
	assert.NoError(t, err)

	type values struct {
		id, secret, username, password string
	}

	expect := values{"id1", "secret1", "username1", "password1"}
	actual := values{c.ID, c.Secret, c.Username, c.Password}
	assert.Equal(t, expect, actual)
}

func TestWithCredentials(t *testing.T) {
	withCredentials := WithCredentials("id1", "secret1", "username1", "password1")
	c, err := NewClient(nil, withCredentials)
	assert.NoError(t, err)
	assert.Equal(t, "id1", c.ID)
	assert.Equal(t, "secret1", c.Secret)
	assert.Equal(t, "username1", c.Username)
	assert.Equal(t, "password1", c.Password)
}

func TestWithBaseURL(t *testing.T) {
	baseURL := "http://localhost:8080"
	c, err := NewClient(nil, WithBaseURL(baseURL))
	assert.NoError(t, err)
	assert.Equal(t, baseURL, c.BaseURL.String())
}

func TestWithTokenURL(t *testing.T) {
	tokenURL := "http://localhost:8080/api/v1/access_token"
	c, err := NewClient(nil, WithTokenURL(tokenURL))
	assert.NoError(t, err)
	assert.Equal(t, tokenURL, c.TokenURL.String())
}
