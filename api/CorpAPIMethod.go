package api
/*
 * Generated CorpAPI methods
 * 2018-04-13 17:42:16.524965705 +0800 CST m=+0.002196426
 */

func (c *CorpAPI) OpenidToUserid(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["OPENID_TO_USERID"], args)
}
		
func (c *CorpAPI) BatchJobGetResult(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["BATCH_JOB_GET_RESULT"], args)
}
		
func (c *CorpAPI) GetJsapiTicket(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_JSAPI_TICKET"], args)
}
		
func (c *CorpAPI) UserUpdate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_UPDATE"], args)
}
		
func (c *CorpAPI) UserList(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_LIST"], args)
}
		
func (c *CorpAPI) DepartmentCreate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["DEPARTMENT_CREATE"], args)
}
		
func (c *CorpAPI) GetUserInfoByCode(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_USER_INFO_BY_CODE"], args)
}
		
func (c *CorpAPI) UpdateInvoiceStatus(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["UPDATE_INVOICE_STATUS"], args)
}
		
func (c *CorpAPI) UserAuthSuccess(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_AUTH_SUCCESS"], args)
}
		
func (c *CorpAPI) DepartmentUpdate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["DEPARTMENT_UPDATE"], args)
}
		
func (c *CorpAPI) TagCreate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_CREATE"], args)
}
		
func (c *CorpAPI) TagGetList(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_GET_LIST"], args)
}
		
func (c *CorpAPI) GetInvoiceInfo(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_INVOICE_INFO"], args)
}
		
func (c *CorpAPI) UserCreate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_CREATE"], args)
}
		
func (c *CorpAPI) UseridToOpenid(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USERID_TO_OPENID"], args)
}
		
func (c *CorpAPI) TagAddUser(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_ADD_USER"], args)
}
		
func (c *CorpAPI) AgentGet(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["AGENT_GET"], args)
}
		
func (c *CorpAPI) MenuGet(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["MENU_GET"], args)
}
		
func (c *CorpAPI) GetCheckinData(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_CHECKIN_DATA"], args)
}
		
func (c *CorpAPI) UserBatchDelete(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_BATCH_DELETE"], args)
}
		
func (c *CorpAPI) DepartmentList(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["DEPARTMENT_LIST"], args)
}
		
func (c *CorpAPI) TagGetUser(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_GET_USER"], args)
}
		
func (c *CorpAPI) TagDeleteUser(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_DELETE_USER"], args)
}
		
func (c *CorpAPI) AgentSet(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["AGENT_SET"], args)
}
		
func (c *CorpAPI) GetTicket(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_TICKET"], args)
}
		
func (c *CorpAPI) UserGet(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_GET"], args)
}
		
func (c *CorpAPI) DepartmentDelete(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["DEPARTMENT_DELETE"], args)
}
		
func (c *CorpAPI) UserSimpleList (args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_SIMPLE_LIST "], args)
}
		
func (c *CorpAPI) GetCheckinOption(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_CHECKIN_OPTION"], args)
}
		
func (c *CorpAPI) BatchInvite(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["BATCH_INVITE"], args)
}
		
func (c *CorpAPI) AgentGetList(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["AGENT_GET_LIST"], args)
}
		
func (c *CorpAPI) MenuDelete(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["MENU_DELETE"], args)
}
		
func (c *CorpAPI) MediaGet(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["MEDIA_GET"], args)
}
		
func (c *CorpAPI) UserDelete(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["USER_DELETE"], args)
}
		
func (c *CorpAPI) TagDelete(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_DELETE"], args)
}
		
func (c *CorpAPI) MessageSend(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["MESSAGE_SEND"], args)
}
		
func (c *CorpAPI) GetUserDetail(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_USER_DETAIL"], args)
}
		
func (c *CorpAPI) GetApprovalData(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["GET_APPROVAL_DATA"], args)
}
		
func (c *CorpAPI) BatchUpdateInvoiceStatus(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["BATCH_UPDATE_INVOICE_STATUS"], args)
}
		
func (c *CorpAPI) BatchGetInvoiceInfo(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["BATCH_GET_INVOICE_INFO"], args)
}
		
func (c *CorpAPI) TagUpdate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["TAG_UPDATE"], args)
}
		
func (c *CorpAPI) MenuCreate(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["MENU_CREATE"], args)
}
		