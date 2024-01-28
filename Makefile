VERSION=#{RELEASE_VERSION}
EXCLUDE_THIRD_PARTY=--exclude-path third_party/errors --exclude-path third_party/google --exclude-path third_party/openapi --exclude-path third_party/validate

setup:
	go mod vendor
	go install github.com/cespare/reflex@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: api
api:
	buf generate ${EXCLUDE_THIRD_PARTY} --path api/v1

build:
	go build -ldflags "-X main.Version=${VERSION}" -v -o bin/app-api cmd/app-api/*.go

start-local:
	make api
	reflex -r "\.(go|yaml)" -s -- sh -c "make build && ./bin/app-api -config=./files/config/local.yaml"

start-prod:
	./bin/app-api -config=./files/config/prod.yaml