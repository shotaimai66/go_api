package go_api

import (
	"context"
	"errors"

	"github.com/guregu/dynamo"
)

func scanUsers(ctx context.Context) ([]User, error) {
	var resp []User
	table := gdb.Table(usersTable)
	if err := table.Scan().AllWithContext(ctx, &resp); err != nil {
		// 0件の場合も正常とします
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, nil
		}
		return resp, err
	}
	return resp, nil
}