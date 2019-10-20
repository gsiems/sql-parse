package sqlparse

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

/*
for each file in test/input/*sql
    - read and parse the file
    - read the like named file in test/expected
    - the compare of the two should match
*/
func TestSQLFiles(t *testing.T) {

	inputDir := "testdata/input"
	expectedDir := "testdata/expected"

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
	}

	for _, file := range files {
		// Ensure that it is a *.sql file
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		inputFile := inputDir + "/" + file.Name()
		expectedFile := expectedDir + "/" + file.Name()

		inBytes, err := ioutil.ReadFile(inputFile)
		if err != nil {
			t.Errorf(fmt.Sprintf("%s", err))
		}

		expBytes, err := ioutil.ReadFile(expectedFile)
		if err != nil {
			t.Errorf(fmt.Sprintf("%s", err))
		}

		input := string(inBytes)
		expected := string(expBytes)
		var dialect int

		// Extract the parsing args from the first line of the input
		// and determine which dialect to use
		l1 := strings.SplitN(input, "\n", 2)[0]
		args := strings.Split(strings.Replace(l1, "-", "", 2), ",")

		for i := 0; i < len(args); i++ {
			kv := strings.SplitN(args[i], ":", 2)
			if len(kv) > 1 {
				key := strings.Trim(kv[0], " ")
				value := strings.Trim(kv[1], " ")

				if key == "dialect" {
					dialect = SQLDialect(value)
				}
			}
		}

		tl := ParseStatements(input, dialect)

		tl.Rewind()
		var resultTokens []string

		for {
			t := tl.Next()
			s := t.Value()
			if s == "" {
				// nothing left to parse
				break
			}
			resultTokens = append(resultTokens, fmt.Sprintf("%s", t))
		}

		result := strings.Join(resultTokens, "\n") + "\n"
		if strings.Compare(result, expected) != 0 {
			t.Errorf("Comparison of %s failed", inputFile)

			// Output the results so they may be investigated/compared
			target := "testdata/result/" + file.Name()

			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				t.Errorf(fmt.Sprintf("File open failed: %q", err))
			} else {
				defer f.Close()

				if _, err := f.Write([]byte(result)); err != nil {
					t.Errorf(fmt.Sprintf("%s", err))
				}

				if err := f.Close(); err != nil {
					t.Errorf(fmt.Sprintf("%s", err))
				}
			}
		}
	}
}
