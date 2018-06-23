package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/vastness-io/coordinator/pkg/errors"
)

type BranchMeta struct {
	Languages Languages
}

func (b BranchMeta) Value() (driver.Value, error) {
	j, err := json.Marshal(b)
	return j, err
}

func (b *BranchMeta) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.NotByteSliceErr
	}

	var out BranchMeta
	err := json.Unmarshal(source, &out)
	if err != nil {
		return err
	}

	*b = out

	return nil
}
