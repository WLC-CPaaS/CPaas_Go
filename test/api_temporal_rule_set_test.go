/*
White Label Communications CPaas API Documentation

Testing TemporalRuleSetAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_TemporalRuleSetAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TemporalRuleSetAPIService V1AccountAccountIDTemporalrulesetGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetGet(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleSetAPIService V1AccountAccountIDTemporalrulesetPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetPost(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleSetAPIService V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleSetID string

		resp, httpRes, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete(context.Background(), accountID, temporalRuleSetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleSetAPIService V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleSetID string

		resp, httpRes, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet(context.Background(), accountID, temporalRuleSetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleSetAPIService V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleSetID string

		resp, httpRes, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut(context.Background(), accountID, temporalRuleSetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
