package ding

func (c *Client) Text() *Text {
	ptr := &Text{
		params: map[string]interface{}{
			"msgtype": "text",
		},
	}
	return ptr
}

func (c *Client) Image() *Image {
	ptr := &Image{
		params: map[string]interface{}{
			"msgtype": "image",
		},
	}
	return ptr
}

func (c *Client) Voice() *Voice {
	ptr := &Voice{
		params: map[string]interface{}{
			"msgtype": "voice",
		},
	}
	return ptr
}

func (c *Client) File() *File {
	ptr := &File{
		params: map[string]interface{}{
			"msgtype": "file",
		},
	}
	return ptr
}

func (c *Client) Link() *Link {
	ptr := &Link{
		params: map[string]interface{}{
			"msgtype": "link",
		},
	}
	return ptr
}

func (c *Client) Oa() *Oa {
	ptr := &Oa{
		params: map[string]interface{}{
			"msgtype": "oa",
		},
		oa:     map[string]interface{}{},
		oaBody: map[string]interface{}{},
	}
	return ptr
}

func (c *Client) Markdown() *Markdown {
	ptr := &Markdown{
		params: map[string]interface{}{
			"msgtype": "markdown",
		},
	}
	return ptr
}

func (c *Client) CardSingleJump() *CardSingleJump {
	ptr := &CardSingleJump{
		params: map[string]interface{}{
			"msgtype": "action_card",
		},
		inner: map[string]string{},
	}
	return ptr
}

func (c *Client) CardMultiJump() *CardMultiJump {
	ptr := &CardMultiJump{
		params: map[string]interface{}{
			"msgtype": "action_card",
		},
		inner: map[string]interface{}{},
	}
	return ptr
}
