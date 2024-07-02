package main

import (
	"flag"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/server"
)

func main() {
	repoType := flag.String("repo", "memory", "Repository type: memory, postgres")
	connStr := flag.String("conn", "postgres://username:password@localhost:5432/dbname", "Connection string for PostgreSQL")
	flag.Parse()

	// Создание и запуск сервера
	s := server.NewServer(*repoType, *connStr)
	s.Run()
}
