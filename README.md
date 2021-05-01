# Fuzzy Search API Engine

This repository is split to two components, `go-app` - a gokit inspired application, and `elasticsearch`

This [document](https://logz.io/blog/elasticsearch-api/) is really good to get started with integrating the API
### go-app

For this directory, the project layout is inspired from [here](https://github.com/golang-standards/project-layout)

This version of go is `1.16`

## Running the application

For the first time, copy the following files (and change the values if necessary)
```
go-app/.env.sample ----> go-app/.env
```

Run the following command to start the API and elasticsearch servers
```
docker-compose up
```

To shell into the API server instance
```
docker exec -it elasticsearch-api sh
```

## Seeding elasticsearch db

Run the following to seed the `user` index with 100 distinct results for testing
```
./tools/seeds/seed.sh
```

## Documentation

View the documentation at [https://weiyuan-lane.github.io/elasticsearch-api](https://weiyuan-lane.github.io/elasticsearch-api/)

## TODO

- Redis caching
