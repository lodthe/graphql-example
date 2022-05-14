package resolve

import (
	"github.com/lodthe/graphql-example/internal/gqlmodel"
	"github.com/lodthe/graphql-example/internal/match"
)

func matchToModel(m *match.Match) *gqlmodel.Match {
	s := m.State

	return &gqlmodel.Match{
		ID:         m.ID.String(),
		CreatedAt:  s.CreatedAt.String(),
		IsFinished: s.IsFinished,
		Comments:   commentsToModel(s.Comments),
		Scoreboard: scoreboardToModel(&s.Scoreboard),
	}
}

func commentsToModel(comments []match.Comment) []gqlmodel.Comment {
	result := make([]gqlmodel.Comment, 0, len(comments))
	for _, c := range comments {
		result = append(result, commentToModel(c))
	}

	return result
}

func commentToModel(comment match.Comment) gqlmodel.Comment {
	return gqlmodel.Comment{
		ID:   comment.ID.String(),
		Text: comment.Text,
	}
}

func scoreboardToModel(s *match.Scoreboard) *gqlmodel.Scoreboard {
	scoreboard := &gqlmodel.Scoreboard{
		Players: make([]gqlmodel.Player, 0, len(s.Players)),
	}
	for _, p := range s.Players {
		scoreboard.Players = append(scoreboard.Players, playerToModel(&p))
	}

	return scoreboard
}

func playerToModel(p *match.Player) gqlmodel.Player {
	return gqlmodel.Player{
		Username: p.Username,
		Role:     roleToModel(p.Role),
		IsAlive:  p.IsAlive,
		Kills:    p.Kills,
	}
}

func roleToModel(role match.Role) gqlmodel.Role {
	switch role {
	case match.RoleVillager:
		return gqlmodel.RoleVillager

	case match.RoleMafia:
		return gqlmodel.RoleMafia

	default:
		return "UNKNOWN"
	}
}
