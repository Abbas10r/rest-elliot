package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Representation of a comment structure of our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Interface for repository
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// struct on which our logic will be built
type Service struct {
	Store Store
}

// Constructor. Returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}
