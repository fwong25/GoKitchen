module gokitchen

go 1.22.2

replace models v1.0.0 => ./src/models

replace db_mgmt v1.0.0 => ./src/db_mgmt

replace interfaces v1.0.0 => ./src/interfaces

require (
	db_mgmt v1.0.0
	interfaces v1.0.0
)

require (
	github.com/lib/pq v1.10.9 // indirect
	models v1.0.0 // indirect
)
