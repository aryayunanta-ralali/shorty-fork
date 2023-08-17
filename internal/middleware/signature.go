// Package middleware
package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/pkg/hash"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
)

// ValidateSignature header signature
func ValidateSignature(r *http.Request, conf *appctx.Config) string {

	var err error
	key := conf.App.ApiKey

	signature := r.Header.Get("Signature")

	save := r.Body

	save, r.Body, err = drainBody(r.Body)

	nameField := "request"

	if err != nil {
		return consts.ResponseSignatureFailure
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger.Warn(fmt.Sprintf("[middleware] cannot read request body , error %s", err.Error()), logger.Any(nameField, string(b)))
		return consts.ResponseSignatureFailure
	}

	buff := &bytes.Buffer{}
	err = json.Compact(buff, b)
	if err != nil {
		logger.Error(fmt.Sprintf("[middleware] cannot compat bytes , error %s", err.Error()), logger.Any(nameField, string(b)))
		return consts.ResponseSignatureFailure
	}

	if !hash.HmacComparator(buff.String(), signature, hash.SHA256(key)) {
		logger.Warn(fmt.Sprintf("[middleware] invalid signature %s", signature), logger.Any(nameField, string(b)))
		return consts.ResponseSignatureFailure
	}

	r.Body = save

	return consts.MiddlewarePassed
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
