/*

Print a list of test cases and test functions found from the requested
directory (or from the "tests/" directory by default).

This can be used for auto-completion.

*/

package main

import "bufio"
import "fmt"
import "os"
import "path/filepath"
import "regexp"
import "strings"


var pythonRegexpClassName = regexp.MustCompile("^class (\\w*)")
var pythonRegexpFuncName = regexp.MustCompile("^(    )?def (test\\w*)")


func parsePythonFile(path string) []string {
	module := strings.TrimSuffix(path, ".py")
	module = strings.Replace(module, "/", ".", -1)
	functions := make([]string, 0)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	currentClass := ""
	for scanner.Scan() {
		line := scanner.Text()
		match := pythonRegexpClassName.FindStringSubmatch(line)
		if len(match) > 0 {
			currentClass = match[1]
			functions = append(functions, path + "::" + currentClass)
		} else {
			match := pythonRegexpFuncName.FindStringSubmatch(line)
			if len(match) > 0 {
				var funcName = path + "::" + match[2]
				if match[1] != "" {
					funcName = path + "::" + currentClass + "::" + match[2]
				}
				functions = append(functions, funcName)
			}
		}
	}
	return functions
}


func walkDir(dir string, filePattern string) ([]string, error) {
	var testFunctions = make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			isTestFile, _ := regexp.MatchString(filePattern, path)
			if isTestFile {
				testFunctions = append(testFunctions, parsePythonFile(path)...)
			}
		}
		return nil  // no error
	})
	return testFunctions, err
}


func main() {
	startDir := ""
	if len(os.Args) < 2 {
		startDir = "tests"
	} else {
		startDir = os.Args[1]
	}
	filePattern := "(tests\\.py$)|(test_.*?\\.py$)"
	var testFunctions, err = walkDir(startDir, filePattern)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, testFunc := range testFunctions {
		fmt.Println(testFunc)
	}
}
