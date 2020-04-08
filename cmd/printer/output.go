/*
Copyright 2020 The blcli Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package printer

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Format is the type of output to display
var Format string

// Output writes the output
func Output(data interface{}) {
	var err error
	switch Format {
	case "json":
		err = writeJSON(data)
	default:
		err = errors.New("unknown output format")
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func writeJSON(data interface{}) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}
