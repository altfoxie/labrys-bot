package storage

import (
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
)

// VoicesRepository represents a repository for voice messages.
type VoicesRepository struct {
	db *sqlx.DB
}

// Voice is a stored voice message.
type Voice struct {
	// ID is a unique identifier of the voice message.
	ID uint64 `db:"id"`
	// Name is a name of the voice message.
	// It is used to search the voice message.
	Name string `db:"name"`
	// FileID is a unique identifier of the file, stored on the Telegram servers.
	FileID string `db:"file_id"`
	// MessageID is a unique identifier of the message.
	MessageID int `db:"message_id"`
}

func (v Voice) InlineResult() telego.InlineQueryResult {
	return tu.ResultCachedVoice(
		strconv.FormatUint(v.ID, 10),
		v.FileID,
		v.Name,
	)
}

// VoiceMatch represents a matched voice message.
type VoiceMatch struct {
	Voice
	// Score is a similarity of the voice message name and query.
	Score float64 `db:"score"`
}

// Base returns Voice.
func (m VoiceMatch) Base() Voice {
	return m.Voice
}

// Init initializes the repository.
func (repo *VoicesRepository) Init() error {
	return lo.T2(repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS voices (
			id INTEGER PRIMARY KEY,
			name TEXT,
			file_id TEXT,
			message_id INTEGER
		);
		CREATE VIRTUAL TABLE IF NOT EXISTS voices_fts USING fts5(name, tokenize=trigram)`,
	)).B
}

// Create creates a new voice message.
func (repo *VoicesRepository) Create(voice Voice) error {
	return lo.T2(repo.db.NamedExec(
		`INSERT INTO voices (name, file_id, message_id) VALUES (:name, :file_id, :message_id);
		INSERT INTO voices_fts (rowid, name) VALUES (last_insert_rowid(), :name)`,
		voice,
	)).B
}

// Search searches voice messages by name.
func (repo *VoicesRepository) Search(name string) (voices []VoiceMatch, err error) {
	return voices, repo.db.Select(
		&voices,
		`SELECT voices.*, similarity(lower(voices.name), lower(?)) AS score
		FROM voices_fts JOIN voices ON voices_fts.rowid = voices.id
		WHERE score >= 0.25 ORDER BY score DESC LIMIT 50`,
		strings.ToLower(name),
	)
}

// GetN returns latest N voice messages.
func (repo *VoicesRepository) GetN(n int) (voices []Voice, err error) {
	return voices, repo.db.Select(&voices, "SELECT * FROM voices ORDER BY id DESC LIMIT ?", n)
}
