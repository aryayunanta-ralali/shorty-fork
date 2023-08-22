package short_url

import (
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
	"github.com/gorilla/mux"
)

type deleteShortUrl struct {
	shortUrlRepo repositories.ShortUrls
}

func NewDeleteShortUrl(shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &deleteShortUrl{shortUrlRepo: shortUrlRepo}
}

func (u *deleteShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx    = tracer.SpanStartUseCase(data.Request.Context(), "Serve")
		userID = data.Request.Header.Get(consts.HeaderXUserID)

		lvState1       = consts.LogEventStateFetchDBData
		lfState1Status = "state_1_fetch_db_status"
		lfState1Data   = "state_1_fetch_db_data"

		lvState2       = consts.LogEventStateCheckUserID
		lfState2Status = "state_2_check_user_id_status"

		lvState3       = consts.LogEventStateDeleteData
		lfState3Status = "state_3_delete_data_to_db_status"

		err       error
		shortCode = mux.Vars(data.Request)["short_code"]

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameUpdateShortUrl),
		}
	)

	defer tracer.SpanFinish(ctx)

	/*---------------------------
	| STEP 1 : get data from db
	* --------------------------*/
	shortUrls, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
		ShortCode: shortCode,
		Limit:     1,
	})

	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState1),
			logger.Any(lfState1Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedFetching, entity.TableNameShortUrls, err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	if len(shortUrls) == 0 {
		lf = append(lf,
			logger.EventState(lvState1),
			logger.Any(lfState1Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBDataNotFound, entity.TableNameShortUrls), lf...)
		response.SetName(consts.ResponseDataNotFound)
		return
	}

	lf = append(lf,
		logger.Any(lfState1Status, consts.LogStatusSuccess),
		logger.Any(lfState1Data, util.DumpToString(shortUrls)),
	)

	/*---------------------------
	| STEP 3 : validate user id
	* --------------------------*/
	shortUrl := shortUrls[0]
	if shortUrl.UserID == "" || shortUrl.UserID != userID {
		lf = append(lf,
			logger.EventState(lvState2),
			logger.Any(lfState2Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageAuthenticationError), lf...)
		response.SetName(consts.ResponseAuthenticationFailure)
		return
	}

	lf = append(lf,
		logger.Any(lfState2Status, consts.LogStatusSuccess),
	)

	/*---------------------------
	| STEP 3 : delete data to db
	* --------------------------*/
	err = u.shortUrlRepo.Delete(ctx, shortUrl.ID)
	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState3),
			logger.Any(lfState3Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedToDelete, entity.TableNameShortUrls, err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	response.SetName(consts.ResponseSuccess)
	response.Message = "Success delete short url"

	lf = append(lf,
		logger.Any(lfState3Status, consts.LogStatusSuccess),
	)
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameDeleteShortUrl), lf...)
	return
}
