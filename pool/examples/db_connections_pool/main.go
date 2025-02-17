package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	pool "github.com/debanjan97/pool"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	usePool              bool
	poolSize             int
	maxConcurrentQueries int
)

func newPostgresClient() *sql.Conn {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	return conn
}

func runQuery(db *sql.Conn, query string) {
	_, _ = db.ExecContext(context.Background(), query)
}

func withoutPool(numQueries int) {
	startTime := time.Now()
	defer func() {
		fmt.Printf("without pool: Time taken: %s\n", time.Since(startTime))
	}()
	var wg sync.WaitGroup
	for i := 0; i < numQueries; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := newPostgresClient()
			// defer client.Close()
			runQuery(client, "SELECT pg_sleep(1);")
		}()
	}
	wg.Wait()
}

func withPool(numQueries, poolSize int) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	p := pool.NewPool(poolSize, logger, newPostgresClient)
	startTime := time.Now()
	defer func() {
		fmt.Printf("with pool:   Time taken: %s\n", time.Since(startTime))
	}()

	var wg sync.WaitGroup
	for i := 0; i < numQueries; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := p.Get()
			defer p.Put(client)
			// defer client.Close()
			runQuery(client, "SELECT pg_sleep(1);")
		}()
	}
	wg.Wait()
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "dbpool",
		Short: "Database connection pool example",
		Run: func(cmd *cobra.Command, args []string) {
			if usePool {
				withPool(maxConcurrentQueries, poolSize)
			} else {
				withoutPool(maxConcurrentQueries)
			}
		},
	}

	rootCmd.Flags().BoolVar(&usePool, "with-pool", true, "Use connection pool")
	rootCmd.Flags().IntVar(&poolSize, "pool-size", 10, "Size of the connection pool")
	rootCmd.Flags().IntVar(&maxConcurrentQueries, "max-queries", 102, "Maximum number of concurrent queries")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
