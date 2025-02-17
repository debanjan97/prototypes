module dbpool

go 1.23.5

replace github.com/debanjan97/pool => ../../pool

require (
	github.com/debanjan97/pool v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)
