package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	languageMetaKey = "languages"
)

var (
	TypeAssertByteErr               = errors.New("unable to type assert to map[string] interface{}")
	TypeAssertMapStringInterfaceErr = errors.New("unable to type assert to map[string] interface{}")
)

type BranchLanguages map[string]float64

type BranchMeta map[string]interface{}

func (b BranchMeta) SetLanguages(branchLanguages BranchLanguages) {
	b[languageMetaKey] = branchLanguages
}

func (b BranchMeta) GetLanguages() BranchLanguages {
	m, ok := b[languageMetaKey].(map[string]interface{})

	if !ok {
		return nil
	}

	out := make(BranchLanguages)

	for k, v := range m {

		fv, ok := v.(float64)

		if !ok {
			continue
		}

		out[k] = fv
	}

	return out
}

func (b BranchMeta) Value() (driver.Value, error) {
	j, err := json.Marshal(b)
	return j, err
}

func (b *BranchMeta) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return TypeAssertByteErr
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*b, ok = i.(map[string]interface{})
	if !ok {
		return TypeAssertMapStringInterfaceErr
	}

	return nil
}

type Branch struct {
	ID              int64 `gorm:"primary_key"`
	Name            string
	Meta            BranchMeta
	Commits         []*Commit `gorm:"many2many:branch_commits"`
	RepositoryName  string
	RepositoryOwner string
	RepositoryType  string
}
