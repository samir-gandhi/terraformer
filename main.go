// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"path/filepath"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/cmd"
)

type TerraformerWriter struct {
	io.Writer
}

func (t TerraformerWriter) Write(p []byte) (n int, err error) {
	if !strings.Contains(string(p), "[TRACE]") && !strings.Contains(string(p), "[DEBUG]") { // hide TF GRPC client log messages
		return os.Stdout.Write(p)
	}
	return len(p), nil
}

func main() {
	log.SetOutput(TerraformerWriter{})
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := filepath.Walk("generated", walkFunc); err != nil {
		panic(err)
	}
}

// DO NOT COMMIT THIS FUNCTION
func walkFunc(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}
	matched, err := filepath.Match("*.tf", fi.Name())
	if err != nil {
		return err
	}

	if matched {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		r := regexp.MustCompile(`"\${file\(\\\"data\/.*json\\\"\)}"`)
		items := r.FindAllStringSubmatch(string(b), -1)
		for _, item := range items {
			replace := strings.Replace(item[0], "\\", "", -1)
			os.WriteFile(path, r.ReplaceAllLiteral(b, []byte(replace)), 0600)
		}
	}
	return nil
}
