package asset_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
)

func TestNewRepositoryUnknown(t *testing.T) {
	repository, err := asset.NewRepository("unknown", "")
	if err == nil {
		t.Fatalf("unknown repository: %T", repository)
	}
}

func TestRegisterRepositoryBuilder(t *testing.T) {
	builderName := "testbuilder"

	repository, err := asset.NewRepository(builderName, "")
	if err == nil {
		t.Fatalf("testbuilder is: %T", repository)
	}

	asset.RegisterRepositoryBuilder(builderName, func(_ string) (asset.Repository, error) {
		return asset.NewInMemoryRepository(), nil
	})

	repository, err = asset.NewRepository(builderName, "")
	if err != nil {
		t.Fatalf("testbuilder is not found: %v", err)
	}

	_, ok := repository.(*asset.InMemoryRepository)
	if !ok {
		t.Fatalf("testbuilder is: %T", repository)
	}
}

func TestNewRepositoryMemory(t *testing.T) {
	repository, err := asset.NewRepository(asset.InMemoryRepositoryBuilderName, "")
	if err != nil {
		t.Fatal(err)
	}

	_, ok := repository.(*asset.InMemoryRepository)
	if !ok {
		t.Fatalf("repository not correct type: %T", repository)
	}
}

func TestNewRepositoryFileSystem(t *testing.T) {
	repository, err := asset.NewRepository(asset.FileSystemRepositoryBuilderName, "testdata")
	if err != nil {
		t.Fatal(err)
	}

	_, ok := repository.(*asset.FileSystemRepository)
	if !ok {
		t.Fatalf("repository not correct type: %T", repository)
	}
}

func TestNewTiingoRepository(t *testing.T) {
	repository, err := asset.NewRepository(asset.TiingoRepositoryBuilderName, "1234")
	if err != nil {
		t.Fatal(err)
	}

	_, ok := repository.(*asset.TiingoRepository)
	if !ok {
		t.Fatalf("repository not correct type: %T", repository)
	}
}
