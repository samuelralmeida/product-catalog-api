package context

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *entity.User {
	val := ctx.Value(userKey)
	user, ok := val.(*entity.User)
	if !ok {
		return nil
	}
	return user
}
