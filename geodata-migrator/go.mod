module github.com/watchedsky-social/applications/geodata-migrator

go 1.25.5

require (
	github.com/alecthomas/kong v1.13.0
	github.com/watchedsky-social/libwatchedsky v0.0.1
)

require (
	github.com/jghiloni/go-commonutils/v3 v3.3.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.33 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/paulmach/orb v0.12.0 // indirect
	github.com/pressly/goose/v3 v3.26.0 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	github.com/watchedsky-social/geodata v0.0.2 // indirect
	github.com/watchedsky-social/go-spatialite v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
)

replace (
	github.com/watchedsky-social/libwatchedsky => ../../libwatchedsky
)
