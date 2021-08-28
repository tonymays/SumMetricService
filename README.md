# SumMetricService
Metric Test

* Author(s): Anthony Mays
* Current Version: 0.0.1
* Release Date: August 27, 2021
* MIT License
___
## Getting Started

Obtain package dependencies

```bash
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/handlers"
$ go get "github.com/gofrs/uuid"
```

To build the service
cd to the cloned directory

```bash
$ go build ./src/service
```

Do the following to execute the MicroService
```bash
$ ./service
```

Do the following to execute the MicroService as a Service

```bash
$ go build ./src/rest_api/cmd/app
$ sudo systemctl start restapi
$ sudo systemctl enable restapi
```

