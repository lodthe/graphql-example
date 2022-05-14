package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/lodthe/graphql-example/internal/gqlgenerated"
	"github.com/lodthe/graphql-example/internal/gqlmodel"
	"github.com/lodthe/graphql-example/internal/match"
	"github.com/pkg/errors"
	zlog "github.com/rs/zerolog/log"
)

func (r *mutationResolver) CreateComment(ctx context.Context, comment gqlmodel.NewComment) (*gqlmodel.Comment, error) {
	parsed, err := uuid.Parse(comment.MatchID)
	if err != nil {
		return nil, errors.Wrap(err, "invalid match id UUID format")
	}

	m, err := r.repo.Get(parsed)
	if err != nil {
		if !errors.Is(err, match.ErrNotFound) {
			zlog.Err(err).Str("id", comment.MatchID).Msg("failed to get match to comment")
		}

		return nil, err
	}

	c := match.Comment{
		ID:   uuid.New(),
		Text: comment.Text,
	}
	m.State.Comments = append(m.State.Comments, c)

	err = r.repo.Update(m)
	if err != nil {
		zlog.Err(err).Str("id", m.ID.String()).Msg("failed to update match")
		return nil, err
	}

	converted := commentToModel(c)

	return &converted, nil
}

func (r *queryResolver) Match(ctx context.Context, id string) (*gqlmodel.Match, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid UUID format")
	}

	m, err := r.repo.Get(parsed)
	if err != nil {
		if !errors.Is(err, match.ErrNotFound) {
			zlog.Err(err).Str("id", id).Msg("failed to get match")
		}

		return nil, err
	}

	return matchToModel(m), err
}

func (r *queryResolver) Matches(ctx context.Context, isFinished *bool, limit *int, offset *int) ([]gqlmodel.Match, error) {
	matches, err := r.repo.GetAll()
	if err != nil {
		zlog.Err(err).Msg("failed to get matches")
		return nil, err
	}

	var result []gqlmodel.Match
	for _, m := range matches {
		if isFinished != nil && *isFinished != m.State.IsFinished {
			continue
		}

		if offset != nil && *offset != 0 {
			*offset--
			continue
		}

		if limit != nil && *limit < len(result)+1 {
			break
		}

		result = append(result, *matchToModel(&m))
	}

	return result, nil
}

// Mutation returns gqlgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgenerated.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgenerated.QueryResolver implementation.
func (r *Resolver) Query() gqlgenerated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
