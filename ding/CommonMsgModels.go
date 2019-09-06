package ding

type CommonMessage struct {
	params map[string]interface{}
}

func (c *CommonMessage) SetSender(userId string) *CommonMessage {
	c.params["sender"] = userId
	return c
}

func (c *CommonMessage) SetConversationId(cid string) *CommonMessage {
	c.params["cid"] = cid
	return c
}

func (c *CommonMessage) SetMsg (msg interface{}) *CommonMessage {
	c.params["msg"] = msg
	return c
}

func (c *CommonMessage) Make() *map[string]interface{} {
	return &c.params
}


