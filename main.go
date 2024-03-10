package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/bangau1/golang-htmx/film"
	"github.com/bangau1/golang-htmx/view"
	"github.com/oklog/ulid/v2"
)

type Controller struct {
	router      *http.ServeMux
	filmService film.FilmService
}

func (c *Controller) homePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	films, err := c.filmService.GetFilms(ctx)
	if err != nil {
		view.Error(fmt.Sprintf("%v", err)).Render(ctx, w)
		return
	}

	view.Index(films).Render(ctx, w)
}

func (c *Controller) getFilm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filmId := r.PathValue("id")
	film, err := c.filmService.GetFilm(ctx, filmId)
	if err != nil {
		view.Error(fmt.Sprintf("%v", err)).Render(ctx, w)
		return
	}

	view.FilmDetail(film).Render(ctx, w)
}

func (c *Controller) deleteFilm(w http.ResponseWriter, r *http.Request) {
	filmId := r.PathValue("id")
	err := c.filmService.DeleteFilm(r.Context(), filmId)
	if errors.Is(err, film.ErrNotFound) {
		w.WriteHeader(404)
		return
	}
	return
}

func main() {
	c := Controller{
		router:      http.NewServeMux(),
		filmService: film.NewInMemoryFilmService(movies...),
	}

	addr := ":5050"

	// serve static assets file a simple fileserver
	assetsFs := http.FileServer(http.Dir("./assets"))
	c.router.Handle("GET /assets/", http.StripPrefix("/assets/", assetsFs))

	// then the rest are related with our pages
	c.router.HandleFunc("GET /", c.homePage)
	c.router.HandleFunc("GET /films/{id}", c.getFilm)
	c.router.HandleFunc("DELETE /films/{id}", c.deleteFilm)

	log.Println("starting server at " + addr)

	err := http.ListenAndServe(addr, c.router)
	if err != nil {
		fmt.Println(err)
	}

}

func init() {
	for i := 0; i < len(movies); i++ {
		movies[i].Id = ulid.Make().String()
	}
}

var (
	movies = []film.Film{
		{Title: "The Shawshank Redemption", Director: "Frank Darabont", PosterURL: "/assets/01.jpeg"},
		{Title: "The Godfather", Director: "Francis Ford Coppola", PosterURL: "/assets/02.jpeg"},
		{Title: "The Dark Knight", Director: "Christopher Nolan", PosterURL: "/assets/03.jpeg"},
		{Title: "Pulp Fiction", Director: "Quentin Tarantino", PosterURL: "/assets/04.jpeg"},
		{Title: "Schindler's List", Director: "Steven Spielberg", PosterURL: "/assets/05.jpeg"},
		{Title: "12 Angry Men", Director: "Sidney Lumet", PosterURL: "/assets/06.jpeg"},
		{Title: "The Lord of the Rings: The Return of the King", Director: "Peter Jackson", PosterURL: "/assets/07.jpeg"},
		{Title: "Fight Club", Director: "David Fincher", PosterURL: "/assets/08.jpeg"},
		{Title: "Parasite", Director: "Bong Joon-ho", PosterURL: "/assets/09.jpeg"},
		{Title: "Inception", Director: "Christopher Nolan", PosterURL: "/assets/10.jpeg"},
	}
)
