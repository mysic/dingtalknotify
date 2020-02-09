package main

import (
	"src/dingtalk/ding"
	"fmt"
)

func main() {

	host := "oapi.dingtalk.com"
	scheme := "https"
	agentid := ""
	appkey := ""
	appsecret := ""
	//实例化Client
	client := ding.NewClient(scheme, host)
	// 获取access token
	accessToken, errCode, errMsg, expiresIn, err := client.GetAccessToken(appkey, appsecret)
	if err != nil {
		fmt.Println(err.Error())
	}
	if errCode != 0 {
		fmt.Println(errMsg)
	}
	//将access token及有效期存入结构体中
	client.TokenContainer.SetExpiresIn(expiresIn).SetAccessToken(accessToken)


	/****
	生成消息
	 */
	// text类型 消息
	msg := client.Text().SetInfo("工作通知消息测试1").Make()

	// oa类型 消息
	/*
	form := &[]map[string]string{
		{"key":"姓名：", "value": "张三"},
		{"key":"年龄：", "value": "20"},
		{"key":"身高：", "value": "1.8米"},
		{"key":"体重：", "value": "130斤"},
		{"key":"学历：", "value": "本科"},
	}
	msg := client.Oa().
		SetMsgUrl("https://hello.xxx.com").
		SetPcMsgUrl("https://hello.xxx.com").
		SetHead("FF0000", "这里是头部大标题").// 普通回话发送才有效，企业回话会被替换成微应用名称
		SetBodyContent("这是body123内容").
		SetBodyAuthor("tony老师").
		SetBodyTitle("这是body标题").
		SetBodyForm(form).
		Make()
	*/


	/****
	 生成消息发送主体
	 */

	// 发送类型：工作通知
	  // 发送给所有人  3次/天
	requestBody := client.WorkNotice().
		SetToAllUser(true).
		SetAgentId(agentid).
		SetMsg(msg).
		Make()

	  // 发送给指定用户
	// user_id 12421853601031587  manager7887
	// 同一个应用下，同一个用户，相同用户 1次/天 ；
	// 同一应用，同一用户，发不同的消息，ISV方式50次/天 企业内部方式 500/天

	/*
	userIdList := &[]string{"manager7887"}
	requestBody := client.WorkNotice().
		SetUserIdList(userIdList).
		SetAgentId(agentid).
		SetMsg(msg).
		Make()
	 */

	// 发送类型：普通消息
	//cid := "14ac70d94e79377b88aa5fc75759fe84" //会话ID
	//senderId := "manager7887"
	//requestBody := client.CommonMsg().
	//	SetMsg(msg).
	//	SetConversationId(cid).
	//	SetSender(senderId).
	//	Make()

	// 发送类型：创建群聊会话
	//userIdList := &[]string{
	//	"manager7887",
	//	"12421853601031587",
	//}
	//requestBody := client.CreateGroup().
	//	SetName("群聊名称").
	//	SetOwner("manager7887").
	//	SetUserIdList(userIdList).
	//	Searchable().
	//	JoinValidation().
	//	OneManage().
	//	DenyMentionAll().
	//	ShowHistory().
	//	Make()



	//发送消息
	//*/
	resp, err := client.Send(requestBody)
	fmt.Println(resp)


	/****
	撤回工作通知消息
	 */

	/*
	taskId := int64(43616347648)
		requestBody := client.WorkNoticeRecall().
		SetAgentId(agentid).
		SetMsgTaskId(taskId).
		Make()
	resp, err := client.Send(requestBody)
	fmt.Println(resp)
	*/


	/****
	通知进度查询
	 */

	/*
	body := client.WorkNoticeProgressQuery().
		SetAgentId(agentid).
		SetTaskId(taskId).
		Make()
	resp, err := client.Send(body)
	fmt.Println(err)
	fmt.Println(resp)
	 */


	/****
	通知结果查询
	 */

	/*
	body := client.WorkNoticeResultQuery().
		SetAgentId(agentid).
		SetTaskId(taskId).
		Make()
	resp, err := client.Send(body)
	fmt.Println(err)
	fmt.Println(resp)
	 */

	/****
	发送普通消息
	*/

	/*
	requestBody := client.CommonMsg().
		SetMsg(msg).
		SetSender("manager7887").
		SetConversationId(""). //cid必须前端通过jsapi拉起回话取得cid
		Make()

	resp, err := client.Send(requestBody)
	println(resp)
	 */

	/****
	创建群消息
	*/

	/*
	userIdList := &[]string{"12421853601031587", "manager7887"}
	body := client.CreateGroup().
		SetUserIdList(userIdList).
		SetName("群会话测试").
		SetOwner("manager7887").
		ShowHistory().
		Searchable().
		DenyMentionAll().
		OneManage().
		Make()

	resp, err := client.Send(body)
	fmt.Println(resp)
	fmt.Println(err)
	*/



	/****
	发送群消息
	 */

	/*
	chatId := "chat610cc6308c64c7332fedea4c87b50df2"
	conversationTag := 2
	msg := client.Text().SetInfo("阿罗啊，阿修罗").Make()
	body := client.SendMsg().
		SetChatId(chatId).
		SetMsg(msg).
		Make()
	resp, err := client.Send(body)
	fmt.Println(resp)
	fmt.Println(err)

	 */


	/****
	查询群消息已读人员列表
	 */

	/*
	messageId := "msgMUyX1XXmSFvKAXkY+csUvw=="
	cursor := 1
	size := 10
	queryString := client.MsgReadList().
		SetQueryString(messageId, cursor, size).
		Make()
	resp, err := client.Send(queryString)
	fmt.Println(resp)
	fmt.Println(err)
	 */



	/****
	修改群会话信息
	 */

	/*
	body := client.UpdateGroup().
		SetChatId(chatId).
		SetOwner("12421853601031587").
		Make()
	resp, err := client.Send(body)
	fmt.Println(resp)
	fmt.Println(err)
	 */

	/****
	获取群信息
	 */

	/*
	queryString := client.GetGroupInfo().SetChatId(chatId).Make()
	resp, err := client.Send(queryString)
	fmt.Println(resp)
	fmt.Println(err)
	 */

}
