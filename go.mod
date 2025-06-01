module gitlab.com/scdb/server

go 1.24.3

require (
	gitlab.com/scdb/core v0.1.2 // indirect
	gitlab.com/scdb/database v0.1.2 // indirect
)

replace gitlab.com/scdb/core => ../core

replace gitlab.com/scdb/database => ../database
