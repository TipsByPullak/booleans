# Implementation of booleans.io using golang
###### [Reference link](http://booleans.io)

### Description
It was a rudimentary attempt on my part to implement a RESTful service that provides a database to create, store, retrieve, update and delete booleans
I took it up as a small project to learn [gin](https://github.com/gin-gonic/gin) and [gorm](http://gorm.io/index.html). Besides, I got a hands on of docker "containerisation" as well.
#### *This project is still a work in progress. The README will be updated as things move*
#### TODO List-
1. Implement automatic creation of database in the DB server
2. Implement tests for the current code
3. Add support for OAuth/JWT, Concurrency with sync and mutex

___

### Starting the boolean service

Please make sure that you have a MySQL server running.
The server should have a DB named mysql (Will be corrected for in future commits)
#### Using a local copy of the source
1. Clone using `git clone https://github.com/TipsByPullak/booleans`
2. Change your directory to `booleans` (as cloned above) and run `go build`. (Assuming a proper setup of the pre-requisites, hoping no error pops up for you)
3. run-  
  `./booleans -ip <IP to your DB server:PORT> -usr <username for the DB> -pass <password to the DB>` (For Mac/Linux)  
  or  
  `booleans.exe -ip <IP to your DB server:PORT> -usr <username for the DB> -pass <password to the DB` (For Windows)
    - if flag `-ip` is not provided, the service assumes that the DB server is running at "127.0.0.1:3306"
    - if `-usr` flag is not provided, the service assumes a username of root (This may not be safe in critical applications)
    - if `-pass` flag ins not provided, the service assumes password to be empty.
*Note: Behaviour on Windows and Linux has not been thoroughly tested. Users are welcome to contribute for the same*
#### Using a Docker image
*Note that due to limitations and testing conveniences, the docker image only connects to a MySQL server on the localhost for now*
1. First run the following-
`docker pull tipsbypullak/booleans:init //If you want a complete Docker image,that is, the src code will be included in the image`

2.  - If you are on Mac/Windows run this next-
`docker run --rm -it -p 8080:8080 booleans`
This will expose the docker container port 8080 to host port 8080
    - Linux has NOT been tested due to `host.docker.internal` being not supported. Testers and contributors are welcome to try the following command and modify the dockerfile accordingly-
`docker run -network=host --rm -it -p 8080:8080 booleans`
---

### Using the boolean service
The API provides the following services for now-
1. POST - Creates a new boolean with the value and key provided by the user. The response is a JSON consisting of UUID, value and key of the boolean.
2. GET - Request url contains the UUID of the boolean. The response is a JSON consisting of UUID, value and key of the boolean.
3. PATCH - Request url contains the UUID of the boolean which is to be updated. The HTTP request data field contains a JSON with value and key. The service updates the corresponding boolean and responds with a JSON consisting of UUID, value and key of the boolean.
4. DELETE- Request url contains the UUID of the boolean. The response is a HTTP 204 (no content) Status Code

#### POST
```
POST /
request:
{
  "value":true,
   "key": "name(Optional)"
}

response:
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
#### GET
```
GET /:id

response:
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
#### PATCH
```
PATCH /:id
request:
{
  "value":false,
  "key": "new name(Optional)"
}

response:
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "new name"
}
```
#### DELETE
```
DELETE /:id

response:
HTTP 204 No Content
```
---
### Example using cURL
### 1.
```
curl -X POST http://localhost:8080 --header "Content-Type: application/json" --data '{"value": false, "label": "brand1"}'
```
The above shall return the following-
`
{"id": <Some_UUID>, "value": false, "key": "brand1"}
`
### 2.
```
curl http://localhost:8080/SOME_UUID
```
The above shall return the following-
`
{"id": <Some_UUID>, "value": false, "key": "brand1"}
`
##### You can work with other API endpoints similarly