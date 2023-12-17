package router

import (
	"github.com/WildEgor/fiber-gql/internal/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(NewRouter, handlers.HandlersSet)
