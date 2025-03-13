# Mongo-Cedar POC
This POC combines MongoDB and [Cedar Authorization Engine](https://docs.cedarpolicy.com/) to implement the authorization checks for CareGivers(CG) to perform actions on behalf of CareReceipents(CR). The code uses Docker Compose for orchestration and creates 2 apps, one for the hosting Cedar engine and the webserver. The other one hosts MongoDB which is where the data about the entities live.

The usecase can be explained as follows:
A CR can assign one or more CGs for managing each Capability (in this case, UserProfile and Documents) and provide them with different Permissions at the capability level.

Whenever the CG invokes a `check` , to check if they have access to a tuple `{cr id, capability id, action}`, we compose the Cedar Entities and Context in real-time and feed it to the Authz Engine to verify if the data abides by the Policy.

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

### Accessing MongoDB
```bash
root@b8b9fa2a5312:/# mongosh
Current Mongosh Log ID: 67d30f30d3027ee69f584a20
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.4.0
Using MongoDB:          8.0.5
Using Mongosh:          2.4.0

For mongosh info see: https://www.mongodb.com/docs/mongodb-shell/


To help improve our products, anonymous usage data is collected and sent to MongoDB periodically (https://www.mongodb.com/legal/privacy-policy).
You can opt-out by running the disableTelemetry() command.

test> 

test> 

test> use admin
switched to db admin
admin> db.auth('rootuser','rootpass')
{ ok: 1 }
admin> use mydb
switched to db mydb

mydb> show collections
caregivers
carereceipents

mydb> db.carereceipents.find().count()
1000000
mydb> db.caregivers.find().count()
100
mydb> 
```

## Workflow
How the authz check is meant to happen using `Cedar + Mongo`

```mermaid
sequenceDiagram
    autonumber
    actor cg as Care-Giver
    participant app as webserver
    participant authz as Cedar
    participant db as Mongo

    note right of cg: Assume Mongo and Cedar are already primed with data.
    cg ->> app: check if authorized to <br/> tuple {cr,capability,action}
    app->>db: fetch info about cg-cr-capability-perm mappings
    db-->>app: data
    app->>app: compose entity and context for cedar to consume
    app->>authz: isAuthorized() with entities and context <br/> along with the initial tuple
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
**NOTE** This script will take a bit of time (~3min) to insert 1M documents.

Once that is done, running a `check` would give us this:
![alt text](k6/k6-1mil-cr-cg.png "Performance Analysis")