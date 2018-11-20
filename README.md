## Assignment description
Assume that there is a service that keeps track of latencies of the API calls it made to other services. Write a service that has two endpoints exposed enabling the latency tracker to POST the latency of every API and also enable the user to GET the top n latencies where n is defined at the start of the server. The user is only concerned about the top n latencies.

## APIs
#### Add latency data

`curl -X POST http://localhost:8081/api/v1/latency -d '{"latency": 1936585}'`

*  201 on success
*  422 (UnprocessableEntity) on faulty object

#### Get top n latency data

`curl -X GET http://localhost:8081/api/v1/latency`

*  200 on success
*  500 otherwise
*  `{"latencies":[]}` as response

## Run

```
go get github.com/Vigneshsekar/exotel-test 
cd <path>/exotel-test
go run main.go
```
