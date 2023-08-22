package short_url

import (
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/dto"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
	"github.com/gorilla/mux"
)

type getStatShortUrl struct {
	shortUrlRepo repositories.ShortUrls
}

func NewGetStatShortUrl(shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &getStatShortUrl{shortUrlRepo: shortUrlRepo}
}

func (u *getStatShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx    = tracer.SpanStartUseCase(data.Request.Context(), "Serve")
		userID = data.Request.Header.Get(consts.HeaderXUserID)

		lvState1       = consts.LogEventStateFetchDBData
		lfState1Status = "state_1_fetch_db_status"
		lfState1Data   = "state_1_fetch_db_data"

		lvState2       = consts.LogEventStateCheckUserID
		lfState2Status = "state_2_check_user_id_status"

		err       error
		shortCode = mux.Vars(data.Request)["short_code"]

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameGetStatShortUrl),
		}
	)

	defer tracer.SpanFinish(ctx)

	/*---------------------------
	| STEP 1 : get data from db
	* --------------------------*/
	dbData, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
		Limit:     1,
		ShortCode: shortCode,
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

	if len(dbData) == 0 {
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
		logger.Any(lfState1Data, util.DumpToString(dbData)),
	)

	/*---------------------------
	| STEP 2 : validate user id
	* --------------------------*/
	if dbData[0].UserID == "" || dbData[0].UserID != userID {
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

	response.SetName(consts.ResponseSuccess).SetData(dto.TransformToGetStatShortUrlResponse(dbData[0]))
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameGetStatShortUrl), lf...)
	return
}
