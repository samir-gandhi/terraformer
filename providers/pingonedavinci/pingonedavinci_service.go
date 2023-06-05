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
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type PingOneDavinciService struct { //nolint
	terraformutils.Service
}

func (s *PingOneDavinciService) generateClient() *davinci.APIClient {
	cInput := davinci.ClientInput{
		Username:        s.Args["username"].(string),
		Password:        s.Args["password"].(string),
		AccessToken:     s.Args["access_token"].(string),
		PingOneRegion:   s.Args["region"].(string),
		PingOneSSOEnvId: s.Args["environment_id"].(string),
	}
	apiClient, err := davinci.NewClient(&cInput)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if s.Args["target_environment_id"] != nil {
		apiClient.CompanyID = s.Args["target_environment_id"].(string)
	}

	return apiClient
}

func (s *PingOneDavinciService) updateEnvId(resourceType string) error {
	targetEnvironmentId := s.Args["environment_id"].(string)
	if s.Args["target_environment_id"] != nil {
		targetEnvironmentId = s.Args["target_environment_id"].(string)
	}
	for k, r := range s.Resources {
		thisResource := s.Resources[k]
		if r.InstanceInfo.Type == resourceType {
			if r.Item["environment_id"] != targetEnvironmentId {
				return fmt.Errorf("environment_id %q is not equal to target_environment_id %q", r.Item["environment_id"], targetEnvironmentId)
			}
			keyValue := "pingone_target_environment_id"
			linkValue := fmt.Sprintf("${var.%s}", keyValue)
			r.Item["environment_id"] = linkValue
			thisResource = r
		}
		s.Resources[k] = thisResource
	}
	return nil
}
