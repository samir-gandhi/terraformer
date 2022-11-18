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
	ApplicationAllowEmptyValues = []string{}
)

type ApplicationGenerator struct {
	PingOneDavinciService
}

func (g ApplicationGenerator) createResources(applications []davinci.App) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, application := range applications {
		resourceName := application.AppID
		resources = append(resources, terraformutils.NewResource(
			resourceName,
			resourceName+"_"+application.Name,
			"davinci_application",
			"davinci",
			map[string]string{
				"environment_id": application.CompanyID,
			},
			ApplicationAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *ApplicationGenerator) InitResources() error {
	m := g.generateClient()
	list := []davinci.App{}
	params := davinci.Params{
		Page:  "",
		Limit: "",
	}
	l, err := m.ReadApplications(&m.CompanyID, &params)
	if err != nil {
		return err
	}
	list = append(list, l...)

	g.Resources = g.createResources(list)
	return nil
}
