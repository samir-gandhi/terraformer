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
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type PingOneDavinciProvider struct { //nolint
	terraformutils.Provider
	username            string
	password            string
	region              string
	environmentId       string
	targetEnvironmentId string
}

func (p *PingOneDavinciProvider) Init(args []string) error {
	uName := os.Getenv("PINGONE_USERNAME")
	if uName == "" {
		return errors.New("set PINGONE_USERNAME env var")
	}
	p.username = uName

	pWord := os.Getenv("PINGONE_PASSWORD")
	if pWord == "" {
		return errors.New("set PINGONE_PASSWORD env var")
	}
	p.password = pWord

	region := os.Getenv("PINGONE_REGION")
	if region == "" {
		return errors.New("set PINGONE_REGION env var")
	}
	p.region = region

	environmentId := os.Getenv("PINGONE_ENVIRONMENT_ID")
	if region == "" {
		return errors.New("set PINGONE_ENVIRONMENT_ID env var")
	}
	p.environmentId = environmentId

	if targetEnvId := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID"); targetEnvId != "" {
		p.targetEnvironmentId = targetEnvId
	}

	if len(args) > 4 && args[4] != "" {
		p.targetEnvironmentId = args[4]
	}

	return nil
}

func (p *PingOneDavinciProvider) GetName() string {
	return "davinci"
}

<<<<<<< HEAD
=======
func (p *PingOneDavinciProvider) GetSource() string {
	return "pingidentity/davinci"
}

>>>>>>> e3ab3c2a26b8e9311e2c2c7f5a5a18d5f83d7944
func (p *PingOneDavinciProvider) GetConfig() cty.Value {

	return cty.ObjectVal(map[string]cty.Value{
		"username":       cty.StringVal(p.username),
		"password":       cty.StringVal(p.password),
		"region":         cty.StringVal(p.region),
		"environment_id": cty.StringVal(p.environmentId),
	})
}

func (p *PingOneDavinciProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username":              p.username,
		"password":              p.password,
		"region":                p.region,
		"environment_id":        p.environmentId,
		"target_environment_id": p.targetEnvironmentId,
	})
	return nil
}

func (p *PingOneDavinciProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"davinci_connections":  &ConnectionGenerator{},
		"davinci_flows":        &FlowGenerator{},
		"davinci_applications": &ApplicationGenerator{},
	}
}

func (p PingOneDavinciProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p PingOneDavinciProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
