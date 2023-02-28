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
	"encoding/json"
	"fmt"
	"strconv"

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
	names := map[string]struct{}{}
	for _, flow := range flows {
		resourceId := flow.FlowID
		resourceName := flow.Name
		if _, ok := names[flow.Name]; !ok {
			names[flow.Name] = struct{}{}
		} else {
			for i := 2; i > 0; i++ {
				thisName := flow.Name + "_" + strconv.Itoa(i)
				if _, ok := names[thisName]; !ok {
					names[thisName] = struct{}{}
					resourceName = thisName
					break
				}
			}
		}

		// prepare additional attributes for terraform resource
		// addlAttrs := map[string]interface{}{}
		// addlAttrs["deploy"] = true
		// addlAttrs["connections"] = expandFlowConnections(flow)
		// sf := expandFlowSubflows(flow)
		// if sf != nil {
		// 	addlAttrs["subflows"] = sf
		// }
		configJson, err := json.MarshalIndent(flow, "", "  ")
		if err != nil {
			panic(fmt.Errorf("unable to marshal davinci flow to json %s: %v", flow.Name, err))
		}
		filename := fmt.Sprintf("flow_%s.json", flow.Name)
		resource := terraformutils.NewResource(
			resourceId,
			resourceName,
			"davinci_flow",
			"davinci",
			map[string]string{
				"environment_id": flow.CompanyID,
			},
			FlowAllowEmptyValues,
			map[string]interface{}{
				// "flow_json": "${file(" + "\"data/" + filename + "\")}",
				// "flow_json": fmt.Sprintf("${file(%q)}", "data/"+filename),
				// "flow_json": fmt.Sprintf(`${file("./data/%s")}`, filename),
				// "flow_json": fmt.Sprint("${file(" + "\"./data/" + filename + "\")}"),
				"flow_json": fmt.Sprintf("${file(\"data/%s\")}", filename),
			},
			// addlAttrs,
		)
		resource.DataFiles = map[string][]byte{
			filename: configJson,
		}

		resources = append(resources, resource)
	}
	return resources
}

// stolen from dv terraform provider - should be exported.
func expandFlowSubflows(flow davinci.Flow) []map[string]string {
	var subflows []map[string]string
	nodes := flow.GraphData.Elements.Nodes
	for _, node := range nodes {
		if node.Data.ConnectorID == "flowConnector" && (node.Data.CapabilityName == "startSubFlow" || node.Data.CapabilityName == "startUiSubFlow") {
			sfProps, err := expandSubFlowProps(node.Data.Properties)
			if err != nil {
				panic(err)
			}
			subflow := map[string]string{
				"subflow_id":   sfProps.SubFlowID.Value.Value,
				"subflow_name": sfProps.SubFlowID.Value.Label,
			}
			if !containsObj(subflows, subflow) && subflow["subflow_id"] != "" {
				subflows = append(subflows, subflow)
			}
		}
	}
	if len(subflows) > 0 {
		return subflows
	}
	return nil
}

func expandSubFlowProps(subflowProps map[string]interface{}) (*davinci.SubFlowProperties, error) {

	sfp := subflowProps["subFlowId"].(map[string]interface{})
	sfpVal := sfp["value"].(map[string]interface{})
	sfId := davinci.SubFlowID{
		Value: davinci.SubFlowValue{
			Value: sfpVal["value"].(string),
			Label: sfpVal["label"].(string),
		},
	}
	subflowVersionId := subflowProps["subFlowVersionId"].(map[string]interface{})
	sfv := davinci.SubFlowVersionID{
		Value: subflowVersionId["value"].(string),
	}
	if sfId.Value.Value == "" || sfv.Value == "" {
		return nil, fmt.Errorf("Error: subflow value or versionId is empty")
	}
	subflow := davinci.SubFlowProperties{
		SubFlowID:        sfId,
		SubFlowVersionID: sfv,
	}
	return &subflow, nil
}

func containsObj(items []map[string]string, subitem map[string]string) bool {
	for _, item := range items {
		if item["connection_id"] == subitem["connection_id"] {
			return true
		}
	}
	return false
}

func (g *FlowGenerator) InitResources() error {
	m := g.generateClient()

	params := davinci.Params{
		Page:  "",
		Limit: "",
	}
	list, err := m.ReadFlows(&m.CompanyID, &params)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(list)
	return nil
}

func (g *FlowGenerator) PostConvertHook() error {
	//function to variablize resource environment_id
	if g.Args["abstract"].(string) == "true" {
		targetEnvironmentId := g.Args["environment_id"].(string)
		if g.Args["target_environment_id"] != nil {
			targetEnvironmentId = g.Args["target_environment_id"].(string)
		}
		for k, r := range g.Resources {
			thisResource := g.Resources[k]
			if r.InstanceInfo.Type == "davinci_flow" {
				if r.Item["environment_id"] != targetEnvironmentId {
					return fmt.Errorf("environment_id %q is not equal to target_environment_id %q", r.Item["environment_id"], targetEnvironmentId)
				}
				keyValue := "pingone_target_environment_id"
				linkValue := fmt.Sprintf("${var.%s}", keyValue)
				r.Item["environment_id"] = linkValue
				thisResource = r
			}
			g.Resources[k] = thisResource
		}

		if g.Resources[0].Variables == nil {
			g.Resources[0].Variables = map[string]interface{}{}
		}
		g.Resources[0].Variables["pingone_target_environment_id"] = targetEnvironmentId
	}
	return nil
}
