package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Profile struct {
	ID      uint
	Name    string
	Surname string
	City    string
}

func (p Profile) String() string {
	return fmt.Sprintf("%d. Фамилия: %s, Имя: %s, Город: %s\n", p.ID, p.Surname, p.Name, p.City)
}

func SearchByNameList(ctx context.Context, pool *pgxpool.Pool, input string) ([]Profile, error) {
	query := `
	SELECT 
		m.uniq_id, 
		n.name, 
		s.surname, 
		c.city
	FROM main m
	JOIN names n ON m.name_ = n.n_id
	JOIN surnames s ON m.surname = s.s_id
	LEFT JOIN cities c ON m.city = c.c_id -- LEFT JOIN на случай, если город не указан
	WHERE n.name ILIKE $1 OR n.name ILIKE '%' || $1 || '%'
	ORDER BY 
		-- 1. Приоритет: полное совпадение (Case-Insensitive)
		(n.name ILIKE $1) DESC,
		-- 2. Приоритет: начало строки (например, "Ив" найдет "Иван")
		(n.name ILIKE $1 || '%') DESC,
		-- 3. Остальные частичные совпадения
		n.name ASC;
	`
	rows, err := pool.Query(ctx, query, input)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Unable to open .env")
	}
	dsn := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic("Error while opening db")
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	res := make([]Profile, 0, 1)
	res, err = SearchByNameList(context.Background(), pool, "иван")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Найдено совпадений: ", len(res))
	fmt.Println(res)
}
