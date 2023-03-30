// Code generated by hertz generator.

package profile

import (
	"context"

	hprofile "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/profile"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/config"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	kprofile "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/profile"
	"github.com/CyanAsterisk/FreeCar/server/shared/tools"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// DeleteProfile .
// @router /profile/admin/delete [DELETE]
func DeleteProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.DeleteProfileRequest
	resp := new(kprofile.DeleteProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.DeleteProfile(ctx, &kprofile.DeleteProfileRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// CheckProfile .
// @router /profile/admin/check [POST]
func CheckProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.CheckProfileRequest
	resp := new(kprofile.CheckProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.CheckProfile(ctx, &kprofile.CheckProfileRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// GetAllProfile .
// @router /profile/admin/all [GET]
func GetAllProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.GetAllProfileRequest
	resp := new(kprofile.GetAllProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.GetAllProfile(ctx, &kprofile.GetAllProfileRequest{})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// GetSomeProfile .
// @router /profile/admin/some [GET]
func GetSomeProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.GetSomeProfileRequest
	resp := new(kprofile.GetSomeProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.GetSomeProfile(ctx, &kprofile.GetSomeProfileRequest{})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// GetPendingProfile .
// @router /profile/admin/pending [GET]
func GetPendingProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.GetPendingProfileRequest
	resp := new(kprofile.GetPendingProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.GetPendingProfile(ctx, &kprofile.GetPendingProfileRequest{})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// GetProfile .
// @router /profile/mini/profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.GetProfileRequest
	resp := new(kprofile.GetProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.GetProfile(ctx, &kprofile.GetProfileRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// SubmitProfile .
// @router /profile/mini/profile [POST]
func SubmitProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.SubmitProfileRequest
	resp := new(kprofile.SubmitProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.SubmitProfile(ctx, &kprofile.SubmitProfileRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// ClearProfile .
// @router /profile/mini/profile [DELETE]
func ClearProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.ClearProfileRequest
	resp := new(kprofile.ClearProfileResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.ClearProfile(ctx, &kprofile.ClearProfileRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// CreateProfilePhoto .
// @router /profile/mini/photo [POST]
func CreateProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.CreateProfilePhotoRequest
	resp := new(kprofile.CreateProfilePhotoResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.CreateProfilePhoto(ctx, &kprofile.CreateProfilePhotoRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// ClearProfilePhoto .
// @router /profile/mini/photo [DELETE]
func ClearProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.ClearProfilePhotoRequest
	resp := new(kprofile.ClearProfilePhotoResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.ClearProfilePhoto(ctx, &kprofile.ClearProfilePhotoRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// GetProfilePhoto .
// @router /profile/mini/photo [GET]
func GetProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.GetProfilePhotoRequest
	resp := new(kprofile.GetProfilePhotoResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.GetProfilePhoto(ctx, &kprofile.GetProfilePhotoRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}
}

// CompleteProfilePhoto .
// @router /profile/mini/complete [GET]
func CompleteProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hprofile.CompleteProfilePhotoRequest
	resp := new(kprofile.CompleteProfilePhotoResponse)
	defer errno.NewSendResponse(c, resp)

	err = c.BindAndValidate(&req)
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		return
	}

	resp, err = config.GlobalProfileClient.CompleteProfilePhoto(ctx, &kprofile.CompleteProfilePhotoRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc profile service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
	}

}
