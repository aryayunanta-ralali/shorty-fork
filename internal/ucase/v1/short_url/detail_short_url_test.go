package short_url

import (
	"errors"
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	mock_repositories "github.com/aryayunanta-ralali/shorty/mocks/repositories"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDetailShortUrl(t *testing.T) {
	var (
		shortUrlIDInt = int64(1692391018317656062)
		shortUrlCode  = "test"

		endpoint = "/in/v1/short-urls/" + shortUrlCode
		ctx      = gomock.Any()
		errDummy = errors.New("dummy error")
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
				vars := map[string]string{
					"short_code": shortUrlCode,
				}

				hReq := httptest.NewRequest(http.MethodGet, endpoint, nil)
				hReq = mux.SetURLVars(hReq, vars)

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
					ShortCode: shortUrlCode,
				}).Return(nil, errDummy)
			},
		},
		{
			testName: "Test error data not found",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlCode,
				}

				hReq := httptest.NewRequest(http.MethodGet, endpoint, nil)
				hReq = mux.SetURLVars(hReq, vars)

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
					ShortCode: shortUrlCode,
					Limit:     1,
				}).Return(nil, nil)
			},
		},
		{
			testName: "Test error increment view count data",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlCode,
				}

				hReq := httptest.NewRequest(http.MethodGet, endpoint, nil)
				hReq = mux.SetURLVars(hReq, vars)

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
					ShortCode: shortUrlCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{{ID: shortUrlIDInt}}, nil)

				mc.shortUrlRepo.EXPECT().IncrementViewCount(ctx, shortUrlIDInt).Return(errDummy)
			},
		},
		{
			testName: "Test success",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlCode,
				}

				hReq := httptest.NewRequest(http.MethodGet, endpoint, nil)
				hReq = mux.SetURLVars(hReq, vars)

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
					ShortCode: shortUrlCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{{ID: shortUrlIDInt}}, nil)

				mc.shortUrlRepo.EXPECT().IncrementViewCount(ctx, shortUrlIDInt).Return(nil)
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

			ucase := NewDetailShortUrl(mShortUrlRepo)

			resp := ucase.Serve(input)

			assert.Equal(t, tt.expected.Name, resp.Name)
		})
	}
}
