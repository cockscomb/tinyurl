package usecase

import (
	"context"
	"github.com/cockscomb/tinyurl/domain/entity"
	mock_repository "github.com/cockscomb/tinyurl/usecase/mocks"
	"github.com/cockscomb/tinyurl/util"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
	"net/url"
	"testing"
)

//go:generate mockgen -destination ./mocks/repository.go -package mock_repository . URLStore

func TestTinyURLUsecase_Generate(t *testing.T) {
	targetURL := util.Must(url.Parse("https://example.com"))

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock_repository.NewMockURLStore(ctrl)

		usecase := NewTinyURLUsecase(store)

		store.EXPECT().
			Create(gomock.Any(), gomock.AssignableToTypeOf(new(entity.TinyURL))).
			Return(nil).
			Times(1)

		tu, err := usecase.Generate(context.Background(), targetURL)
		assert.NilError(t, err)
		assert.Assert(t, tu.ID != "")
	})

	t.Run("retry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock_repository.NewMockURLStore(ctrl)

		usecase := NewTinyURLUsecase(store)

		store.EXPECT().
			Create(gomock.Any(), gomock.AssignableToTypeOf(new(entity.TinyURL))).
			Return(entity.ErrAlreadyExists).
			Times(1)
		store.EXPECT().
			Create(gomock.Any(), gomock.AssignableToTypeOf(new(entity.TinyURL))).
			Return(nil).
			Times(1)

		tu, err := usecase.Generate(context.Background(), targetURL)
		assert.NilError(t, err)
		assert.Assert(t, tu.ID != "")
	})

	t.Run("invalid scheme", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock_repository.NewMockURLStore(ctrl)

		usecase := NewTinyURLUsecase(store)

		tu, err := usecase.Generate(context.Background(), &url.URL{Scheme: "ftp"})
		assert.ErrorType(t, err, new(entity.ValidationError))
		assert.ErrorContains(t, err, "invalid scheme")
		assert.Assert(t, tu == nil)
	})
}

func TestTinyURLUsecase_Access(t *testing.T) {
	targetURL := util.Must(url.Parse("https://example.com"))
	tinyURL := entity.GenerateTinyURL(targetURL)

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock_repository.NewMockURLStore(ctrl)

		usecase := NewTinyURLUsecase(store)

		store.EXPECT().
			Find(gomock.Any(), tinyURL.ID).
			Return(tinyURL, nil).
			Times(1)

		tu, err := usecase.Access(context.Background(), tinyURL.ID)
		assert.NilError(t, err)
		assert.DeepEqual(t, tu, tinyURL)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock_repository.NewMockURLStore(ctrl)

		usecase := NewTinyURLUsecase(store)

		store.EXPECT().
			Find(gomock.Any(), "not_exist").
			Return(nil, entity.ErrNotFound).
			Times(1)

		_, err := usecase.Access(context.Background(), "not_exist")
		assert.ErrorIs(t, err, entity.ErrNotFound)
	})
}
