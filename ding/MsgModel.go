package ding

// text
type Text struct {
	params map[string]interface{}
}

func (t *Text) SetInfo(content string) *Text {
	t.params["text"] = map[string]string{
		"content" : content,
	}
	return t
}

func (t *Text) Make() *map[string]interface{}{
	return &t.params
}

// image
type Image struct {
	params map[string]interface{}
}


func (i *Image) SetInfo(mediaId string) *Image {
	i.params["image"] = map[string]string{
		"media_id": mediaId,
	}
	return i
}

func (i *Image) Make() *map[string]interface{}{
	return &i.params
}

// voice
type Voice struct {
	params map[string]interface{}
}


func (v *Voice) SetInfo(id, duration string) *Voice {
	v.params["voice"] = map[string]string{
		"media_id": id,
		"duration": duration,
	}
	return v
}

func (v *Voice) Make() *map[string]interface{}{
	return &v.params
}


// file
type File struct {
	params map[string]interface{}
}

func (f *File) SetInfo(mediaId string) *File {
	f.params["file"] = map[string]string{
		"media_id": mediaId,
	}
	return f
}

func (f *File) Make() *map[string]interface{}{
	return &f.params
}

// link
type Link struct {
	params map[string]interface{}
}

func (l *Link) SetInfo(messageUrl, picUrl, title, text string) *Link {
	l.params["link"] = map[string]string{
		"messageUrl": messageUrl,
		"picUrl": picUrl,
		"title": title,
		"text": text,
	}
	return l
}

func (l *Link) Make() *map[string]interface{}{
	return &l.params
}

// oa
type Oa struct {
	params map[string]interface{}
	oa map[string]interface{}
	oaBody map[string]interface{}
}

func (o *Oa) SetMsgUrl (url string) *Oa {
	o.oa["message_url"] = url
	return o
}

func (o *Oa) SetPcMsgUrl (url string) *Oa {
	o.oa["pc_message_url"] = url
	return o
}

func (o *Oa) SetHead(color, text string) *Oa {
	o.oa["head"] = map[string]string{
		"text": text,
		"bgcolor": color,
	}
	return o
}

func (o *Oa) SetBodyTitle(title string) *Oa {
	o.oaBody["title"] = title
	return o
}

func (o *Oa) SetBodyContent(content string) *Oa {
	o.oaBody["content"] = content
	return o
}

func (o *Oa) SetBodyImage(image string) *Oa {
	o.oaBody["image"] = image
	return o
}

func (o *Oa) SetBodyFileCount(fileCount string) *Oa {
	o.oaBody["file_count"] = fileCount
	return o
}

func (o *Oa) SetBodyAuthor (author string) *Oa {
	o.oaBody["author"] = author
	return o
}

func (o *Oa) SetBodyForm (form *[]map[string]string) *Oa {
	o.oaBody["form"] = form
	return o
}

func (o *Oa) SetBodyRich (rich *map[string]string) *Oa {
	o.oaBody["rich"] = rich
	return o
}

func (o *Oa) Make() *map[string]interface{}{
	o.oa["body"] = &o.oaBody
	o.params["oa"] = &o.oa
	return &o.params
}

// markdown
type Markdown struct {
	params map[string]interface{}
}

func (m *Markdown) SetInfo(title, text string) *Markdown {
	m.params["markdown"] = map[string]string{
		"title": title,
		"text": text,
	}
	return m
}

func (m *Markdown) Make() *map[string]interface{}{
	return &m.params
}

// action card single jump
type CardSingleJump struct{
	params map[string]interface{}
	inner map[string]string
}

func (sj *CardSingleJump) SetInfo (title, markdown string) *CardSingleJump {
	sj.inner["title"] =  title
	sj.inner["markdown"] = markdown
	return sj
}


func (sj *CardSingleJump) SetSingleJump(singleTitle, singleUrl string) *CardSingleJump {
	sj.inner["single_title"] = singleTitle
	sj.inner["single_url"] = singleUrl
	return sj
}


func (sj *CardSingleJump) Make() *map[string]interface{}{
	sj.params["action_card"] = sj.inner
	return &sj.params
}

//action card multi jump
type CardMultiJump struct{
	params map[string]interface{}
	inner map[string]interface{}
}

func (mj *CardMultiJump) SetInfo (title, markdown string) *CardMultiJump {
	mj.inner["title"] = title
	mj.inner["markdown"] = markdown
	return mj
}

func (mj *CardMultiJump) SetBtn(orient string, jsonList *[]map[string]string) *CardMultiJump {
	mj.inner["btn_orientation"] = orient
	mj.inner["btn_json_list"] = jsonList
	return mj
}

func (mj *CardMultiJump) Make() *map[string]interface{}{
	mj.params["action_card"] = mj.inner
	return &mj.params
}



