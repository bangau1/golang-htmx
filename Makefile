setup-dev:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest

pre-build:
	templ generate

live-run:
	air

TAG=latest
docker-build:
	docker build -t golang-htmx:$(TAG) .