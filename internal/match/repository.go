package match

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type Repository interface {
	Create(match *Match) error
	Update(match *Match) error
	Get(id uuid.UUID) (*Match, error)
	GetAll() ([]Match, error)
}

type Repo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(match *Match) error {
	query := `INSERT INTO matches (id, state) VALUES (:id, :state)`
	_, err := r.db.NamedExec(query, match)
	if err != nil {
		return errors.Wrap(err, "database error")
	}

	return nil
}

func (r *Repo) Update(match *Match) error {
	_, err := r.db.NamedExec(`UPDATE matches SET state = :state WHERE id = :id`, match)

	return err
}

func (r *Repo) Get(id uuid.UUID) (*Match, error) {
	match := new(Match)
	err := r.db.Get(match, `SELECT * FROM matches WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "database error")
	}

	return match, nil
}

func (r *Repo) GetAll() ([]Match, error) {
	var matches []Match
	err := r.db.Select(&matches, "SELECT * FROM matches ORDER BY state->'created_at'")
	if err != nil {
		return nil, err
	}

	return matches, nil
}
