package ding

func (c *Client) CreateGroup() *CreateGroupMsg {
	ptr := &CreateGroupMsg{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "chat/create"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) SendMsg() *GroupMsg {
	ptr := &GroupMsg{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "chat/send"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) MsgReadList() *GroupMsgReadList {
	ptr := &GroupMsgReadList{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "chat/getReadList"
	c.requestHandler.method = "get"
	return ptr
}

func (c *Client) UpdateGroup() *UpdateGroup {
	ptr := &UpdateGroup{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "chat/send"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) GetGroupInfo() *GetGroupInfo {
	ptr := &GetGroupInfo{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "chat/get"
	c.requestHandler.method = "get"
	return ptr
}
