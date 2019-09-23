# cloudblog
Over engineered CQRS blog system, with NoSQL caches and message queues

In order to properly understand various architectural principles I have created this blog.

The over all system is built along CQRS lines, with Redis caching blog entries, Postgres being the source of truth.

The code is laid out using Clean architecture as a foundation. 

SOLID principles have been used to determine integration, the business logic is NOT dependant upon the repositories, they are dependant upon it via interfaces.


```
.
├── LICENSE
├── readarticle
│   ├── builder.sh
│   ├── cmd
│   │   └── main.go
│   ├── Dockerfile
│   ├── endtoend_test.sh
│   ├── go.mod
│   ├── go.sum
│   ├── integration
│   │   ├── grpc
│   │   │   ├── client
│   │   │   │   └── v1
│   │   │   │       └── readarticle-client.go
│   │   │   ├── gen
│   │   │   │   └── v1
│   │   │   │       └── article.pb.go
│   │   │   └── proto
│   │   │       └── v1
│   │   │           └── article.proto
│   │   ├── messagequeue
│   │   │   └── v1
│   │   │       └── refill-cache.go
│   │   └── repository
│   │       └── cache
│   │           └── v1
│   │               └── cache.go
│   ├── internal
│   │   ├── pkg
│   │   │   └── articleservice
│   │   │       ├── articleservice.go
│   │   │       └── articleservice_test.go
│   │   └── repository
│   │       └── redis
│   │           ├── actions.go
│   │           ├── actions_test.go
│   │           └── connection.go
│   └── readarticle.sh
├── README.md
└── REST
    ├── cmd
    │   ├── main.go
    │   └── routes.go
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    └── Makefile
```
