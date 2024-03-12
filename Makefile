setup-dev:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest

pre-build:
	templ generate

live-run:
	air

GCR=asia-southeast1-docker.pkg.dev/personal-232705/agung-docker-repo
TAG=latest
PLATFORM=linux/amd64
docker-build:
	docker build --platform $(PLATFORM) -t golang-htmx:$(TAG) -t $(GCR)/golang-htmx:$(TAG) . 

gcr-push:
	docker push $(GCR)/golang-htmx:$(TAG)

deploy-cloud-run:
	gcloud run deploy golang-htmx --image $(GCR)/golang-htmx:$(TAG) --region asia-southeast1 --port=5050
