package short_url

import (
	"github.com/aryayunanta-ralali/shorty/internal/dto"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
	"github.com/thedevsaddam/govalidator"
	"net/url"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
)

type insertShortUrl struct {
	generatorID  func() int64
	shortUrlRepo repositories.ShortUrls
}

func NewInsertShortUrl(generatorID func() int64, shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &insertShortUrl{
		generatorID:  generatorID,
		shortUrlRepo: shortUrlRepo,
	}
}

// Serve
// API Contract: https://www.notion.so/ralalicom/
func (u *insertShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx = tracer.SpanStartUseCase(data.Request.Context(), "Serve")

		lvState1       = consts.LogEventStateValidateRequestBody
		lfState1Status = "state_1_validate_request_status"

		lvState2       = consts.LogEventStateFetchDBData
		lfState2Status = "state_2_fetch_db_data_status"

		lvState3       = consts.LogEventStateInsertData
		lfState3Status = "state_3_insert_to_db_status"

		err     error
		payload presentations.InsertShortUrlPayload

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameInsertShortUrl),
		}

		userID = data.Request.Header.Get(consts.HeaderXUserID)
	)

	defer tracer.SpanFinish(ctx)

	/*------------------------------
	| STEP 1 : validate request body
	* -----------------------------*/
	err = data.Cast(&payload)
	if err != nil {
		lf = append(lf,
			logger.EventState(lvState1),
			logger.Any(lfState1Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageFailedToValidateRequestBody, err), lf...)
		response.SetName(consts.ResponseValidationFailure)
		return
	}

	payload.UserID = userID

	status, ev := u.validateRequestBody(payload)
	if status != consts.ResponseSuccess {
		lf = append(lf,
			logger.EventState(lvState1),
			logger.Any(lfState1Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageFailedToValidateRequestBody, ev), lf...)
		response.SetName(status)
		response.SetError(ev)
		return
	}

	lf = append(lf,
		logger.Any(lfState1Status, consts.LogStatusSuccess),
	)

	/*------------------------------
	| STEP 2: Check data existing by short code
	* -----------------------------*/
	shortUrl, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
		ShortCode: payload.ShortCode,
		Limit:     1,
	})
	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState2),
			logger.Any(lfState2Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedFetching, entity.TableNameShortUrls, err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	if len(shortUrl) != 0 {
		lf = append(lf,
			logger.Any(lfState2Status, consts.LogStatusFounded),
			logger.EventOutputHttp(response.GetCode(), response, util.DumpToString(response)),
		)
		logger.InfoWithContext(ctx,
			logger.SetMessageFormat("short url with code %s already exist", payload.ShortCode),
			lf...,
		)
		response.SetName(consts.ResponseValidationFailure)
		response.SetError(
			map[string][]string{"short_url": {"The given short code is already exists"}},
		)
		return
	}

	lf = append(lf,
		logger.Any(lfState2Status, consts.LogStatusSuccess),
	)

	/*---------------------------
	| STEP 3 : insert data from db
	* --------------------------*/
	dataInsert := dto.TransformToInsertEntity(payload, u.generatorID())
	err = u.shortUrlRepo.Insert(ctx, dataInsert)
	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState3),
			logger.Any(lfState3Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedToStore, "table_name", err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	response.SetName(consts.ResponseSuccess)
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameInsertShortUrl), lf...)
	return
}

// validateRequestBody represent function to validate request payload
func (u *insertShortUrl) validateRequestBody(reqBody presentations.InsertShortUrlPayload) (string, url.Values) {
	rules := govalidator.MapData{
		"user_id":    consts.RulesUserID,
		"url":        consts.RulesURL,
		"short_code": consts.RulesShortCodeURL,
	}

	options := govalidator.Options{
		Data:          &reqBody,
		Rules:         rules,
		TagIdentifier: "json",
	}

	v := govalidator.New(options)
	ev := v.ValidateStruct()
	if len(ev) != 0 {
		return consts.ResponseValidationFailure, ev
	}

	return consts.ResponseSuccess, nil
}
