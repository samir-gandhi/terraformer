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

package pingonedavinci

import (
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var (
	VariableAllowEmptyValues = []string{}
)

type VariableGenerator struct {
	PingOneDavinciService
}

func (g VariableGenerator) createResources(variables map[string]davinci.Variable) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	names := map[string]struct{}{}
	for id, variable := range variables {
		s := strings.Split(id, "##SK##")
		name := s[0]
		resourceId := id
		resourceName := name
		if _, ok := names[name]; !ok {
			names[name] = struct{}{}
		} else {
			for i := 2; i > 0; i++ {
				thisName := resourceName + "_" + strconv.Itoa(i)
				if _, ok := names[thisName]; !ok {
					names[thisName] = struct{}{}
					resourceName = thisName
					break
				}
			}
		}
		resources = append(resources, terraformutils.NewResource(
			resourceId,
			resourceName,
			"davinci_variable",
			"davinci",
			map[string]string{
				"environment_id": variable.CompanyID,
				"name":           name,
				"context":        variable.Context,
			},
			VariableAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *VariableGenerator) InitResources() error {
	m := g.generateClient()
	list := map[string]davinci.Variable{}
	// params := davinci.Params{
	// 	Page:  "",
	// 	Limit: "",
	// }
	list, err := m.ReadVariables(&m.CompanyID, nil)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(list)
	return nil
}

func (g *VariableGenerator) PostConvertHook() error {
	//function to variablize resource environment_id
	if g.Args["abstract"].(string) == "true" {
		if err := g.updateEnvId("davinci_variable"); err != nil {
			return err
		}
	}
	return nil
}
