package main

import (
	"os"
	"time"

	"github.com/tidwall/buntdb"
)

type BuntRunner struct {
	path string
}

func (r *BuntRunner) Run(c RunnerConfig) (*BenchResult, error) {
	r.path = c.path
	res := new(BenchResult)
	t := time.Now().UnixMilli()
	db, err := buntdb.Open(c.path)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res.InitTime = time.Now().UnixMilli() - t
	keys := generateKeys(c.count, 16, 16)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		err := db.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(string(k), string(randVal()), nil)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	res.WriteTime = time.Now().UnixMilli() - t
	shuffle(keys)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		err := db.View(func(tx *buntdb.Tx) error {
			_, e := tx.Get(string(k))
			if err != nil {
				return e
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	res.ReadTime = time.Now().UnixMilli() - t
	s, err := os.Stat(c.path)
	if err != nil {
		return nil, err
	}
	res.FSize += s.Size()

	return res, nil
}

func (r *BuntRunner) Clean() error {
	err := os.Remove(r.path)
	if err != nil {
		return err
	}
	return nil
}