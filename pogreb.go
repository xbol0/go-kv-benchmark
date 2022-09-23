package main

import (
	"os"
	"time"

	"github.com/akrylysov/pogreb"
)

type PogrebRunner struct {
	path string
}

func (r *PogrebRunner) Run(c RunnerConfig) (*BenchResult, error) {
	r.path = c.path
	res := new(BenchResult)
	t := time.Now().UnixMilli()
	db, err := pogreb.Open(c.path, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res.InitTime = time.Now().UnixMilli() - t
	keys := generateKeys(c.count, 16, 16)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		err := db.Put(k, []byte(randVal()))
		if err != nil {
			return nil, err
		}
	}

	res.WriteTime = time.Now().UnixMilli() - t
	shuffle(keys)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		_, err := db.Get(k)
		if err != nil {
			return nil, err
		}
	}
	res.ReadTime = time.Now().UnixMilli() - t
	size, err := DirSize(c.path)
	if err != nil {
		return nil, err
	}
	res.FSize = size

	return res, nil
}

func (r *PogrebRunner) Clean() error {
	err := os.RemoveAll(r.path)
	if err != nil {
		return err
	}
	return nil
}