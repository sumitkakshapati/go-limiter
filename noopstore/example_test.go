package noopstore_test

import (
	"context"
	"log"

	"github.com/sumitkakshapati/go-limiter/noopstore"
)

func ExampleNew() {
	ctx := context.Background()

	store, err := noopstore.New()
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close(ctx)

	limit, remaining, reset, _, ok, err := store.Take(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, _ = limit, remaining, reset, ok
}
