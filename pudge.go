package main

import (
	"os"
	"time"

	"github.com/recoilme/pudge"
)

type PudgeRunner struct {
	path string
}

func (r *PudgeRunner) Run(c RunnerConfig) (*BenchResult, error) {
	r.path = c.path
	res := new(BenchResult)
	t := time.Now().UnixMilli()
	db, err := pudge.Open(c.path, pudge.DefaultConfig)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res.InitTime = time.Now().UnixMilli() - t
	keys := generateKeys(c.count, 16, 16)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		err := db.Set(k, randVal())
		if err != nil {
			return nil, err
		}
	}

	res.WriteTime = time.Now().UnixMilli() - t
	shuffle(keys)

	t = time.Now().UnixMilli()
	for _, k := range keys {
		var val string
		err := db.Get(k, &val)
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
	s, err = os.Stat(c.path + ".idx")
	if err != nil {
		return nil, err
	}
	res.FSize += s.Size()

	return res, nil
}

func (r *PudgeRunner) Clean() error {
	err := os.Remove(r.path)
	if err != nil {
		return err
	}
	err = os.Remove(r.path + ".idx")
	if err != nil {
		return err
	}
	return nil
}
