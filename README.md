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

* CountStrategy 	- the number of minutes of elapsed time to filter out a metric (default is 60 minutes)
* ClearOnSum 		- off|on to clear outdated metrics on a summation operation
* InitWithTestData	- off|on if on the system will generate the following:
```bash
- 1 record 3 hours old with active_visitors at 15
- 1 record 2 hours old with active_vistiors at 10
- 1 record 30 minutes old with active_vistors at 5
- 1 record 15 minutes old with active_vistors at 20
```
* HTTPS				- off|on to use HTTPS or not IGNORE for the this API
* Cert       		- For HTTPS Certs ignore here
* Key        		- FOR HTTPS Certs ignore here
* ServerListenPort	- which port the server will run on

The service must be restarted for config changes to take affect.
This service only uses in-memory data cache which means data posted via the routes below are lost when the service is stopped

4. Compile the service
```bash
$ go build ./src/service
```

5. Start the service
```bash
$ ./service
```
___

This will respond with
Listening on port :8080

## Running the API
1. Open the conf.json file an set InitWithTestData to on
```bash
Example:
{
	"CountStrategy": 60,
	"ClearOnSum": "off",
	"InitWithTestData": "on",
	"HTTPS": "off",
	"Cert": "/etc/ssl/certs/cert.pem",
	"Key": "/etc/ssl/certs/key.pem",
	"ServerListenPort": ":8080"
}
```
2.
