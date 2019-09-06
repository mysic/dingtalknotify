package ding

import "strings"

type WorkNotice struct {
	params map[string]interface{}
}

func (w *WorkNotice) SetAgentId(agentId string) *WorkNotice {
	w.params["agent_id"] = agentId
	return w
}

func (w *WorkNotice) SetUserIdList(userIds *[]string) *WorkNotice {
	w.params["userid_list"] = strings.Join(*userIds, ",")
	return w
}

func (w *WorkNotice) SetDeptIdList(deptIds *[]string) *WorkNotice {
	w.params["dept_id_list"] = strings.Join(*deptIds, ",")
	return w
}

func (w *WorkNotice) SetToAllUser(isAll bool) *WorkNotice {
	w.params["to_all_user"] = isAll
	return w
}

func (w *WorkNotice) SetMsg(msg interface{}) *WorkNotice {
	w.params["msg"] = msg
	return w
}

func (w *WorkNotice) Make() *map[string]interface{} {
	return &w.params
}

type WorkNoticeProgressQuery struct {
	params map[string]interface{}
}

func (w *WorkNoticeProgressQuery) SetAgentId(agentId string) *WorkNoticeProgressQuery {
	w.params["agent_id"] = agentId
	return w
}

func (w *WorkNoticeProgressQuery) SetTaskId(taskId string) *WorkNoticeProgressQuery {
	w.params["task_id"] = taskId
	return w
}

func (w *WorkNoticeProgressQuery) Make() *map[string]interface{} {
	return &w.params
}

type WorkNoticeResultQuery struct {
	params map[string]interface{}
}

func (w *WorkNoticeResultQuery) SetAgentId(agentId string) *WorkNoticeResultQuery {
	w.params["agent_id"] = agentId
	return w
}

func (w *WorkNoticeResultQuery) SetTaskId(taskId string) *WorkNoticeResultQuery {
	w.params["task_id"] = taskId
	return w
}

func (w *WorkNoticeResultQuery) Make() *map[string]interface{} {
	return &w.params
}

type WorkNoticeRecall struct {
	params map[string]interface{}
}

func (w *WorkNoticeRecall) SetAgentId(agentId string) *WorkNoticeRecall {
	w.params["agent_id"] = agentId
	return w
}

func (w *WorkNoticeRecall) SetMsgTaskId(taskId string) *WorkNoticeRecall {
	w.params["msg_task_id"] = taskId
	return w
}

func (w *WorkNoticeRecall) Make() *map[string]interface{} {
	return &w.params
}
