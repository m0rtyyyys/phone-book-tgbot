package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Profile struct {
	ID      uint
	Name    string
	Surname string
	City    string
}

func (p Profile) String() string {
	return fmt.Sprintf("%d. Фамилия: %s, Имя: %s, Город: %s", p.ID, p.Surname, p.Name, p.City)
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

func SearchBySurnameList(ctx context.Context, pool *pgxpool.Pool, input string) ([]Profile, error) {
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
	WHERE s.surname ILIKE $1 OR s.surname ILIKE '%' || $1 || '%'
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

func SearchByCityList(ctx context.Context, pool *pgxpool.Pool, input string) ([]Profile, error) {
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
	WHERE c.city ILIKE $1 OR c.city ILIKE '%' || $1 || '%'
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
