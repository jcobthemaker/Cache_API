package db

import (
    "context"
    "database/sql"
    "errors"
)

func Set(ctx context.Context, db *sql.DB, key, value string) error {
    query := `
        INSERT INTO result_tab ("key", "value")
        VALUES ($1, $2)
        ON CONFLICT ("key") DO UPDATE SET "value" = EXCLUDED.value
    `
    _, err := db.ExecContext(ctx, query, key, value)
    return err
}

func Get(ctx context.Context, db *sql.DB, key string) (string, error) {
    var value string

    query := `SELECT "value" FROM result_tab WHERE "key" = $1`

    err := db.QueryRowContext(ctx, query, key).Scan(&value)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return "", nil
        }
        return "", err
    }

    return value, nil
}

func GetAll(ctx context.Context, db *sql.DB) (map[string]string, error) {
	valueMap := make(map[string]string)
    query := `SELECT "key", "value" FROM result_tab`

    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return valueMap, err
    }
    defer rows.Close()

    for rows.Next() {
		var key, value string
        if err := rows.Scan(&key, &value); err != nil {
            return valueMap, err
        }
        valueMap[key] = value
    }


    if err = rows.Err(); err != nil {
        return valueMap, err
    }

    return valueMap, nil
}