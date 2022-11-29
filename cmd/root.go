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

package cmd

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}
	cmd.AddCommand(newImportCmd())
	cmd.AddCommand(newPlanCmd())
	cmd.AddCommand(versionCmd)
	return cmd
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}

func providerImporterSubcommands() []func(options ImportOptions) *cobra.Command {
	return []func(options ImportOptions) *cobra.Command{
		// Major Cloud
<<<<<<< HEAD
		// 		newCmdGoogleImporter,
		// 		newCmdAwsImporter,
		// 		newCmdAzureImporter,
		// 		newCmdAliCloudImporter,
		// 		newCmdIbmImporter,
		// Cloud
		// 		newCmdDigitalOceanImporter,
		// 		newCmdEquinixMetalImporter,
		// 		newCmdHerokuImporter,
		// 		newCmdLaunchDarklyImporter,
		// 		newCmdLinodeImporter,
		// 		newCmdOpenStackImporter,
		// 		newCmdTencentCloudImporter,
		// 		newCmdVultrImporter,
		// 		newCmdYandexImporter,
		// Infrastructure Software
		// 		newCmdKubernetesImporter,
		// 		newCmdOctopusDeployImporter,
		// 		newCmdRabbitMQImporter,
		// Network
		// 		newCmdMyrasecImporter,
		// 		newCmdCloudflareImporter,
		// 		newCmdFastlyImporter,
		// 		newCmdNs1Importer,
		// 		newCmdPanosImporter,
		// VCS
		// 		newCmdAzureDevOpsImporter,
		// 		newCmdAzureADImporter,
		// 		newCmdGithubImporter,
		// 		newCmdGitLabImporter,
		// Monitoring & System Management
		// 		newCmdDatadogImporter,
		// 		newCmdNewRelicImporter,
		// 		newCmdMackerelImporter,
		// 		newCmdGrafanaImporter,
		// 		newCmdPagerDutyImporter,
		// 		newCmdOpsgenieImporter,
		// 		newCmdHoneycombioImporter,
		// Community
		// 		newCmdKeycloakImporter,
		// 		newCmdLogzioImporter,
		// 		newCmdCommercetoolsImporter,
		// 		newCmdMikrotikImporter,
		// 		newCmdXenorchestraImporter,
		// 		newCmdGmailfilterImporter,
		// 		newCmdVaultImporter,
		// 		newCmdOktaImporter,
		// 		newCmdAuth0Importer,
=======
// 		newCmdGoogleImporter,
// 		newCmdAwsImporter,
// 		newCmdAzureImporter,
// 		newCmdAliCloudImporter,
// 		newCmdIbmImporter,
		// Cloud
// 		newCmdDigitalOceanImporter,
// 		newCmdEquinixMetalImporter,
// 		newCmdHerokuImporter,
// 		newCmdLaunchDarklyImporter,
// 		newCmdLinodeImporter,
// 		newCmdOpenStackImporter,
// 		newCmdTencentCloudImporter,
// 		newCmdVultrImporter,
// 		newCmdYandexImporter,
		// Infrastructure Software
// 		newCmdKubernetesImporter,
// 		newCmdOctopusDeployImporter,
// 		newCmdRabbitMQImporter,
		// Network
// 		newCmdMyrasecImporter,
// 		newCmdCloudflareImporter,
// 		newCmdFastlyImporter,
// 		newCmdNs1Importer,
// 		newCmdPanosImporter,
		// VCS
// 		newCmdAzureDevOpsImporter,
// 		newCmdAzureADImporter,
// 		newCmdGithubImporter,
// 		newCmdGitLabImporter,
		// Monitoring & System Management
// 		newCmdDatadogImporter,
// 		newCmdNewRelicImporter,
// 		newCmdMackerelImporter,
// 		newCmdGrafanaImporter,
// 		newCmdPagerDutyImporter,
// 		newCmdOpsgenieImporter,
// 		newCmdHoneycombioImporter,
		// Community
// 		newCmdKeycloakImporter,
// 		newCmdLogzioImporter,
// 		newCmdCommercetoolsImporter,
// 		newCmdMikrotikImporter,
// 		newCmdXenorchestraImporter,
// 		newCmdGmailfilterImporter,
// 		newCmdVaultImporter,
// 		newCmdOktaImporter,
// 		newCmdAuth0Importer,
>>>>>>> e3ab3c2a26b8e9311e2c2c7f5a5a18d5f83d7944
		newCmdPingOneDavinciImporter,
	}
}

func providerGenerators() map[string]func() terraformutils.ProviderGenerator {
	list := make(map[string]func() terraformutils.ProviderGenerator)
	for _, providerGen := range []func() terraformutils.ProviderGenerator{
		// Major Cloud
<<<<<<< HEAD
		// 		newGoogleProvider,
		// 		newAWSProvider,
		// 		newAzureProvider,
		// 		newAliCloudProvider,
		// 		newIbmProvider,
		// Cloud
		// 		newDigitalOceanProvider,
		// 		newEquinixMetalProvider,
		// 		newFastlyProvider,
		// 		newHerokuProvider,
		// 		newLaunchDarklyProvider,
		// 		newLinodeProvider,
		// 		newNs1Provider,
		// 		newOpenStackProvider,
		// 		newTencentCloudProvider,
		// 		newVultrProvider,
		// Infrastructure Software
		// 		newKubernetesProvider,
		// 		newOctopusDeployProvider,
		// 		newRabbitMQProvider,
		// Network
		// 		newMyrasecProvider,
		// 		newCloudflareProvider,
		// VCS
		// 		newAzureDevOpsProvider,
		// 		newAzureADProvider,
		// 		newGitHubProvider,
		// 		newGitLabProvider,
		// Monitoring & System Management
		// newOpalProvider,
		// 		newDataDogProvider,
		// 		newNewRelicProvider,
		// 		newPagerDutyProvider,
		// 		newHoneycombioProvider,
		// Community
		// 		newKeycloakProvider,
		// 		newLogzioProvider,
		// 		newCommercetoolsProvider,
		// 		newMikrotikProvider,
		// 		newXenorchestraProvider,
		// 		newGmailfilterProvider,
		// 		newVaultProvider,
		// 		newOktaProvider,
		// 		newAuth0Provider,
=======
// 		newGoogleProvider,
// 		newAWSProvider,
// 		newAzureProvider,
// 		newAliCloudProvider,
// 		newIbmProvider,
		// Cloud
// 		newDigitalOceanProvider,
// 		newEquinixMetalProvider,
// 		newFastlyProvider,
// 		newHerokuProvider,
// 		newLaunchDarklyProvider,
// 		newLinodeProvider,
// 		newNs1Provider,
// 		newOpenStackProvider,
// 		newTencentCloudProvider,
// 		newVultrProvider,
		// Infrastructure Software
// 		newKubernetesProvider,
// 		newOctopusDeployProvider,
// 		newRabbitMQProvider,
		// Network
// 		newMyrasecProvider,
// 		newCloudflareProvider,
		// VCS
// 		newAzureDevOpsProvider,
// 		newAzureADProvider,
// 		newGitHubProvider,
// 		newGitLabProvider,
		// Monitoring & System Management
// 		newDataDogProvider,
// 		newNewRelicProvider,
// 		newPagerDutyProvider,
// 		newHoneycombioProvider,
		// Community
// 		newKeycloakProvider,
// 		newLogzioProvider,
// 		newCommercetoolsProvider,
// 		newMikrotikProvider,
// 		newXenorchestraProvider,
// 		newGmailfilterProvider,
// 		newVaultProvider,
// 		newOktaProvider,
// 		newAuth0Provider,
>>>>>>> e3ab3c2a26b8e9311e2c2c7f5a5a18d5f83d7944
		newPingOneDavinciProvider,
	} {
		list[providerGen().GetName()] = providerGen
	}
	return list
}
