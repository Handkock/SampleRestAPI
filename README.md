### This is Sample RESTful API to manipulate a simple counter

In order to run the application, you have to ensure you have the docker and docker-compose running https://docs.docker.com/get-docker/

First of all, you need to create an `.env` file, you can do this by copying the `.env.development` file and renaming it. 

Then build and run the app with:
```bigquery
make up # this invokes docker-compose which installs the server and the database
```
The server starts at `localhost:8000` and is ready to recieve requests. Check out the OpenAPI specification here: https://app.swaggerhub.com/apis/Handkock/SampleRestAPI/1.0.0

The project contains three main parts:
- **main.go** - main entrypoint to invoke the server, init the counter and clean termination
- **app.go** - runs the server and handles the requests
- **counter.go** - Counter model, which takes over the db interaction and handles the counter operations

Tests are currently only to run due to host issues inside of container, so log in inside of the server containter and run the tests:
1. `docker exec -it restapi_server_1 /bin/bash`
2. ``make test``

