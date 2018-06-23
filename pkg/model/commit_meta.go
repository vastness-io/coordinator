package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/vastness-io/coordinator/pkg/errors"
)

type CommitMeta struct {
	Languages Languages
}

func (c CommitMeta) Value() (driver.Value, error) {
	j, err := json.Marshal(c)
	return j, err
}

func (c *CommitMeta) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.NotByteSliceErr
	}

	var out CommitMeta
	err := json.Unmarshal(source, &out)
	if err != nil {
		return err
	}

	*c = out

	return nil
}
