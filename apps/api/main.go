package main

import (
	"dragon-threads/apps/api/bootstrap"

	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Bootstrap()
}
