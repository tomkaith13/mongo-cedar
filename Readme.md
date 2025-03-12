# Mongo-Cedar POC
This POC combines MongoDB and Cedar Authorization Engine to implement the authorization checks for CareGivers and CareReceipents. The code uses Docker Compose for orchestration and creates 2 apps, one for the hosting Cedar engine and the webserver. The other one hosts MongoDB.

## Entity Relationship Diagram
Explains how the entities are setup

TODO: Add mermaid ERD
## Instructions
Use instructions in `Makefile` to start the service
```bash
make up
```
To clean up after the run, do:
```bash
make clean
```

Use the following to set up a sample CG-CR pair:
```bash
curl --location --request POST 'localhost:8888/insert-example'
```

And use `check` to verify authz:
```bash
curl --location 'localhost:8888/check' --header 'Content-Type: application/json' --data '{"cg":"cg1","cr":"cr14","action":"READ","resource":"UserProfile"}'
```

**NOTE:** Currently the 2 resources configured are `UserProfile` and `Documents`

## K6 Data
We can use k6 to run perf tests.
Inorder to do this, we need to run the following cURL to setup the perf test data:
```bash
curl --location --request POST 'localhost:8888/insert-perf-data'
```

Once that is done, running a `check` would give us this:
![alt text](k6/1%20million%20CG-CR.png "Title")