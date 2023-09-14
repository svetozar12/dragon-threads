package main

import (
	"dragon-threads/apps/api/internal/bootstrap"

	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Bootstrap()
}
