package seeds

import (
	"context"
	"path"
	"transaction-service/internal/pkg/file"
	seedfiles "transaction-service/internal/testresources/seeds"
	"transaction-service/it/config/db"

	"github.com/pkg/errors"
)

// Load loads seeds in database
func Load(ctx context.Context, seeds ...string) error {
	for _, seed := range seeds {
		filepath := path.Join(seedfiles.Path, seed)

		stmt, err := file.LoadString(filepath)
		if err != nil {
			return errors.Wrapf(err, "[%s] failed to open seed file", seed)
		}

		if err := db.Exec(ctx, stmt); err != nil {
			return errors.Wrapf(err, "[%s] failed to load seed", seed)
		}
	}

	return nil
}
