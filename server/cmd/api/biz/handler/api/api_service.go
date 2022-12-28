// Code generated by hertz generator.

package api

import (
	"context"
	"time"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/errno"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/server/cmd/api"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/global"
	models "github.com/CyanAsterisk/FreeCar/server/cmd/api/model"
	"github.com/CyanAsterisk/FreeCar/server/cmd/auth/kitex_gen/auth"
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/kitex_gen/car"
	"github.com/CyanAsterisk/FreeCar/server/cmd/profile/kitex_gen/profile"
	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/kitex_gen/trip"
	"github.com/CyanAsterisk/FreeCar/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
)

// Login .
// @router /v1/auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	// rpc to get accountID
	resp, err := global.AuthClient.Login(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	// create a JWT
	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID: resp.AccountID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "FreeCar",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		errno.SendResponse(c, errno.GenerateTokenFail, nil)
		return
	}
	// return token
	errno.SendResponse(c, errno.Success, api.LoginResponse{
		Token:     token,
		ExpiredAt: (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

// CreateCar .
// @router /v1/car [POST]
func CreateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req car.CreateCarRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.CarClient.CreateCar(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCar .
// @router /v1/car [GET]
func GetCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req car.GetCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.CarClient.GetCar(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetProfile .
// @router /v1/profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.GetProfileRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)
	resp, err := global.ProfileClient.GetProfile(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// SubmitProfile .
// @router /v1/profile [POST]
func SubmitProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.SubmitProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.SubmitProfile(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfile .
// @router /v1/profile [DELETE]
func ClearProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.ClearProfileRequest
	err = c.BindAndValidate(&req)
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.ClearProfile(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetProfilePhoto .
// @router /v1/profile/photo [GET]
func GetProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.GetProfilePhotoRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.GetProfilePhoto(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateProfilePhoto .
// @router /v1/profile/photo [POST]
func CreateProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.CreateProfilePhotoRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.CreateProfilePhoto(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CompleteProfilePhoto .
// @router /v1/profile/photo/complete [POST]
func CompleteProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.CompleteProfilePhotoRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.CompleteProfilePhoto(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfilePhoto .
// @router /v1/profile/photo [DELETE]
func ClearProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req profile.ClearProfilePhotoRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.ProfileClient.ClearProfilePhoto(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateTrip .
// @router /v1/trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req trip.CreateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.TripClient.CreateTrip(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrip .
// @router /v1/trip/:id [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req trip.GetTripRequest
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)
	req.Id = c.Param("id")

	resp, err := global.TripClient.GetTrip(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrips .
// @router /v1/trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req trip.GetTripsRequest
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)

	resp, err := global.TripClient.GetTrips(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// UpdateTrip .
// @router /v1/trip/:id [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req trip.UpdateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get("accountID")
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
	}
	req.AccountId = aid.(int64)
	req.Id = c.Param("id")

	resp, err := global.TripClient.UpdateTrip(ctx, &req)
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}
