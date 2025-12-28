package main

import (
	"fmt"
	"strings"

	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("cli> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	fmt.Println("CLI")

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		args := strings.Fields(line)

		switch args[0] {

		case "exit", "quit":
			return

		case "help":
			help()

		case "clear":
			fmt.Println("\033[H\033[2J")

		case "ls":
			viewTables()

		case "query", "-q":
			sql := strings.Join(args[1:], " ")
			query(sql)

		default:
			fmt.Println("comando desconocido")
		}
	}
}

func query(sql string) {
	var results []map[string]interface{}

	err := db.DB.Raw(sql).Scan(&results).Error
	if err != nil {
		fmt.Println("Error al ejecutar consulta:", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("(sin resultados)")
		return
	}

	for i, row := range results {
		fmt.Printf("Row %d:\n", i+1)
		for k, v := range row {
			fmt.Printf("  %s = %v\n", k, v)
		}
	}
}

func dropTable(table string) {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)

	if err := db.DB.Exec(query); err != nil {
		fmt.Println("Database not exists")
		return
	}

	fmt.Println("Database deleted")
}

func viewTables() {
	var tables []string

	err := db.DB.
		Raw("SELECT name FROM sqlite_master WHERE type='table';").
		Scan(&tables).Error

	if err != nil {
		fmt.Println("Error al obtener tablas:", err)
		return
	}

	fmt.Println("Tablas:")
	for _, t := range tables {
		fmt.Println(" -", t)
	}
}

func help() {
	fmt.Println(`
Comandos disponibles:

drop <tabla>     Elimina una tabla
migrate          Ejecuta migraciones
help             Muestra ayuda
exit             Salir del CLI

Ejemplo:
cli> drop users
`)
}
