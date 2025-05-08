package repository

import (
	"context"
	"fmt"

	"github-tracker/repository/entity"
)

type Commit interface {
	Insert(ctx context.Context, commit *entity.Commit) error
}

type CommitRepo struct{}

func (r *CommitRepo) Insert(ctx context.Context, commit *entity.Commit) error {
	fmt.Printf("Simulated insert: %+v\n", commit)
	return nil
}
