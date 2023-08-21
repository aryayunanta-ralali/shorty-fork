package short_url

import (
	"github.com/aryayunanta-ralali/shorty/internal/dto"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type updateShortUrl struct {
	shortUrlRepo repositories.ShortUrls
}

func NewUpdateShortUrl(shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &updateShortUrl{shortUrlRepo: shortUrlRepo}
}

func (u *updateShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx    = tracer.SpanStartUseCase(data.Request.Context(), "Serve")
		userID = data.Request.Header.Get(consts.HeaderXUserID)

		lvState1       = consts.LogEventStateValidateRequestBody
		lfState1Status = "state_1_validate_request_status"

		lvState2       = consts.LogEventStateFetchDBData
		lfState2Status = "state_2_fetch_db_status"
		lfState2Data   = "state_2_fetch_db_data"

		lvState3       = consts.LogEventStateCheckUserID
		lfState3Status = "state_3_check_user_id_status"

		lvState4       = consts.LogEventStateUpdateData
		lfState4Status = "state_4_update_data_to_db_status"

		err       error
		shortCode = mux.Vars(data.Request)["short_code"]
		payload   presentations.UpdateShortUrlPayload

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameUpdateShortUrl),
		}
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

	/*---------------------------
	| STEP 2 : get data from db
	* --------------------------*/
	shortUrls, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
		ShortCode: shortCode,
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

	if len(shortUrls) == 0 {
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
		logger.Any(lfState2Data, util.DumpToString(shortUrls)),
	)

	/*---------------------------
	| STEP 3 : validate user id
	* --------------------------*/
	shortUrl := shortUrls[0]
	if shortUrl.UserID == "" || shortUrl.UserID != payload.UserID {
		lf = append(lf,
			logger.EventState(lvState3),
			logger.Any(lfState3Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageAuthenticationError), lf...)
		response.SetName(consts.ResponseAuthenticationFailure)
		return
	}

	/*---------------------------
	| STEP 4 : update data to db
	* --------------------------*/
	updateOrder := dto.TransformToUpdateEntity(shortUrl, payload)
	err = u.shortUrlRepo.Update(ctx, updateOrder)
	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState4),
			logger.Any(lfState4Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedToUpdate, entity.TableNameShortUrls, err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	response.SetName(consts.ResponseSuccess)
	response.Message = "Success update short url"
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameUpdateShortUrl), lf...)
	return
}

// validateRequestBody represent function to validate request payload
func (u *updateShortUrl) validateRequestBody(reqBody presentations.UpdateShortUrlPayload) (string, url.Values) {
	rules := govalidator.MapData{
		"user_id":    consts.RulesUserID,
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
