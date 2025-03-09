package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Fortune struct {
	Id        int64
	Kind      FortuneKind
	Body      string
	CreatedAt time.Time
	Opened    bool
}

type FortuneKind string

const (
	KindMemory FortuneKind = "memory"
	KindWish   FortuneKind = "wish"
)

func GetMemories(db *sql.DB) ([]Fortune, error) {
	return GetFortunes(db, KindMemory, false)
}

func GetWishes(db *sql.DB) ([]Fortune, error) {
	return GetFortunes(db, KindWish, false)
}

func GetOpened(db *sql.DB) ([]Fortune, error) {
	openedMemories, err := GetFortunes(db, KindMemory, true)
	if err != nil {
		return nil, err
	}
	openedWishes, err := GetFortunes(db, KindWish, true)
	if err != nil {
		return nil, err
	}
	return append(openedMemories, openedWishes...), nil
}

func GetFortunes(db *sql.DB, kind FortuneKind, opened bool) ([]Fortune, error) {
	var fortunes []Fortune
	rows, err := db.Query("SELECT * FROM fortune WHERE kind = ? AND opened = ? ORDER BY created_at DESC", kind, opened)
	if err != nil {
		return nil, fmt.Errorf("fortunes: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var a Fortune
		if err := rows.Scan(&a.Id, &a.Kind, &a.Body, &a.CreatedAt, &a.Opened); err != nil {
			return nil, fmt.Errorf("fortune: %v", err)
		}
		fortunes = append(fortunes, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fortunes: %v", err)
	}
	return fortunes, nil
}

func CreateFortune(db *sql.DB, kind FortuneKind, body string) (int64, error) {
	result, err := db.Exec("INSERT INTO fortune (kind, body) VALUES (?, ?)", kind, body)
	if err != nil {
		return 0, fmt.Errorf("addFortune: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addFortune: %v", err)
	}
	return id, nil
}

func GetFortune(db *sql.DB, id int64) (Fortune, error) {
	var a Fortune
	row := db.QueryRow("SELECT * FROM fortune WHERE id = ?", id)
	if err := row.Scan(&a.Id, &a.Kind, &a.Body, &a.CreatedAt, &a.Opened); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("fortunesById %d: no such fortune", id)
		}
		return a, fmt.Errorf("fortunesById %d: %v", id, err)
	}
	return a, nil
}

func OpenFortune(db *sql.DB, id int64) (int64, error) {
	result, err := db.Exec("UPDATE fortune SET opened = true WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("openFortune %d: %v", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("openFortune %d: %v", id, err)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("openFortune %d: no such fortune", id)
	}
	return id, nil
}

func DeleteFortune(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM fortune WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteFortune %d: %v", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteFortune %d: %v", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("deleteFortune %d: no such fortune", id)
	}
	return nil
}

func DeleteAllFortunes(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM fortune")
	if err != nil {
		return fmt.Errorf("deleteAllFortunes: %v", err)
	}
	return nil
}

func (a *Fortune) GetStrId() string {
	return fmt.Sprintf("%d", a.Id)
}

func (a *Fortune) GetTitle() string {
	if a.Kind == KindWish {
		return "✨"
	}
	return "🪂"
}
