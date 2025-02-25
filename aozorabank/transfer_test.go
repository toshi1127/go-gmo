package aozorabank

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"

	"github.com/bxcodec/faker"
)

func TestGetTransferStatus(
	t *testing.T,
) {

	testcases := map[string]struct {
		request  *GetTransferStatusRequest
		rawQuery string
		expected *GetTransferStatusResponse
	}{
		"ok": {
			request: &GetTransferStatusRequest{
				AccessToken:             "access_token",
				AccountID:               "111111111111",
				QueryKeyClass:           QueryKeyClassTransferApplies,
				ApplyNo:                 "2018072902345678",
				DateFrom:                "2018-07-30",
				DateTo:                  "2018-08-10",
				NextItemKey:             "1234567890",
				RequestTransferStatuses: []*RequestTransferStatus{{TransferStatusApplying}},
				RequestTransferClass:    RequestTransferClassAll,
				RequestTransferTerm:     RequestTransferTermTransferDesignatedDate,
			},
			rawQuery: "accountId=111111111111&applyNo=2018072902345678&dateFrom=2018-07-30&dateTo=2018-08-10&nextItemKey=1234567890&queryKeyClass=1&requestTransferClass=1&requestTransferStatus=%5Bmap%5BrequestTransferStatus%3A2%5D%5D&requestTransferTerm=2",
			expected: &GetTransferStatusResponse{
				AcceptanceKeyClass: "acceptance_key_class",
				BaseDate:           "2023-08-01",
				BaseTime:           "00:00:01",
			},
		},
		"ok (required only)": {
			request: &GetTransferStatusRequest{
				AccessToken:   "access_token",
				AccountID:     "111111111111",
				QueryKeyClass: QueryKeyClassTransferApplies,
			},
			rawQuery: "accountId=111111111111&queryKeyClass=1",
			expected: &GetTransferStatusResponse{
				AcceptanceKeyClass: "acceptance_key_class",
				BaseDate:           "2023-08-01",
				BaseTime:           "00:00:02",
			},
		},
	}

	for title, tc := range testcases {
		t.Run(title, func(t *testing.T) {
			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				assert.Equal(t, tc.rawQuery, r.URL.RawQuery)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(APIHostTypeTest)
			result, err := cli.GetTransferStatus(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func TestTransferRequest(
	t *testing.T,
) {
	testcases := map[string]struct {
		request  *TransferRequestRequest
		expected *TransferRequestResponse
	}{
		"ok": {
			request: &TransferRequestRequest{
				AccessToken:             "access_token",
				IdempotencyKey:          "111111111111",
				AccountID:               "101011234567",
				RemitterName:            "ｼﾞ-ｴﾑｵ-ｼｮｳｼﾞ(ｶ",
				TransferDesignatedDate:  "2018-07-30",
				TransferDateHolidayCode: TransferDateHolidayCodeNextBusinessDay,
				TotalCount:              0,
				TotalAmount:             1000,
				ApplyComment:            "緊急で承認をお願いします",
				Transfers: []*Transfer{
					{
						ItemID:                "1",
						TransferAmount:        100,
						EDIInfo:               "ｾｲｷﾕｳｼﾖﾊﾞﾝｺﾞｳ1234",
						BeneficiaryBankCode:   "0398",
						BeneficiaryBankName:   "ｱｵｿﾞﾗ",
						BeneficiaryBranchCode: "111",
						BeneficiaryBranchName: "ﾎﾝﾃﾝ",
						AccountTypeCode:       AccountTypeCodeOrdinary,
						AccountNumber:         "1234567",
						BeneficiaryName:       "ｶ)ｱｵｿﾞﾗｻﾝｷﾞｮｳ",
					},
				},
			},
			expected: fakeData[TransferRequestResponse](),
		},
	}

	for title, tc := range testcases {
		t.Run(title, func(t *testing.T) {
			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(APIHostTypeTest)
			result, err := cli.TransferRequest(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func TestGetRequestResult(
	t *testing.T,
) {
	t.Parallel()

	testcases := map[string]struct {
		request  *GetRequestResultRequest
		expected *GetRequestResultResponse
	}{
		"ok": {
			request: &GetRequestResultRequest{
				AccessToken: "xxxxxxxxxxxx",
				AccountID:   "111111111111",
				ApplyNo:     "2018072902345678",
			},
			expected: fakeData[GetRequestResultResponse](),
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(APIHostTypeTest)
			result, err := cli.GetRequestResult(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func fakeData[T any]() *T {
	ret := new(T)
	if err := faker.FakeData(ret); err != nil {
		panic(err)
	}
	return ret
}
