package db

import "context"

func (q *FakeQuerier) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	// TODO: Implement
	return User{}, nil
}
