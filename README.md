# Go example REST API layout

Routing, middleware, logging, mocks, asserts, error stacktrace, database driver.

This is an example CRUD API that showcases a simple server structured using clean code architecture.
It's a CRUD about articles, that a author can write, set a title and a content.

## How to run

```
go run main.go
```

## Clean code architecture

A dependency diagram of clean code arch is:

```
    Entity
        ^
        |
    Usecase
    ^       ^
    |       |
Transport   Store
```

Entity is where all the models (structs) live. These structs can be shared across layers and it's how the layers
communicate data.
Usecase is where all the business logic is embedded.
Transport handles everything related with the request and translating into usecase terms. For example, all HTTP requests are transformed into usecase methods and entities structs.
Store handles everything related to storing and accessing the data. It's the data access layer. For example, if 
we're using a SQL database, it translates the CRUD methods into sql queries.

Each layer defines an interface on how it wants the other layers to accept and return information. This interface
is implemented by each layer.

## Middlewares

I've added some middlewares usually useful in a server. One adds a request id to the requests. 
Another one is a request timer that logs the request duration. And other logs the method and path of arriving requests.

## Mocks

I've added a unit test in the usecase layer to show how the interfaces are mocked.
If the interfaces are modified, the following command needs to be run to update the mocks.

```
go generate ./...
```

