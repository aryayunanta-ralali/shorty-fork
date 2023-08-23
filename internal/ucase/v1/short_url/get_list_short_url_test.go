package short_url

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/helper"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	mock_repositories "github.com/aryayunanta-ralali/shorty/mocks/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetListShortUrl(t *testing.T) {
	var (
		shortUrlIDInt = int64(1692391018317656062)
		shortUrlCode  = "test"
		userId        = "171d30a5-597b-4637-b2e1-654e75d567f4"

		endpoint = "/in/v1/short-urls/" + shortUrlCode + "/stats"
		ctx      = gomock.Any()
		errDummy = errors.New("dummy error")

		payload = presentations.GetListShortUrlPayload{}
	)

	type (
		output struct {
			appctx.Response
		}

		mockConfig struct {
			shortUrlRepo *mock_repositories.MockShortUrls
		}
	)

	testTable := []struct {
		testName       string
		configureInput func() *appctx.Data
		expected       output
		configureMock  func(mockConfig)
	}{
		{
			testName: "Test error get data from DB",
			configureInput: func() *appctx.Data {
				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(hReqBody))

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseInternalFailure,
				},
			},
			configureMock: func(mc mockConfig) {
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					Limit:  helper.LimitDefaultValue(payload.Limit),
					Offset: getOffset(payload.Page, helper.LimitDefaultValue(payload.Limit)),
				}).Return(nil, errDummy)
			},
		},
		{
			testName: "Test error data not found",
			configureInput: func() *appctx.Data {
				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(hReqBody))

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseDataNotFound,
				},
			},
			configureMock: func(mc mockConfig) {
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					Limit:  helper.LimitDefaultValue(payload.Limit),
					Offset: getOffset(payload.Page, helper.LimitDefaultValue(payload.Limit)),
				}).Return(nil, nil)
			},
		},
		{
			testName: "Test success",
			configureInput: func() *appctx.Data {
				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(hReqBody))

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseSuccess,
				},
			},
			configureMock: func(mc mockConfig) {
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					Limit:  helper.LimitDefaultValue(payload.Limit),
					Offset: getOffset(payload.Page, helper.LimitDefaultValue(payload.Limit)),
				}).Return([]entity.ShortUrls{{ID: shortUrlIDInt, UserID: userId}}, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mShortUrlRepo := mock_repositories.NewMockShortUrls(ctrl)
	mConfig := mockConfig{
		shortUrlRepo: mShortUrlRepo,
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			input := tt.configureInput()

			tt.configureMock(mConfig)

			ucase := NewGetListShortUrl(mShortUrlRepo)

			resp := ucase.Serve(input)

			assert.Equal(t, tt.expected.Name, resp.Name)
		})
	}
}
