package api

var SERVICE_CORP_API_TYPE = map[string][]string{
	"GET_CORP_TOKEN":         []string{"/cgi-bin/service/get_corp_token?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
	"GET_SUITE_TOKEN":        []string{"/cgi-bin/service/get_suite_token", "POST"},
	"GET_PRE_AUTH_CODE":      []string{"/cgi-bin/service/get_pre_auth_code?suite_access_token=SUITE_ACCESS_TOKEN", "GET"},
	"SET_SESSION_INFO":       []string{"/cgi-bin/service/set_session_info?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
	"GET_PERMANENT_CODE":     []string{"/cgi-bin/service/get_permanent_code?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
	"GET_AUTH_INFO":          []string{"/cgi-bin/service/get_auth_info?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
	"GET_ADMIN_LIST":         []string{"/cgi-bin/service/get_admin_list?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
	"GET_USER_INFO_BY_3RD":   []string{"/cgi-bin/service/getuserinfo3rd?suite_access_token=SUITE_ACCESS_TOKEN", "GET"},
	"GET_USER_DETAIL_BY_3RD": []string{"/cgi-bin/service/getuserdetail3rd?suite_access_token=SUITE_ACCESS_TOKEN", "POST"},
}

// TODO
type ServiceCorpAPI struct {
}
