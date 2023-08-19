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

type detailShortUrl struct {
	shortUrlRepo repositories.ShortUrls
}

func NewDetailShortUrl(shortUrlRepo repositories.ShortUrls) contract.UseCase {
	return &detailShortUrl{shortUrlRepo: shortUrlRepo}
}

// Serve
// API Contract: https://www.notion.so/ralalicom/
func (u *detailShortUrl) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx = tracer.SpanStartUseCase(data.Request.Context(), "Serve")

		lvState1       = consts.LogEventStateFetchDBData
		lfState1Status = "state_1_fetch_db_status"
		lfState1Data   = "state_1_fetch_db_data"

		lvState2       = consts.LogEventStateUpdateData
		lfState2Status = "state_3_update_data_to_db_status"

		err error

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameDetailShortUrl),
		}
	)

	defer tracer.SpanFinish(ctx)

	shortCode := mux.Vars(data.Request)["short_code"]

	/*---------------------------
	| STEP 1 : get data from db
	* --------------------------*/
	shortUrls, err := u.shortUrlRepo.FindBy(ctx, repositories.FindShortUrlsCriteria{
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

	if len(shortUrls) == 0 {
		lf = append(lf,
			logger.EventState(lvState1),
			logger.Any(lfState1Status, consts.LogStatusFailed),
		)
		logger.WarnWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBDataNotFound, entity.TableNameShortUrls), lf...)
		response.SetName(consts.ResponseDataNotFound)
		return
	}

	shortUrl := shortUrls[0]

	lf = append(lf,
		logger.Any(lfState1Status, consts.LogStatusSuccess),
		logger.Any(lfState1Data, util.DumpToString(shortUrl)),
	)

	/*---------------------------
	| STEP 2 : increment view count
	* --------------------------*/
	err = u.shortUrlRepo.IncrementViewCount(ctx, shortUrl.ID)
	if err != nil {
		tracer.SpanError(ctx, err)
		lf = append(lf,
			logger.EventState(lvState2),
			logger.Any(lfState2Status, consts.LogStatusFailed),
		)
		logger.ErrorWithContext(ctx, logger.SetMessageFormat(consts.LogMessageDBFailedToUpdate, entity.TableNameShortUrls, err), lf...)
		response.SetName(consts.ResponseInternalFailure)
		return
	}

	response.SetName(consts.ResponseSuccess).SetData(dto.TransformToDetailShortUrlResponse(shortUrl))
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameDetailShortUrl), lf...)
	return
}
