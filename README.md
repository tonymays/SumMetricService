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
```
WARNING: if this setting is turned on the Summation Route, defined below, will as it's first
step remove outdated metric data for the specified metric
```
* InitWithTestData	- off|on if on the system will generate the following:
```
	- 1 record 3 hours old with active_visitors at 15
	- 1 record 2 hours old with active_vistiors at 10
	- 1 record 30 minutes old with active_vistors at 5
	- 1 record 15 minutes old with active_vistors at 20
```
* HTTPS				- off|on, on to use HTTPS or off to HTTPS
* Cert       		- Cert File
* Key        		- Cert Key
```
	* Warning: This are placed in the config file for future effect and are ignored
	         by this Service
```
* ServerListenPort	- which port the server will run on

```
* Warning: The service must be restarted for config changes to take affect.
* This service only uses in-memory data cache which means data posted via the routes
  below are lost when the service is stopped
```

4. Compile the service
```bash
$ go build ./src/service
```
___

This will respond with
Listening on port :8080

## Running the API
1. Open the conf.json file an set InitWithTestData to on to obtain test data described above
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
2. Start the service
```bash
$ ./service
```

## API Route Guide

#### API List
```
* GET - /metric/active_visitors (returns all metrics for a specified key)
* GET - /metric/active_visitors/active (returns all non-outdated metrics for a specified key)
* GET - /metric/active_visitors/sum (returns the sum of metric for a specified key)
* POST - /metric/active_visitors (adds a new metric for a specified key)
* DELETE - /metric/active_visitors/clear (clears outdated metrics for a specified key)
```

### I) Getting and Summing Metrics
___
#### 1. Get All Metrics Entered
* GET - /metric/active_visitors

##### Request

* Headers

```
{
  Content-Type: application/json
}
```

##### Example Response
* Body
```
[
    {
        "id": "ba63e7de-b204-41d3-9ca9-523dc23c3bce",
        "key": "active_vistors",
        "value": 15,
        "entry_time": "2021-08-28 22:17:34.418686251 +0000 UTC m=-10799.998222152"
    },
    {
        "id": "0f8da723-7fd5-4576-a7ee-4f158dc67e16",
        "key": "active_vistors",
        "value": 10,
        "entry_time": "2021-08-28 23:17:34.418686251 +0000 UTC m=-7199.998222152"
    },
    {
        "id": "a2493c7f-6332-4f65-8695-030009b997f2",
        "key": "active_vistors",
        "value": 5,
        "entry_time": "2021-08-29 00:47:34.418686251 +0000 UTC m=-1799.998222152"
    },
    {
        "id": "19e374bb-9220-4cb6-aedb-2c9b05d16dd6",
        "key": "active_vistors",
        "value": 20,
        "entry_time": "2021-08-29 01:02:34.418686251 +0000 UTC m=-899.998222152"
    }
]
```
***
#### 2. Get All Active Metrics Entered
Active Metrics are those metrics that can be sum via the Config CountStrategy setting

* GET - /metric/active_visitors/active

##### Request
* Headers
```
{
  Content-Type: application/json
}
```

##### Example Response
* Body
```
[
    {
        "id": "a2493c7f-6332-4f65-8695-030009b997f2",
        "key": "active_vistors",
        "value": 5,
        "entry_time": "2021-08-29 00:47:34.418686251 +0000 UTC m=-1799.998222152"
    },
    {
        "id": "19e374bb-9220-4cb6-aedb-2c9b05d16dd6",
        "key": "active_vistors",
        "value": 20,
        "entry_time": "2021-08-29 01:02:34.418686251 +0000 UTC m=-899.998222152"
    }
]
```
***
#### 3. Sum Metric against CountStrategy
* GET - /metric/active_visitors/sum

##### Request
* Headers

```
{
  Content-Type: application/json
}
```

##### Example Response
* Body
```
{
    "key": "active_vistors",
    "sum": 25
}
```

### II) Add and Clear Key Metrics

___
#### 1. Add a New Metric for a specified key
* POST - /metric/active_visitors

##### Request

* Headers

```
{
  Content-Type: application/json
}
```

*Body

```
{
    "value":6
}
```

##### Example Response
* Body
```
{
    "key": "active_vistors",
    "value": 6,
    "entry_time": "2021-08-28 23:58:11.115465451 +0000 UTC"
}
```

#### 2. Clear Outdated Metrics for a specified key
* DELETE - /metric/active_visitors/clear

##### Request

* Headers

```
{
  Content-Type: application/json
}
```


##### Example Response
* Header
```
Status: 200Ok
```
