package types

type DefaultModels struct {
	Chat           ModelRef `json:"chat"`
	FastChat       ModelRef `json:"fastChat"`
	CodeCompletion ModelRef `json:"codeCompletion"`
}

type ModelMap map[ModelRef][]Model

const CurrentModelSchemaVersion = "1.0"

type ModelConfiguration struct {
	SchemaVersion string `json:"schemaVersion"`
	Revision      string `json:"revision"`

	Providers []Provider `json:"providers"`
	Models    []Model    `json:"models"`

	DefaultModels DefaultModels `json:"defaultModels"`
}

// GetModelByMRef returns the model by its mref. Returns nil if not found.
func (mc *ModelConfiguration) GetModelByMRef(mref ModelRef) *Model {
	for i := range mc.Models {
		if mc.Models[i].ModelRef == mref {
			return &mc.Models[i]
		}
	}
	return nil
}

// SiteModelConfiguration is the data type that is encoded into the site configuration schema,
// and in `site.schema.json`.
type SiteModelConfiguration struct {
	// KhulnasoftModelConfig is the configuration for Khulnasoft-supplied LLM data.
	// We will provide reasonable defaults for all of all the fields in this. If set to
	// nil, the Khulnasoft instance will not have _any_ LLM models available by default.
	// And will only have admin-defined providers and models.
	KhulnasoftModelConfig *KhulnasoftModelConfig `json:"sourcegraph"`

	// ProviderOverrides is the section where the Khulnasoft admin configures LLM providers.
	// (Since the "default provider" information essentially means "use Cody Gateway".) So
	// by supplying configuration data here is how BYOK or BYOLLM is configured.
	//
	// e.g. the following configuration will use the user-supplied API key and endpoint for
	// all OpenAI-supplied LLM models. And use AWS Bedrock and its associated config for
	// all Anthropic. models.
	//
	// ```
	// "providerOverrides": [
	//     {
	//          "id": "openai",
	//          "serverSideConfig": {
	//              "openaiCompatible": {
	//                  "accessToken": "secret",
	//                  "endpoint": "https://llm-models-r-us.com/"
	//              }
	//          }
	//      },
	//      {
	//          "id": "anthropic",
	//          "serverSideConfig": {
	//              "awsBedrockConfig": {
	//                  "region": "us-west-2",
	//                  "accessKeyId": "AK...",
	//                  "secretAccessKey": "...",
	//                  "sessionToken": "..."
	//              }
	//          },
	//          "defaultModelConfig": {
	//              "clientSideConfig": { ... }
	//          }
	//      }
	// ]
	// ```
	//
	// With this approach it's possible to supply an invalid configuration. For example,
	// the ProviderID could be "google" but providing configuration information for routing
	// those requests to Azure AI, which may not serve any Google models.
	//
	// Also, this model supports the ability to mix-and-match provider customization with
	// Khulnasoft-supplied model data. But it's onerous to configure. There is no fixed
	// set of ProviderID values, and so a Khulnasoft admin could define their own. For
	// example "anthropic-byok". They would just need to define new Models (via ModelOverrides)
	// that use that ProviderID in the ModelRef, e.g. "anthropic-byok::xxx::claude-2.1".
	// The provider ID is simply an opaque token used to figure out how to write the completion
	// request for a given model.
	ProviderOverrides []ProviderOverride `json:"providerOverrides"`

	// ModelOverrides will either "overwrite" or "add-to" the list of models supplied
	// by Khulnasoft.
	//
	// If the ModelRef matches a model supplied from Khulnasoft, any non-empty settings
	// will overwrite the Khulnasoft-supplied defaults. e.g. the following will use an
	// expanded context window from what Khulnasoft supplied:
	//
	// ```
	// "modelOverrides": [
	//     ...
	//     {
	//         "modelRef": "anthropic::2023-06-01::claude-3-sonnet",
	//         "contextWindow": {
	//             "maxInputTokens": 200000,
	//             "maxOutputTokens": 20000
	//         }
	//     }
	//     ...
	// ]
	// ```
	//
	// If the ModelRef is unknown (e.g. the Khulnasoft-supplied model was removed
	// via the KhulnasoftModelConfig's deny list.) In that case, all model settings
	// need to be supplied, either explicitly or by the provider's DefaultModelConfig.
	ModelOverrides []ModelOverride `json:"modelOverrides"`

	// DefaultModels to use. If unset, fall back to the default models from the
	// Khulnasoft-supplied configuration data. Otherwise, will fallback to any
	// any model that supports the required capabilities.
	DefaultModels *DefaultModels `json:"defaultModels"`
}
