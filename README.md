# bimble
you know, for dating

### How to Run
#### 1. Running dev env with docker
11. Download docker-desktop with docker-compose cli https://www.docker.com/products/docker-desktop/
12. Run `docker-compose up -d`

#### 2. Migrate DB with Goose
21. Set these env
```
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/postgres"
```
22. Run `goose -dir migrations up`

#### 3. Running go service
31. Install make cli https://formulae.brew.sh/formula/make
32. Run `make setup`
33. Run `make start-dev`