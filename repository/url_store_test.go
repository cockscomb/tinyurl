package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/cockscomb/tinyurl/domain/entity"
	"github.com/cockscomb/tinyurl/util"
	"github.com/ory/dockertest/v3"
	"gotest.tools/v3/assert"
	"log"
	"net/url"
	"os"
	"os/exec"
	"testing"
)

var db *dynamodb.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "amazon/dynamodb-local",
		Tag:        "latest",
		Cmd:        []string{"-jar", "DynamoDBLocal.jar", "-sharedDb", "-inMemory"},
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	if err := pool.Retry(func() error {
		command := exec.Command(
			"terraform",
			"-chdir=../terraform/local",
			"apply",
			"-var=dynamodb_endpoint=http://localhost:"+resource.GetPort("8000/tcp"),
			"-auto-approve",
		)
		return command.Run()
	}); err != nil {
		log.Fatalf("terraform apply failed: %s", err)
	}
	db = dynamodb.New(dynamodb.Options{
		Credentials: credentials.NewStaticCredentialsProvider("dummy", "dummy", ""),
	}, dynamodb.WithEndpointResolver(
		dynamodb.EndpointResolverFromURL(fmt.Sprintf("http://localhost:%s", resource.GetPort("8000/tcp"))),
	))

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestURLStore_Create(t *testing.T) {
	store := NewURLStore(&URLStoreConfig{TableName: "url"}, db)

	tinyURL := &entity.TinyURL{
		ID:  "abc",
		URL: util.Must(url.Parse("https://example.com")),
	}

	t.Run("success", func(t *testing.T) {
		err := store.Create(context.Background(), tinyURL)
		assert.NilError(t, err)
	})

	t.Run("already exists", func(t *testing.T) {
		err := store.Create(context.Background(), tinyURL)
		assert.ErrorIs(t, err, entity.ErrAlreadyExists)
	})
}

func TestURLStore_Find(t *testing.T) {
	store := NewURLStore(&URLStoreConfig{TableName: "url"}, db)

	tinyURL := &entity.TinyURL{
		ID:  "def",
		URL: util.Must(url.Parse("https://example.com")),
	}
	err := store.Create(context.Background(), tinyURL)
	assert.NilError(t, err)

	t.Run("success", func(t *testing.T) {
		result, err := store.Find(context.Background(), "def")
		assert.NilError(t, err)
		assert.DeepEqual(t, result, tinyURL)
	})

	t.Run("not found", func(t *testing.T) {
		result, err := store.Find(context.Background(), "not_exist")
		assert.ErrorIs(t, err, entity.ErrNotFound)
		assert.Assert(t, result == nil)
	})
}
