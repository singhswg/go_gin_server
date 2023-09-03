# Run local postgres 

`docker compose up` and then go to `localhost:8080`

# Login using psql - 

```
❯ psql -h 127.0.0.1 -U postgres                   
Password for user postgres: 
psql (14.8 (Homebrew), server 15.3 (Debian 15.3-1.pgdg120+1))
WARNING: psql major version 14, server major version 15.
         Some psql features might not work.
Type "help" for help.

postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)

postgres=# \q

```

Run `source setenv.sh` to set up Postgres specific environment variables 