# Mongo-Cedar POC
This POC combines MongoDB and [Cedar Authorization Engine](https://docs.cedarpolicy.com/) to implement the authorization checks for CareGivers(CG) and CareReceipents(CR). The code uses Docker Compose for orchestration and creates 2 apps, one for the hosting Cedar engine and the webserver. The other one hosts MongoDB which is where the data about the entities live.

In this usecase, a CR can assign one or more CGs for managing each Capability and provide them with different Permissions.

Whenever the user invokes a `check` we compose the Cedar Entities and Context in real-time and feed it to the Authz Engine to verify if the data abides by the Policy.

## Entity Relationship Diagram
```mermaid
erDiagram
    CARE-GIVER }|--|{ CARE-RECEIPENT : has
    CARE-RECEIPENT }o--|{ CAPABILITY : accesses
    CAPABILITY ||--|{ PERMISSION-SET: has
    CARE-GIVER ||--|| PERMISSION-SET: has
```

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

## Workflow

```mermaid
sequenceDiagram
    actor cg as Care-Giver
    participant app as webserver
    participant authz as Cedar
    participant db as Mongo

    cg ->> app: check if authorized to <br/> tuple {cr,capability,action}
    app->>db: fetch info about cg-cr-capability mappings
    db-->>app: data
    app->>app: compose entity and context
    app->>authz: isAuthorized() with entity and context
    authz->>authz: check against policy
    authz-->>app: result: deny/allow
    app-->>cg: result

    
    
    
```

## K6 Data
We can use k6 to run perf tests.
Inorder to do this, we need to run the following cURL to setup the perf test data:
```bash
curl --location --request POST 'localhost:8888/insert-perf-data'
```

Once that is done, running a `check` would give us this:
![alt text](k6/1%20million%20CG-CR.png "Title")