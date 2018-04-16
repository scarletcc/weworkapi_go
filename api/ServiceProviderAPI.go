package api

var SERVICE_PROVIDER_API_TYPE = map[string][]string{
	"GET_PROVIDER_TOKEN":       []string{"/cgi-bin/service/get_provider_token", "POST"},
	"GET_LOGIN_INFO":           []string{"/cgi-bin/service/get_login_info?access_token=PROVIDER_ACCESS_TOKEN", "POST"},
	"GET_REGISTER_CODE":        []string{"/cgi-bin/service/get_register_code?provider_access_token=PROVIDER_ACCESS_TOKEN", "POST"},
	"GET_REGISTER_INFO":        []string{"/cgi-bin/service/get_register_info?provider_access_token=PROVIDER_ACCESS_TOKEN", "POST"},
	"SET_AGENT_SCOPE":          []string{"/cgi-bin/agent/set_scope", "POST"}, //TODO
	"SET_CONTACT_SYNC_SUCCESS": []string{"/cgi-bin/sync/contact_sync_success", "GET"},
}

// TODO
type ServiceProviderApi struct {
}
