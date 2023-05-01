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
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

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
	abstract            string
}

func (p PingOneDavinciProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"davinci_application": {
			"davinci_flow": []string{"policy.policy_flow.flow_id", "id"},
		},
		"davinci_flow": {
			"davinci_connection": []string{"connection_link.id", "id"},
			"davinci_flow":       []string{"subflow_link.id", "id"},
		},
	}
}

func (p PingOneDavinciProvider) GetResourceVariables() map[string][]terraformutils.VariableSet {
	return map[string][]terraformutils.VariableSet{
		"davinci_connection": {{Path: "property.value", Key: "property.name"}},
		"davinci_variable":   {{Path: "value", Key: "self_link"}},
	}
}

// reads a key=value config file and overwrites the values in the varsMap
func ReadPingOneConfig(filename string, varsMap map[string]*string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			log.Printf("invalid line in config file: %s", line)
		}
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		v = strings.Trim(v, "'")  // trim quotes
		v = strings.Trim(v, "\"") // trim quotes
		for key, value := range varsMap {
			if key == k {
				*value = v
			}
		}
	}
	return nil
}

func (p *PingOneDavinciProvider) Init(args []string) error {
	uName := args[0]
	if uName == "" {
		return errors.New("set PINGONE_USERNAME env var")
	}
	p.username = uName

	pWord := args[1]
	if pWord == "" {
		return errors.New("set PINGONE_PASSWORD env var")
	}
	p.password = pWord

	region := args[2]
	if region == "" {
		return errors.New("set PINGONE_REGION env var")
	}
	p.region = region

	environmentId := args[3]
	if region == "" {
		return errors.New("set PINGONE_ENVIRONMENT_ID env var")
	}
	p.environmentId = environmentId

	if targetEnvId := args[4]; targetEnvId != "" {
		p.targetEnvironmentId = targetEnvId
	}

	// override os env vars with args if needed.
	if len(args) > 4 && args[4] != "" {
		p.targetEnvironmentId = args[4]
	}

	//set abstract
	if len(args) > 5 {
		p.abstract = args[5]
	}

	return nil
}

func (p *PingOneDavinciProvider) GetName() string {
	return "davinci"
}
func (p *PingOneDavinciProvider) GetSource() string {
	return "pingidentity/davinci"
}

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
		"abstract":              p.abstract,
	})
	return nil
}

func (p *PingOneDavinciProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"davinci_connection":  &ConnectionGenerator{},
		"davinci_flow":        &FlowGenerator{},
		"davinci_application": &ApplicationGenerator{},
		"davinci_variable":    &VariableGenerator{},
	}
}

func (p PingOneDavinciProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
