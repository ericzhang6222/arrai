package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/arr-ai/arrai/examples/bundle/internal/arrai"
	"github.com/arr-ai/arrai/pkg/arraictx"
	"github.com/arr-ai/arrai/rel"
	"github.com/arr-ai/arrai/syntax"
)

func eval() (rel.Value, error) {
	bundle, err := arrai.Asset("internal/arrai/echo.arraiz")
	if err != nil {
		return nil, err
	}

	args := append([]string{""}, os.Args[1:]...)
	return syntax.EvaluateBundle(arraictx.InitRunCtx(context.Background()), bundle, args...)
}

func main() {
	val, err := eval()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(val.String())
	}
}
