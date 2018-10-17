package user

import "context"

type IUser interface {
	Register(ctx context.Context, phone string) (userId string, err error)
}
