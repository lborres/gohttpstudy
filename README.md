# Go HTTP Study
Learning how to build a Go HTTP Server. 
Containerized and deployed to AWS via self-managed Elastic Container Service.

# Usage



# Field Notes
## Limited to Go Standard Library
For this project I have given myself the challenge of only using the Go Standard Library.

For Database driver, I went with the one recommended in the Go SQL Drivers Page [3]
Particularly, github.com/lib/pq for PostgreSQL

### Future Considerations
I plan to learn the following:
- [Gorilla Mux]() - Over standard http.ServeMux. But might change as the standard library is improved.
- [Labstack Echo - HTTP Framework]()
- Go with HTMX

## main() only calls run()
[1] 

## FormatDSN Function for postgres
Mysql driver has the convenience method, FormatDSN, for Config. [2]
PostgreSQL doesn't seem to have this so I made one emulating the one found in the Mysql driver.


## Next Steps


# References
[1]: https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/ "How I write HTTP services in Go after 13 years, Mat Ryer,  9 Feb, 2024"

[2]: https://github.com/go-sql-driver/mysql/blob/v1.8.1/dsn.go#L226 "FormatDSN Method, github.com/go-sql-driver/mysql"

[3]: https://go.dev/wiki/SQLDrivers "Go Wiki: SQL Database Drivers"


