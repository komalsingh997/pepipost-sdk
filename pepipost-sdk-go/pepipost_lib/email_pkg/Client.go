/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package email_pkg

import (
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/apihelper_pkg"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/configuration_pkg"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/models_pkg"
)

/*
 * Client structure as interface implementation
 */
type EMAIL_IMPL struct {
	config configuration_pkg.CONFIGURATION
}

/**
 * This Endpoint sends emails with the credentials passed.
 * @param    *string                      apiKey      parameter: Optional
 * @param    *models_pkg.EmailBody        body        parameter: Optional
 * @return	Returns the *models_pkg.SendEmailResponse response from the API call
 */
func (me *EMAIL_IMPL) CreateSendEmail(
	apiKey *string,
	body *models_pkg.EmailBody) (*models_pkg.SendEmailResponse, error) {
	//the base uri for api requests
	_queryBuilder := configuration_pkg.BASEURI

	//prepare query string for API call
	//	_queryBuilder = _queryBuilder + "/v2/sendEmail"
	_queryBuilder = _queryBuilder + "/v4/sendEmail"
	//variable to hold errors
	var err error = nil
	//validate and preprocess url
	_queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}

	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":   "pepi-sdk-go v2",
		"accept":       "application/json",
		"content-type": "application/json; charset=utf-8",
		"api_key":      apihelper_pkg.ToString(apiKey, ""),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request)
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 405 {
		err = apihelper_pkg.NewAPIError("Method not allowed", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models_pkg.SendEmailResponse = &models_pkg.SendEmailResponse{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}
