package internal_test

import (
	"github.com/beloys/goenum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDataFromDirectory(t *testing.T) {
	files := make(map[string]string)
	require.NoError(t, filepath.Walk("test_data", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			require.NoError(t, err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(path, ".in.go") {
			files[path] = strings.Replace(path, ".in.go", ".out.go", 1)
		}

		return nil
	}))
	for in, out := range files {
		do(in, out, t)
	}
}

func do(filePathIn string, filePathOut string, t *testing.T) {
	t.Logf("Read %s and compare with %s", filePathIn, filePathOut)
	parser, err := goenum.New(filePathIn)
	require.NoError(t, err)
	ref, err := parser.Scan()
	require.NoError(t, err)
	factory, err := goenum.NewTypeFactory().Create(ref)
	require.NoError(t, err)
	w := strings.Builder{}
	err = goenum.NewTemplatePrinter(ref, factory).Print(&w)
	require.NoError(t, err)
	res, err := ioutil.ReadFile(filePathOut)
	assert.NoError(t, err)
	assert.Equal(t, string(res), w.String())
}
