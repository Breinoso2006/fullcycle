package uow

import (
	"context"
	"database/sql"
	"fmt"
)

// any é o mesmo que a interface{} em Go, mas é mais explícito
type RepositoryFactory func(tx *sql.Tx) any

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	Unregister(name string)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	Rollback() error
	CommitOrRollback() error
	GetRepository(ctx context.Context, name string) (any, error)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx

	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("error during rollback: %v, original error: %v", errRb, err)
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return fmt.Errorf("no transaction to rollback")
	}

	if err := u.Tx.Rollback(); err != nil {
		return err
	}

	u.Tx = nil
	return nil
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("error during rollback: %v, original error: %v", errRb, err)
		}
		return err
	}

	u.Tx = nil
	return nil
}

func (u *Uow) GetRepository(ctx context.Context, name string) (any, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction: %w", err)
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}
