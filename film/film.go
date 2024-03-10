package film

import (
	"context"
	"errors"
)

type Film struct {
	Id        string
	Title     string
	Director  string
	PosterURL string
}

var ErrNotFound = errors.New("not found error")

type FilmService interface {
	GetFilms(ctx context.Context) ([]Film, error)
	GetFilm(ctx context.Context, filmId string) (Film, error)
	DeleteFilm(ctx context.Context, filmId string) error
}

var _ FilmService = &InMemoryFilmService{}

type InMemoryFilmService struct {
	data map[string]Film
}

func NewInMemoryFilmService(films ...Film) *InMemoryFilmService {
	res := &InMemoryFilmService{
		data: make(map[string]Film),
	}

	for _, film := range films {
		res.data[film.Id] = film
	}
	return res
}

func (i *InMemoryFilmService) filmList() []Film {
	var res []Film
	for _, film := range i.data {
		res = append(res, film)
	}
	return res
}

func (i *InMemoryFilmService) GetFilms(ctx context.Context) ([]Film, error) {
	return i.filmList(), nil
}

func (i *InMemoryFilmService) GetFilm(ctx context.Context, filmId string) (Film, error) {
	_, found := i.data[filmId]
	if !found {
		return Film{}, ErrNotFound
	}

	return i.data[filmId], nil
}

func (i *InMemoryFilmService) DeleteFilm(ctx context.Context, filmId string) error {
	_, found := i.data[filmId]
	if !found {
		return ErrNotFound
	}
	delete(i.data, filmId)
	return nil
}
