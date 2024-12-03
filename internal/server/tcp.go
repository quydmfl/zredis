package server

import (
	"context"
)

type IProtocol interface {
	Start(context.Context, ) error
	Stop() error
}