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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var (
	ConnectionAllowEmptyValues = []string{}
)

type ConnectionGenerator struct {
	PingOneDavinciService
}

func (g ConnectionGenerator) createResources(connections []davinci.Connection) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	names := map[string]struct{}{}
	for _, connection := range connections {
		resourceId := connection.ConnectionID
		resourceName := connection.Name
		if _, ok := names[connection.Name]; !ok {
			names[connection.Name] = struct{}{}
		} else {
			for i := 2; i > 0; i++ {
				thisName := connection.Name + "_" + strconv.Itoa(i)
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
			"davinci_connection",
			"davinci",
			map[string]string{
				"environment_id": connection.CompanyID,
			},
			ConnectionAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *ConnectionGenerator) InitResources() error {
	m := g.generateClient()
	list := []davinci.Connection{}

	l, err := m.ReadConnections(&m.CompanyID, nil)
	if err != nil {
		return err
	}

	list = append(list, l...)
	g.Resources = g.createResources(list)

	return nil
}

func (g *ConnectionGenerator) PostConvertHook() error {
	//function to variablize resource environment_id
	if g.Args["abstract"].(string) == "true" {
		if err := g.updateEnvId("davinci_connection"); err != nil {
			return err
		}
	}
	return nil
}
