package short_url

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	mock_repositories "github.com/aryayunanta-ralali/shorty/mocks/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertShortUrl(t *testing.T) {
	var (
		endpoint  = "/in/v1/short-urls"
		ctx       = gomock.Any()
		errDummy  = errors.New("dummy error")
		generator = func() int64 { return int64(1) }

		payload = presentations.InsertShortUrlPayload{
			URL:       "http://www.example.com/index.html",
			ShortCode: "test",
		}
	)

	type (
		output struct {
			appctx.Response
		}

		mockConfig struct {
			shortUrlRepo *mock_repositories.MockShortUrls
			generator    func() int64
		}
	)

	testTable := []struct {
		testName       string
		configureInput func() *appctx.Data
		expected       output
		configureMock  func(mockConfig)
	}{
		{
			testName: "Test error validation",
			configureInput: func() *appctx.Data {
				wrongShortCode := "test_"

				hReqBody, _ := json.Marshal(presentations.InsertShortUrlPayload{
					URL:       payload.URL,
					ShortCode: wrongShortCode,
				})
				hReq := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(hReqBody))

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseValidationFailure,
				},
			},
			configureMock: func(mc mockConfig) {
			},
		},
		{
			testName: "Test error short code existing",
			configureInput: func() *appctx.Data {

				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

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
					Limit:     1,
					ShortCode: payload.ShortCode,
				}).Return(nil, errDummy)
			},
		},
		{
			testName: "Test error insert data",
			configureInput: func() *appctx.Data {

				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

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
					Limit:     1,
					ShortCode: payload.ShortCode,
				}).Return([]entity.ShortUrls{}, nil)

				mc.shortUrlRepo.EXPECT().Insert(ctx, entity.ShortUrls{
					ID:        generator(),
					UserID:    payload.UserID,
					URL:       payload.URL,
					ShortCode: payload.ShortCode,
				}).Return(errDummy)
			},
		},
		{
			testName: "Test success",
			configureInput: func() *appctx.Data {

				hReqBody, _ := json.Marshal(payload)
				hReq := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

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
					Limit:     1,
					ShortCode: payload.ShortCode,
				}).Return([]entity.ShortUrls{}, nil)

				mc.shortUrlRepo.EXPECT().Insert(ctx, entity.ShortUrls{
					ID:        generator(),
					UserID:    payload.UserID,
					URL:       payload.URL,
					ShortCode: payload.ShortCode,
				}).Return(nil)
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

			ucase := NewInsertShortUrl(generator, mShortUrlRepo)

			resp := ucase.Serve(input)

			assert.Equal(t, tt.expected.Name, resp.Name)
		})
	}
}
