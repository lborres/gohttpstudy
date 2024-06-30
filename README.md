# Go HTTP Study
Learning how to build a Go HTTP Server. 
Containerized and deployed to AWS via self-managed Elastic Container Service.

# Usage


---

# Field Notes
## Limited to Go Standard Library
For this project I have given myself the challenge of only using the Go Standard Library.

For Database driver, I went with the one recommended in the Go SQL Drivers Page [3].
Particularly, github.com/lib/pq for PostgreSQL

### Future Considerations
I plan to learn the following:
- [Gorilla Mux]() - Over standard http.ServeMux. But might change as the standard library is improved.
- [Labstack Echo - HTTP Framework]()
- Go with HTMX

## main() only calls run()
Reading up on the topic of go http servers I came across this article [1]. I figured its best to adapt what I can early in my learning journey.

## FormatDSN Function for postgres
Mysql driver has the convenience method, FormatDSN, for Config. [2]
PostgreSQL doesn't seem to have this so I made one emulating the one found in the Mysql driver.

## Database Migrations
I can't seem to find a straightforward way to automate database migrations without using non-standard packages. Decided to put this lower in the priority list and have the migrations manually for now.

```sh
migrateName="test"; touch $(date +%g%m%d%H%M%S)_${migrateName}.up.sql $(date +%g%m%d%H%M%S)_${migrateName}.down.sql
```

PostgreSQL has a built in function to create a new uuid
```sql
select gen_random_uuid()
```

## Subrouters
Wanted to emulate subrouter available with Gorilla Mux
but couldn't quite figure this out with the standard library.
Setting this as low priority.
```go
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
```

## Docker Dev mode
```sh
docker compose -f ./compose-dev.yml up -d
```

## Next Steps
- [ ] Password Encryption
- [ ] Database Migrations
- [ ] Experiment with having Handler functions return an error type
- [ ] Emulate subrouter available with Gorilla Mux



---

# References
[1] - [How I write HTTP services in Go after 13 years, Mat Ryer,  9 Feb, 2024](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) \
[2] - [FormatDSN Method, github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql/blob/v1.8.1/dsn.go#L226) \
[3] - [Go Wiki: SQL Database Drivers](https://go.dev/wiki/SQLDrivers) 


