package match

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID uuid.UUID `db:"id"`

	State State `db:"state"`
}

type State struct {
	CreatedAt  time.Time  `json:"created_at"`
	IsFinished bool       `json:"is_finished"`
	Comments   []Comment  `json:"comments"`
	Scoreboard Scoreboard `json:"scoreboard"`
}

func (s *State) Value() (driver.Value, error) {
	return json.Marshal(*s)
}

func (s *State) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("value cannot be converted to []byte")
	}

	return json.Unmarshal(b, s)
}

type Comment struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type Scoreboard struct {
	Players []Player `json:"players"`
}

type Role uint

const (
	RoleVillager Role = iota + 1
	RoleMafia
)

type Player struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
	IsAlive  bool   `json:"is_alive"`
	Kills    int    `json:"kills"`
}
