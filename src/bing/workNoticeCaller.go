package ding

func (c *Client) WorkNotice() *WorkNotice {
	ptr := &WorkNotice{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "topapi/message/corpconversation/asyncsend_v2"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) WorkNoticeProgressQuery() *WorkNoticeProgressQuery {
	ptr := &WorkNoticeProgressQuery{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "topapi/message/corpconversation/getsendprogress"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) WorkNoticeResultQuery() *WorkNoticeResultQuery {
	ptr := &WorkNoticeResultQuery{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "topapi/message/corpconversation/getsendresult"
	c.requestHandler.method = "post"
	return ptr
}

func (c *Client) WorkNoticeRecall() *WorkNoticeRecall {
	ptr := &WorkNoticeRecall{
		make(map[string]interface{}),
	}
	c.requestHandler.path = "topapi/message/corpconversation/recall"
	c.requestHandler.method = "post"
	return ptr
}
