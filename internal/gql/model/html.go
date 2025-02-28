package model

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type Html string

func MarshalHtml(h string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, h))
	})
}

func UnmarshalHtml(v interface{}) (string, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return "", err
	}

	return value, nil
}
