##ACL (Access Control List Role-Based Access)

### Programing Langauage used

1) Golang 1.10

### File structure
```
-- .acl
    |-- src
        |-- handler
            |--handler.go (contains api methods)
            |--handler_test.go
        |-- role
            |--role.go (contain setting and getting acl)
            |--role_test.go
        |-- router
            |--router.go (contain all api routes)
        |-- user
            |--user.go (contain setting and getting user by role id and user id and subordinates of given user)
            |--user_test.go
        |-- vednor (all package)
        |-- acl.go (entry point of app)
        |-- docker-compose.yml
        |-- Dockerfile (building image)
        |-- Gopkg.lock (package dependency lock)
        |-- Gopkg.toml
        |-- README.md (contain information about app like how to setup and use)
        |-- run.sh (command to run the application)
```

### How to run
1) Install Docker
2) Start Docker
3) docker-compose build
4) docker-compose up


### How to test
1. Run unit tests
go inside the dir github.com/srvsngh200892/acl and run
`go test ./... -v -cover`

2. How to create role
For example, here's the curl command

```
# set roles
curl -X POST http://localhost:3000/roles -d "replcae this with valid roles json"

output: 201
```

2. How to create role

```
# set users
curl -X POST http://localhost:3000/users -d ""replace this with valid users json""

output: 201
```

3. How to find subordinates

```
curl http://localhost:3000/subordinates/3

output: [{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]

```

4. Instead of curl you can use Postman to test rest api

### Validations

1. For role/user - all three attributes are mandatory other wise you will get bad request

2. For role/user - ID value should be greater than zero

3. For role/user - ID value should be unique

4. For role - Parent value should be greater than -1
