package ctx

import (
	"context"
	"fmt"
)

type ID int

const userIDKey ID = 1

func valueSetter(ctx context.Context, val any, callback func(context.Context) (string, error)) string {

	ctxWithVal := context.WithValue(ctx, userIDKey, val)

	id, err := callback(ctxWithVal)
	if err != nil {
		fmt.Println("could not retrieve id: ", err)
		return ""
	}

	return id
}

func valueGetter(ctx context.Context) (string, error) {

	id, ok := ctx.Value(userIDKey).(string)

	if !ok {
		return "", fmt.Errorf("id is not a string")
	}

	return id, nil

}

func UseCTXWithValue() {
	ctx := context.Background()

	idError := valueSetter(ctx, 123, valueGetter)
	fmt.Println(idError)

	idSuccess := valueSetter(ctx, "adad21fada", valueGetter)
	fmt.Println(idSuccess)
}
