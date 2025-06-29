// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/topiccategoryservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type TopicCategoryServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedTopicCategoryServiceServer
}

func NewTopicCategoryServiceServer(svcCtx *svc.ServiceContext) *TopicCategoryServiceServer {
	return &TopicCategoryServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加话题分类
func (s *TopicCategoryServiceServer) AddTopicCategory(ctx context.Context, in *cmsclient.AddTopicCategoryReq) (*cmsclient.AddTopicCategoryResp, error) {
	l := topiccategoryservicelogic.NewAddTopicCategoryLogic(ctx, s.svcCtx)
	return l.AddTopicCategory(in)
}

// 删除话题分类
func (s *TopicCategoryServiceServer) DeleteTopicCategory(ctx context.Context, in *cmsclient.DeleteTopicCategoryReq) (*cmsclient.DeleteTopicCategoryResp, error) {
	l := topiccategoryservicelogic.NewDeleteTopicCategoryLogic(ctx, s.svcCtx)
	return l.DeleteTopicCategory(in)
}

// 更新话题分类
func (s *TopicCategoryServiceServer) UpdateTopicCategory(ctx context.Context, in *cmsclient.UpdateTopicCategoryReq) (*cmsclient.UpdateTopicCategoryResp, error) {
	l := topiccategoryservicelogic.NewUpdateTopicCategoryLogic(ctx, s.svcCtx)
	return l.UpdateTopicCategory(in)
}

// 更新话题分类状态
func (s *TopicCategoryServiceServer) UpdateTopicCategoryStatus(ctx context.Context, in *cmsclient.UpdateTopicCategoryStatusReq) (*cmsclient.UpdateTopicCategoryStatusResp, error) {
	l := topiccategoryservicelogic.NewUpdateTopicCategoryStatusLogic(ctx, s.svcCtx)
	return l.UpdateTopicCategoryStatus(in)
}

// 查询话题分类详情
func (s *TopicCategoryServiceServer) QueryTopicCategoryDetail(ctx context.Context, in *cmsclient.QueryTopicCategoryDetailReq) (*cmsclient.QueryTopicCategoryDetailResp, error) {
	l := topiccategoryservicelogic.NewQueryTopicCategoryDetailLogic(ctx, s.svcCtx)
	return l.QueryTopicCategoryDetail(in)
}

// 查询话题分类列表
func (s *TopicCategoryServiceServer) QueryTopicCategoryList(ctx context.Context, in *cmsclient.QueryTopicCategoryListReq) (*cmsclient.QueryTopicCategoryListResp, error) {
	l := topiccategoryservicelogic.NewQueryTopicCategoryListLogic(ctx, s.svcCtx)
	return l.QueryTopicCategoryList(in)
}
