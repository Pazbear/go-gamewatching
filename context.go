package main

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type ReqContext[T any] struct{
	Config Config
	Request T
}

const contextKey = "pazbear-context"

func GinContext[T any] (c *gin.Context, cnf Config)(context.Context, error){
	var reqbody T
	if err:=c.Bind(&reqbody); err != nil {
		return nil, err
	}
	return context.WithValue(c, contextKey, ReqContext[T]{
		Config: cnf,
		Request: reqbody,
	}), nil
}

func ReqContextWithType[T any](ctx context.Context)(T, error){
	var retval T
	reqany := ctx.Value(contextKey)
	if reqany == nil {
		return retval, errors.New("not found")
	}
	reqconc, ok := reqany.(T)
	if !ok {
		return retval, errors.New("invalid type")
	}
	return reqconc, nil
}