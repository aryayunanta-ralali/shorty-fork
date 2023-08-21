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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateShortUrl(t *testing.T) {
	var (
		shortUrlWithUserID = entity.ShortUrls{
			ID:        1692391018317656062,
			UserID:    "171d30a5-597b-4637-b2e1-654e75d567f4",
			URL:       "http://www.example.com/index.html",
			ShortCode: "test",
		}

		shortUrlWithoutUserID = entity.ShortUrls{
			ID:        1692391018317656062,
			URL:       "http://www.example.com/index.html",
			ShortCode: "test",
		}

		endpoint = "/in/v1/short-urls/" + shortUrlWithoutUserID.ShortCode
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
					"short_code": shortUrlWithoutUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})

				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
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
					ShortCode: shortUrlWithoutUserID.ShortCode,
				}).Return(nil, errDummy)
			},
		},
		{
			testName: "Test error data not found",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithoutUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					UserID:    "",
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
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
					ShortCode: shortUrlWithoutUserID.ShortCode,
					Limit:     1,
				}).Return(nil, nil)
			},
		},
		{
			testName: "Test error authentication without user id",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithoutUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq = mux.SetURLVars(hReq, vars)

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseAuthenticationFailure,
				},
			},
			configureMock: func(mc mockConfig) {
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: shortUrlWithoutUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithoutUserID}, nil)
			},
		},
		{
			testName: "Test error authentication with user id",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq.Header.Set(consts.HeaderXUserID, shortUrlWithUserID.UserID)
				hReq = mux.SetURLVars(hReq, vars)

				return &appctx.Data{
					Request:     hReq,
					ServiceType: consts.ServiceTypeHTTP,
				}
			},
			expected: output{
				appctx.Response{
					Name: consts.ResponseAuthenticationFailure,
				},
			},
			configureMock: func(mc mockConfig) {
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: shortUrlWithoutUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithoutUserID}, nil)
			},
		},
		{
			testName: "Test error check existing by short code",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq.Header.Set(consts.HeaderXUserID, shortUrlWithUserID.UserID)
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
					ShortCode: shortUrlWithUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithUserID}, nil)

				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: "test-new",
					Limit:     1,
				}).Return(nil, errDummy)
			},
		},
		{
			testName: "Test already exist check existing by short code",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq.Header.Set(consts.HeaderXUserID, shortUrlWithUserID.UserID)
				hReq = mux.SetURLVars(hReq, vars)

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
				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: shortUrlWithUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithUserID}, nil)

				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: "test-new",
					Limit:     1,
				}).Return([]entity.ShortUrls{{}}, nil)
			},
		},
		{
			testName: "Test error update data",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq.Header.Set(consts.HeaderXUserID, shortUrlWithUserID.UserID)
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
					ShortCode: shortUrlWithUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithUserID}, nil)

				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: "test-new",
					Limit:     1,
				}).Return([]entity.ShortUrls{}, nil)

				mc.shortUrlRepo.EXPECT().Update(ctx, entity.ShortUrls{
					ID:         shortUrlWithUserID.ID,
					UserID:     shortUrlWithUserID.UserID,
					URL:        shortUrlWithUserID.URL,
					ShortCode:  "test-new",
					VisitCount: shortUrlWithUserID.VisitCount,
					CreatedAt:  shortUrlWithUserID.CreatedAt,
					UpdatedAt:  shortUrlWithUserID.UpdatedAt,
					DeletedAt:  shortUrlWithUserID.DeletedAt,
				}).Return(errDummy)
			},
		},
		{
			testName: "Test success",
			configureInput: func() *appctx.Data {
				vars := map[string]string{
					"short_code": shortUrlWithoutUserID.ShortCode,
				}

				hReqBody, _ := json.Marshal(presentations.UpdateShortUrlPayload{
					ShortCode: "test-new",
				})
				hReq := httptest.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(hReqBody))
				hReq.Header.Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				hReq.Header.Set(consts.HeaderXUserID, shortUrlWithUserID.UserID)
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
					ShortCode: shortUrlWithUserID.ShortCode,
					Limit:     1,
				}).Return([]entity.ShortUrls{shortUrlWithUserID}, nil)

				mc.shortUrlRepo.EXPECT().FindBy(ctx, repositories.FindShortUrlsCriteria{
					ShortCode: "test-new",
					Limit:     1,
				}).Return([]entity.ShortUrls{}, nil)

				mc.shortUrlRepo.EXPECT().Update(ctx, entity.ShortUrls{
					ID:         shortUrlWithUserID.ID,
					UserID:     shortUrlWithUserID.UserID,
					URL:        shortUrlWithUserID.URL,
					ShortCode:  "test-new",
					VisitCount: shortUrlWithUserID.VisitCount,
					CreatedAt:  shortUrlWithUserID.CreatedAt,
					UpdatedAt:  shortUrlWithUserID.UpdatedAt,
					DeletedAt:  shortUrlWithUserID.DeletedAt,
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

			ucase := NewUpdateShortUrl(mShortUrlRepo)

			resp := ucase.Serve(input)

			assert.Equal(t, tt.expected.Name, resp.Name)
		})
	}
}
