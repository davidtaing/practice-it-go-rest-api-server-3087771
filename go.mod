module example.com/main

go 1.18

replace example.com/library => ./library

require example.com/library v0.0.0-00010101000000-000000000000

require github.com/mattn/go-sqlite3 v1.14.16 // indirect
