package ding

// CreateGroupMsg
type CreateGroupMsg struct {
	params map[string]interface{}
}

func (c *CreateGroupMsg) SetName(name string) *CreateGroupMsg {
	c.params["name"] = name
	return c
}

func (c *CreateGroupMsg) SetOwner(userId string) *CreateGroupMsg {
	c.params["owner"] = userId
	return c
}

func (c *CreateGroupMsg) SetUserIdList(userIdList *[]string) *CreateGroupMsg {
	c.params["useridlist"] = userIdList
	return c
}

func (c *CreateGroupMsg) ShowHistory() *CreateGroupMsg {
	c.params["showHistoryType"] = 1
	return c
}

func (c *CreateGroupMsg) Searchable() *CreateGroupMsg {
	c.params["searchable"] = 1
	return c
}

func (c *CreateGroupMsg) JoinValidation() *CreateGroupMsg {
	c.params["validationType"] = 1
	return c
}

func (c *CreateGroupMsg) DenyMentionAll() *CreateGroupMsg {
	c.params["mentionAllAuthority"] = 1
	return c
}

func (c *CreateGroupMsg) ChatAllBanned() *CreateGroupMsg {
	c.params["chatBannedType"] = 1
	return c
}

func (c *CreateGroupMsg) OneManage() *CreateGroupMsg {
	c.params["managementType"] = 1
	return c
}

func (g *CreateGroupMsg) Make() *map[string]interface{} {
	return &g.params
}

// GroupMsg
type GroupMsg struct {
	params map[string]interface{}
}

func (g *GroupMsg) SetChatId(chatId string) *GroupMsg {
	g.params["chatid"] = chatId
	return g
}

func (g *GroupMsg) SetMsg(msg interface{}) *GroupMsg {
	g.params["msg"] = msg
	return g
}

func (g *GroupMsg) Make() *map[string]interface{} {
	return &g.params
}

// GroupMsgReadList
type GroupMsgReadList struct {
	params map[string]interface{}
}

func (g *GroupMsgReadList) SetQueryString(messageId string, cursor, size int) *GroupMsgReadList {
	g.params["messageId"] = messageId
	g.params["cursor"] = cursor
	g.params["size"] = size
	return g
}

func (g *GroupMsgReadList) Make() *map[string]interface{} {
	return &g.params
}

// UpdateGroup
type UpdateGroup struct {
	params map[string]interface{}
}

func (u *UpdateGroup) SetChatId(chatId string) *UpdateGroup {
	u.params["chatid"] = chatId
	return u
}

func (u *UpdateGroup) SetName(name string) *UpdateGroup {
	u.params["name"] = name
	return u
}

func (u *UpdateGroup) SetOwner(userId string) *UpdateGroup {
	u.params["owner"] = userId
	return u
}

func (u *UpdateGroup) AddUsers(userIdList []string) *UpdateGroup {
	u.params["add_useridlist"] = userIdList
	return u
}

func (u *UpdateGroup) DelUsers(userIdList []string) *UpdateGroup {
	u.params["del_useridlist"] = userIdList
	return u
}

func (u *UpdateGroup) SetIcon(mediaId string) *UpdateGroup {
	u.params["icon"] = mediaId
	return u
}

func (u *UpdateGroup) ShowHistory() *UpdateGroup {
	u.params["showHistoryType"] = 1
	return u
}

func (u *UpdateGroup) ShowHistoryCancel() *UpdateGroup {
	u.params["showHistoryType"] = 0
	return u
}

func (u *UpdateGroup) Searchable() *UpdateGroup {
	u.params["searchable"] = 1
	return u
}

func (u *UpdateGroup) SearchableCancel() *UpdateGroup {
	u.params["searchable"] = 0
	return u
}

func (u *UpdateGroup) JoinValidation() *UpdateGroup {
	u.params["validationType"] = 1
	return u
}

func (u *UpdateGroup) JoinValidationCancel() *UpdateGroup {
	u.params["validationType"] = 0
	return u
}

func (u *UpdateGroup) DenyMentionAll() *UpdateGroup {
	u.params["mentionAllAuthority"] = 1
	return u
}

func (u *UpdateGroup) DenyMentionAllCancel() *UpdateGroup {
	u.params["mentionAllAuthority"] = 0
	return u
}

func (u *UpdateGroup) ChatAllBanned() *UpdateGroup {
	u.params["chatBannedType"] = 1
	return u
}

func (u *UpdateGroup) ChatAllBannedCancel() *UpdateGroup {
	u.params["chatBannedType"] = 0
	return u
}

func (u *UpdateGroup) OneManage() *UpdateGroup {
	u.params["managementType"] = 1
	return u
}

func (u *UpdateGroup) OneMangeCancel() *UpdateGroup {
	u.params["managementType"] = 0
	return u
}

func (u *UpdateGroup) Make() *map[string]interface{} {
	return &u.params
}

// GetGroupInfo
type GetGroupInfo struct {
	params map[string]interface{}
}

func (g *GetGroupInfo) SetChatId(chatId string) *GetGroupInfo {
	g.params["chatid"] = chatId
	return g
}

func (g *GetGroupInfo) Make() *map[string]interface{} {
	return &g.params
}
