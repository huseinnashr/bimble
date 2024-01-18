# bimble
you know, for dating. Try it at https://bimble-backend-http.ordinarytechfolks.com (self-hosted, sometimes down)

## Information
### A. Folder Structure
- Structure according to Golang Clean Architecture https://github.com/bxcodec/go-clean-arch
- Find code in `./internal` folder, there will be 3 layer. They are repo, usecase and handler (from bottom to top). Top layer can reuse bottom, bottom can call top layer
- All interface and struct definition in `./internal/domain` folder
- API Definition is on `./api` folder
- Config is on `./config` folder

### B. Test and others
Please see documentation.pdf for the integration test criteria and result (the last section in the bottom).

Due to time constraint the following is not implemented:
- CI/CD, but you can check my implementation here https://github.com/OrdinaryTechFolks/budgetme-backend
- Unit Test. I usually use https://github.com/vektra/mockery, which can generate mock from interface using `go gen`, than you can inject the mock object into the test struct like this `usecase := New(config, mockAccountRepo)`

### C. How to Run
#### 1. Running dev env with docker
11. Download docker-desktop with docker-compose cli https://www.docker.com/products/docker-desktop/
12. Run `docker-compose up -d`

*Alternatively we can use services in the k8s directly using these steps (requires kubectl access):*

11. Run these commands on a new terminal
```
kubectl port-forward services/yb-tservers -n bimble 5433:5433 & \
kubectl port-forward services/redis-standalone -n bimble 6379:6379 & \
kubectl port-forward services/otel-collector-collector -n bimble 4317:4317
```
12. Keep it running

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

#### 4. Modifying HTTP Server Endpoint
41. Stop the go service
42. Modify protos in `./api/v1` folder
43. Run `make api`, this will generate stub that you can use. It also generate openapi.yaml spec in `./`
44. Start the go service
