# Go DB Service 

Go service that leverages `gin` package and exposes various endpoints to interact with a Postgres endpoint.


### Project setup

- Need to have `Go` installed

- After forking the repo, execute - 
    ```
    go mod init github.com/singhswg/go_gin_server # Modify as needed
    go mod tidy
    ```

- Local database should be running and database credentials should be present in local environment/session

### Dependencies

`go get -u github.com/lib/pq`

### Endpoints supported so far 

- `users`: To list all users

- `ping`: Returns a string

- `adduser`: POST request to add a user
    ```
    ❯ curl localhost:8080/adduser --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "User", "title": "XYZ", "city": "Toronto"}'
    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Sun, 03 Sep 2023 01:20:26 GMT
    Content-Length: 25

    {"data":"User was added"}%   
    ```

- `users/<some_user> `: GET request to check if a user exists
    ```
    ❯ curl localhost:8080/users/James

    ❯ curl localhost:8080/users/Balpreet
    "User found"
    
    ❯ curl localhost:8080/users/Balpreeth
    "User not found"
    ```

- `users` : GET request to get all users
    ```
    ❯ curl localhost:8080/users
    ```

- Default message for all other endpoints

## Database setup
Run `docker compose up` within folder `postgres`

Create the following if doesnt already exists - 

```
CREATE TABLE users( name TEXT, title TEXT, city TEXT );
INSERT INTO users (name, title, city) VALUES ('Balpreet','Lead', 'Chicago');
SELECT * from users; # Test
```

### Yet to be added

- [ ] Authentication

- [ ] Containerization and Helm support

- [ ] Registry and K8S manifests support

- [ ] Makefile

- [ ] Vault support

- [ ] Logging

- [ ] Github Actions CI/CD

### References

- https://pkg.go.dev/github.com/gin-gonic/gin 
- https://gin-gonic.com/docs/examples 