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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var (
	FlowAllowEmptyValues = []string{}
)

type FlowGenerator struct {
	PingOneDavinciService
}

func (g FlowGenerator) createResources(flows []davinci.Flow) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, flow := range flows {
		resourceName := flow.FlowID
		resources = append(resources, terraformutils.NewResource(
			resourceName,
			resourceName+"_"+flow.Name,
			"davinci_flow",
			"davinci",
			map[string]string{
				"environment_id": flow.CompanyID,
			},
			FlowAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *FlowGenerator) InitResources() error {
	m := g.generateClient()
	list := []davinci.Flow{}

	params := davinci.Params{
		Page:  "",
		Limit: "",
	}
	l, err := m.ReadFlows(&m.CompanyID, &params)
	if err != nil {
		return err
	}
	list = append(list, l...)

	g.Resources = g.createResources(list)
	return nil
}
