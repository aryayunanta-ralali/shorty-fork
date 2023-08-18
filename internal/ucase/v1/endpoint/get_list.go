package endpoint

import (
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"net/url"

	"github.com/thedevsaddam/govalidator"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
)

type getList struct {
}

func NewGetList() contract.UseCase {
	return &getList{}
}

// TO-DO: MOVE THIS CODE BELOW TO router.go
// getList := endpoint.NewGetList()
/*
	exV1.HandleFunc("/{route}", rtr.handle(
		handler.HttpRequest,
		getList,
		middleware.ValidateLanguage,
		middleware.ValidateAccountID,
		middleware.ValidateUserID,
		middleware.ValidateUserEmail,
	)).Methods(http.MethodGet)
*/

// Serve
// API Contract: https://www.notion.so/ralalicom/
func (u *getList) Serve(data *appctx.Data) (response appctx.Response) {
	var (
		ctx = tracer.SpanStartUseCase(data.Request.Context(), "Serve")

		lvState1       = consts.LogEventStateValidateRequestBody
		lfState1Status = "state_1_validate_request_status"

		err     error
		payload presentations.GetListPayload

		lf = []logger.Field{
			logger.EventName(consts.LogEventNameGetList),
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

	response.SetName(consts.ResponseSuccess).SetData(TransformToGetListResponse(consts.Endpoints))
	logger.InfoWithContext(ctx, logger.SetMessageFormat(consts.LogMessageSuccess, consts.LogEventNameGetList), lf...)
	return
}

// validateRequestBody represent function to validate request payload
func (u *getList) validateRequestBody(reqBody presentations.GetListPayload) (string, url.Values) {
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

// YOU CAN ALSO MOVE THIS FUNCTION TO DTO
// TO-DO: CHANGE INTERFACE TO ARRAY OF ENTITY REPOSITORY RESPONSE
func TransformToGetListResponse(data []string) []presentations.GetListResponse {
	var resp []presentations.GetListResponse

	for _, val := range data {
		currData := presentations.GetListResponse{
			URL: val,
		}

		resp = append(resp, currData)
	}

	return resp
}

func getOffset(page int64, limit int64) int64 {
	if page < 1 {
		return 0
	}

	return (page - 1) * limit
}
