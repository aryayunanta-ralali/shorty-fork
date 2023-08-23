package short_url

import (
	"github.com/aryayunanta-ralali/shorty/internal/dto"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/helper"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"net/url"

	"github.com/thedevsaddam/govalidator"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type getListShortUrl struct {
	shortUrlRepo repositories.ShortUrls
}

func NewGetListShortUrl(shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &getListShortUrl{shortUrlRepo: shortUrlRepo}
}

func (u *getListShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx = tracer.SpanStartUseCase(data.Request.Context(), "Serve")

		lvState1       = consts.LogEventStateValidateRequestBody
		lfState1Status = "state_1_validate_request_status"

		lvState2       = consts.LogEventStateFetchDBData
		lfState2Status = "state_2_fetch_db_status"
		lfState2Data   = "state_2_fetch_db_data"

		err     error
		payload presentations.GetListShortUrlPayload

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameGetListShortUrl),
		}
	)

	payload.Page = helper.PageDefaultValue(payload.Page)
	payload.Limit = helper.LimitDefaultValue(payload.Limit)

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

	/*---------------------------
	| STEP 2 : get data from db
	* --------------------------*/
	dbData, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
		Limit:  payload.Limit,
		Offset: getOffset(payload.Page, payload.Limit),
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

	if len(dbData) == 0 {
		lf = append(lf,
			logger.EventState(lvState2),
			logger.Any(lfState2Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBDataNotFound, entity.TableNameShortUrls), lf...)
		response.SetName(consts.ResponseDataNotFound)
		return
	}

	lf = append(lf,
		logger.Any(lfState2Status, consts.LogStatusSuccess),
		logger.Any(lfState2Data, util.DumpToString(dbData)),
	)

	response.SetName(consts.ResponseSuccess).SetData(dto.TransformToGetListShortUrlResponse(dbData))
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameGetListShortUrl), lf...)
	return
}

// validateRequestBody represent function to validate request payload
func (u *getListShortUrl) validateRequestBody(reqBody presentations.GetListShortUrlPayload) (string, url.Values) {
	// TO-DO: MOVE RULES BELOW TO VALIDATOR RULE CONSTANTS IF NOT ALREADY EXISTS
	RulesLimit := []string{"numeric", "numeric_between:1,50"}
	RulesPage := []string{"numeric", "numeric_between:1,"}

	rules := govalidator.MapData{
		"page":  RulesPage,
		"limit": RulesLimit,
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

func getOffset(page int64, limit int64) int64 {
	if page < 1 {
		return 0
	}

	return (page - 1) * limit
}
