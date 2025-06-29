package coupon_type

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponTypeLogic 更新优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateCouponTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCouponTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponTypeLogic {
	return &UpdateCouponTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCouponType 更新优惠券类型
func (l *UpdateCouponTypeLogic) UpdateCouponType(req *types.UpdateCouponTypeReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CouponTypeService.UpdateCouponType(l.ctx, &smsclient.UpdateCouponTypeReq{
		Id:           req.Id,           // 主键ID
		Name:         req.Name,         // 类型名称
		Code:         req.Code,         // 类型编码
		Description:  req.Description,  // 描述
		DiscountType: req.DiscountType, // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
		Status:       req.Status,       // 是否启用
		Sort:         req.Sort,         // 排序
		UpdateBy:     userId,           // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券类型失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新优惠券类型成功",
	}, nil
}
