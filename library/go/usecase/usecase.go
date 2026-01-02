package usecase

import "context"

type UseCase[Arg any, Res any] interface {
	Exec(ctx context.Context, arg Arg) (Res, error)
}
