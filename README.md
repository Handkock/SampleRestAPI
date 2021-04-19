###This is Sample RESTful API to manipulate a simple counter

In order to run the application, you have to install the docker and docker-compose running https://docs.docker.com/get-docker/

Build and run the app with:
```bigquery
    make up # this invokes docker-compose which installs the server and the database
```
The server starts at `localhost:8000` and is ready to recieve requests. Check out the OpenAPI specification here: https://app.swaggerhub.com/apis/Handkock/SampleRestAPI/1.0.0 

The project contains three main parts:
- **main.go** - main entrypoint to invoke the server, init the counter and clean termination
- **app.go** - runs the server and handles the requests
- **counter.go** - Counter model, which takes over the db interaction and handles the counter operations


