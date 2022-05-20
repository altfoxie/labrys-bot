package storage

import (
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
)

// PastasRepository represents a repository for copypastas.
type PastasRepository struct {
	db *sqlx.DB
}

// Pasta is a stored copypasta.
type Pasta struct {
	// ID is a unique identifier of the copypasta.
	ID uint64 `db:"id"`
	// Name is a name of the copypasta.
	// It is used to search the copypasta.
	Name string `db:"name"`
	// Content is a text of the copypasta.
	Content string `db:"content"`
	// MessageID is a unique identifier of the message.
	MessageID uint64 `db:"message_id"`
}

// InlineResult returns inline query result for the copypasta.
func (p Pasta) InlineResult() telego.InlineQueryResult {
	return tu.ResultArticle(
		strconv.FormatUint(p.ID, 10),
		p.Name,
		tu.TextMessage(p.Content),
	)
}

// PastaMatch represents a matched copypasta.
type PastaMatch struct {
	Pasta
	// NameScore is a similarity of the copypasta's name and query.
	NameScore float64 `db:"name_score"`
	// ContentScore is a similarity of the copypasta's text and query.
	ContentScore float64 `db:"content_score"`
}

// Base returns Pasta.
func (m PastaMatch) Base() Pasta {
	return m.Pasta
}

// Init initializes the repository.
func (repo *PastasRepository) Init() error {
	return lo.T2(repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS pastas (
			id INTEGER PRIMARY KEY,
			name TEXT,
			content TEXT,
			message_id INTEGER
		);
		CREATE VIRTUAL TABLE IF NOT EXISTS pastas_fts USING fts5(name, tokenize=trigram)`,
	)).B
}

// Create creates a new copypasta.
func (repo *PastasRepository) Create(pasta Pasta) error {
	return lo.T2(repo.db.NamedExec(
		`INSERT INTO pastas (name, content, message_id) VALUES (:name, :content, :message_id);
		INSERT INTO pastas_fts (rowid, name) VALUES (last_insert_rowid(), :name)`,
		pasta,
	)).B
}

// Search searches copypastas by query.
func (repo *PastasRepository) Search(query string) (pastas []PastaMatch, err error) {
	l := strings.ToLower(query)
	return pastas, repo.db.Select(
		&pastas,
		`SELECT pastas.*, similarity(lower(pastas.name), lower(?)) AS name_score,
			similarity(lower(pastas.content), lower(?)) AS content_score
		FROM pastas_fts JOIN pastas ON pastas_fts.rowid = pastas.id
		WHERE name_score >= 0.25 OR content_score >= 0.5
		ORDER BY content_score DESC, name_score DESC LIMIT 50`,
		l, l,
	)
}

// GetN returns latest N copypastas.
func (repo *PastasRepository) GetN(n int) (pastas []Pasta, err error) {
	return pastas, repo.db.Select(&pastas, "SELECT * FROM pastas ORDER BY id DESC LIMIT ?", n)
}
