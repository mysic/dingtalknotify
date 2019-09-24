package ding

func (c *Client) CommonMsg() *CommonMessage {
	ptr := &CommonMessage{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "message/send_to_conversation"
	c.requestHandler.method = "post"
	return ptr
}
