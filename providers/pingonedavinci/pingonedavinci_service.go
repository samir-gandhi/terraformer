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
	fmt.Printf(`apiClient.CompanyID is: %q`, apiClient.CompanyID)
	return apiClient
}
