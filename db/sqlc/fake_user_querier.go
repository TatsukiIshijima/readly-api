package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserTable struct {
	// 自動インクリメンタル用
	NextID  int64
	Columns []User
}

var userTable = UserTable{NextID: 1}

func (q *FakeQuerier) CreateUser(_ context.Context, arg CreateUserParams) (User, error) {
	for _, user := range userTable.Columns {
		if user.Email == arg.Email {
			return User{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	now := time.Now().UTC()
	u := User{
		ID:             userTable.NextID,
		Name:           arg.Name,
		Email:          arg.Email,
		HashedPassword: arg.HashedPassword,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	userTable.Columns = append(userTable.Columns, u)
	userTable.NextID++
	return u, nil
}

func (q *FakeQuerier) GetUserByEmail(_ context.Context, email string) (User, error) {
	for _, user := range userTable.Columns {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, sql.ErrNoRows
}

func (q *FakeQuerier) GetUserByID(_ context.Context, id int64) (User, error) {
	return scanUser(id)
}

func (q *FakeQuerier) UpdateUser(_ context.Context, arg UpdateUserParams) (User, error) {
	user, err := scanUser(arg.ID)
	if err != nil {
		return User{}, err
	}
	if user.Email != arg.Email {
		for _, u := range userTable.Columns {
			// 変更後のメールアドレスが既に登録されている場合は重複エラー
			if u.Email == arg.Email {
				return User{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
			}
		}
	}
	for i, u := range userTable.Columns {
		if u.ID == arg.ID {
			userTable.Columns[i].Name = arg.Name
			userTable.Columns[i].Email = arg.Email
			userTable.Columns[i].HashedPassword = arg.HashedPassword
			userTable.Columns[i].UpdatedAt = time.Now().UTC()
			return userTable.Columns[i], nil
		}
	}
	return User{}, sql.ErrNoRows
}

func scanUser(id int64) (User, error) {
	for _, user := range userTable.Columns {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, sql.ErrNoRows
}
