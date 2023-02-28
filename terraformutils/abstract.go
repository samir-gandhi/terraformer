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

package terraformutils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type VariableSet struct {
	Path string
	Key  string
}

func AbstractServices(importResources map[string][]Resource, isServicePath bool, resourceVariables map[string][]VariableSet) map[string][]Resource {
	for resource, variables := range resourceVariables {
		if resourceList, exist := importResources[resource]; exist {
			for k, resourceToMap := range resourceList {
				if resourceToMap.Variables == nil {
					resourceToMap.Variables = map[string]interface{}{}
				}
				for _, variableSet := range variables {
					if !isServicePath {
						absResource(resource, variableSet, resourceToMap, "local")
					} else {
						absResource(resource, variableSet, resourceToMap, strconv.Itoa(k))
					}
				}
				importResources[resource][k] = resourceToMap
			}
		}
	}
	return importResources
}

func absResource(resource string, variableSet VariableSet, resourceToMap Resource, k string) {
	resourceIdentifiers := WalkAndGet(variableSet.Path, resourceToMap.Item)
	for j, identifier := range resourceIdentifiers {
		if identifier == nil {
			continue
		}
		key := ""
		pathArr := strings.Split(variableSet.Path, ".")
		keyArr := strings.Split(variableSet.Key, ".")
		if variableSet.Key == "self_link" || variableSet.Key == "id" {
			key = pathArr[len(pathArr)-1]
		} else {
			//validate key base path matches Variableset.Path
			if strings.Join(pathArr[:len(pathArr)-1], ".") != strings.Join(keyArr[:len(keyArr)-1], ".") {
				log.Panicf("For VariableSet: %v, key: %v does not match path: %v", variableSet, strings.Join(keyArr[:len(keyArr)-1], "."), strings.Join(pathArr[:len(pathArr)-1], "."))
			}
			keys := WalkAndGet(variableSet.Key, resourceToMap.Item)
			key = keys[j].(string)
		}
		keyValue := resourceToMap.InstanceInfo.Type + "_" + resourceToMap.ResourceName + "_" + key
		linkValue := fmt.Sprintf("${var.%s}", keyValue)
		// store key value map of variables to export.
		resourceToMap.Variables[keyValue] = identifier.(string)
		WalkAndOverride(variableSet.Path, identifier.(string), linkValue, resourceToMap.Item)
	}
}
