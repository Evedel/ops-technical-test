# MYOBTestService [![Build Status](https://travis-ci.com/Evedel/ops-technical-test.svg?branch=master)](https://travis-ci.com/Evedel/ops-technical-test)

> ## Platform Enablement Technical Test

> Your application should be a simple, small, operable web-style API or service provider.

To keep things simple, I chose to use `go` with no additional third party packages for my application.

The application implements:
- `appaddr/` a root endpoint, which responds `hello world`
- `appaddr/healthcheck` a health endpoint, which returns http code `200` if application status is healthy and `500` if it is not
- `appaddr/metadata` a metadata endpoint, which returns basic information about application
```
"myapplication": [
  {
    "version": "1.0",
    "description" : "pre-interview technical test",
    "lastcommitsha": "abc57858585"
  }
]
```

### The project also contains:
- a test suite: the set of unit tests to check the correctness of the code and python script to check if the endpoints are accessible.
- a means of packaging my application as a single deployable artefact ([Dockerfile](Dockerfile))
- a pipeline that builds my application on each commit ([.travis.yml](.travis.yml))

### The application and its deployment steps:
The application has following packages:
- `main` (registers the endpoints and starts listener at `:8080`)
- `handlers` (functions that convert go structures to strings and write them into html responses)
- `healthcheck` (runs health check and returns the response code)
- `storage` (obtains and constructs the application metadata)

Deployment contains three stages:
- executing unit test
- executing api call tests
- building docker image and pushing it to the `evedel/myobtestservice`

One can run the application with

```
docker run -rm -p12345:8080 evedel/myobtestservice
```

### Possible risks associated with my application/deployment:
- there is no https (service might be put behind secure reverse proxy)
- there is no serving port configuration (docker port forwarding might be used for this means)
- there might be bugs
