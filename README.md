# SumMetricService
Metric Test

* Author(s): Anthony Mays
* Current Version: 0.0.1
* Release Date: August 28, 2021
* MIT License
___
## Getting Started

Follow the instructions below to get the Go RESTful API up and running on your Linux Environment

### Prerequisites
* An Ubuntu 18+ or CentOS 7+ Operating System
* Hopefully some patience with my stupidity LOL

### Setup

1. Clone the git repo:
```bash
$ git clone https://github.com/tonymays/SumMetricService.git
$ cd SumMetric
```

2. Obtain the package dependencies

```bash
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/handlers"
$ go get "github.com/gofrs/uuid"
```

3. Examine conf.json file specified:

* CountStrategy 	- the number of minutes of elapsed time to filter out a metric
* ClearOnSum 		- off|on to clear outdated metrics on a summation operation
* HTTPS				- off|on to use HTTPS or not IGNORE for the this API
* Cert       		- For HTTPS Certs ignore here
* Key        		- FOR HTTPS Certs ignore here
* ServerListenPort	- which port the server will run on

___
## Running the API


