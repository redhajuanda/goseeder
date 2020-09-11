package seeds

import (
	"io/ioutil"
	"path"
	"runtime"
)

// ProductSeed seeds product data
func (s Seed) ProductSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/products.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
