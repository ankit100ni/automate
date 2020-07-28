// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/tokens.proto

package v2

import (
	policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.iam.v2.Tokens/CreateToken", "iam:tokens", "iam:tokens:create", "POST", "/apis/iam/v2/tokens", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "value":
					return m.Value
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Tokens/GetToken", "iam:tokens:{id}", "iam:tokens:get", "GET", "/apis/iam/v2/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Tokens/UpdateToken", "iam:tokens:{id}", "iam:tokens:update", "PUT", "/apis/iam/v2/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Tokens/DeleteToken", "iam:tokens:{id}", "iam:tokens:delete", "DELETE", "/apis/iam/v2/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteTokenReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Tokens/ListTokens", "iam:tokens", "iam:tokens:list", "GET", "/apis/iam/v2/tokens", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
