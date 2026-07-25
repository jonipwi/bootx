package openaiadvisory

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/ethicalreview"
)

func TestReviewUsesBoundedResponsesRequest(t *testing.T) {
	var captured map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if got := request.Header.Get("Authorization"); got != "Bearer test-secret" {
			t.Errorf("authorization header = %q", got)
		}
		if err := json.NewDecoder(request.Body).Decode(&captured); err != nil {
			t.Fatal(err)
		}
		writer.Header().Set("x-request-id", "req_test")
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{
		  "id":"resp_test",
		  "status":"completed",
		  "model":"gpt-5.6-sol",
		  "output":[{"type":"message","content":[{"type":"output_text","text":"{\"summary\":\"Review completed without factual verification.\",\"statement_reviews\":[],\"findings\":[],\"missing_perspectives\":[],\"counterarguments\":[],\"questions_before_action\":[\"Have the declared sources been independently checked?\"],\"suggested_posture\":\"continue_human_review\",\"rewrite\":{\"provided\":true,\"draft\":\"A careful synthetic draft.\",\"change_summary\":[\"Added uncertainty language.\"]},\"limitations\":[\"No source was fetched or authenticated.\"]}"}]}]
		}`))
	}))
	defer server.Close()

	client, err := New("test-secret", DefaultModel, server.URL, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	request := testRequest()
	preflight, err := ethicalreview.Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	result, receipt, err := client.Review(context.Background(), request, preflight)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "completed" || result.Advice == nil {
		t.Fatalf("unexpected result: %+v", result)
	}
	if receipt.StoreRequested || receipt.ToolsEnabled || receipt.ExternalActionsEnabled || receipt.ApplicationPersistence {
		t.Fatalf("unsafe receipt: %+v", receipt)
	}
	if captured["store"] != false {
		t.Fatalf("store must be false: %#v", captured["store"])
	}
	if _, ok := captured["tools"]; ok {
		t.Fatal("tools must not be sent")
	}
	text := captured["text"].(map[string]any)
	format := text["format"].(map[string]any)
	if format["type"] != "json_schema" || format["strict"] != true {
		t.Fatalf("structured output not strict: %#v", format)
	}
	if captured["safety_identifier"] == request.UserID || !strings.HasPrefix(captured["safety_identifier"].(string), "bootx_") {
		t.Fatal("safety identifier must be privacy-preserving")
	}
	if strings.Contains(captured["input"].(string), request.UserID) {
		t.Fatal("local user_id must not be sent in model input")
	}
}

func TestReviewHandlesRefusal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"id":"resp_refusal","status":"completed","model":"gpt-5.6-sol","output":[{"type":"message","content":[{"type":"refusal","refusal":"Unable to assist with this content."}]}]}`))
	}))
	defer server.Close()
	client, _ := New("test-secret", DefaultModel, server.URL, server.Client())
	request := testRequest()
	preflight, _ := ethicalreview.Evaluate(request)
	result, _, err := client.Review(context.Background(), request, preflight)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "refused" || result.Refusal == "" || result.Advice != nil {
		t.Fatalf("unexpected refusal: %+v", result)
	}
}

func TestReviewSanitizesAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusUnauthorized)
		_, _ = writer.Write([]byte(`{"error":{"message":"test-secret was rejected","type":"invalid_request_error","code":"invalid_api_key"}}`))
	}))
	defer server.Close()
	client, _ := New("test-secret", DefaultModel, server.URL, server.Client())
	request := testRequest()
	preflight, _ := ethicalreview.Evaluate(request)
	_, _, err := client.Review(context.Background(), request, preflight)
	if err == nil || strings.Contains(err.Error(), "test-secret") || !strings.Contains(err.Error(), "invalid_api_key") {
		t.Fatalf("unexpected safe error: %v", err)
	}
}

func TestAdvisoryPostureCannotLowerDeterministicWarning(t *testing.T) {
	if advisoryPostureCompatible("W4_STOP", "continue_human_review") {
		t.Fatal("W4 must reject a weaker model posture")
	}
	if !advisoryPostureCompatible("W4_STOP", "seek_qualified_review") {
		t.Fatal("W4 must allow qualified review")
	}
	if advisoryPostureCompatible("W2_VERIFY", "continue_human_review") {
		t.Fatal("W2 must reject a weaker model posture")
	}
	if !advisoryPostureCompatible("W1_REVIEW", "continue_human_review") {
		t.Fatal("W1 may continue human review")
	}
}

func testRequest() ethicalreview.Request {
	return ethicalreview.Request{
		RequestID:                   "review-api-test",
		CapabilityID:                ethicalreview.CapabilityID,
		UserID:                      "test-user",
		CreatedAt:                   time.Date(2026, 7, 25, 12, 0, 0, 0, time.UTC),
		ContentType:                 "social_post",
		Purpose:                     "Synthetic API client test.",
		Audience:                    "Public",
		Context:                     "Synthetic and non-sensitive.",
		DraftText:                   "A careful synthetic draft.",
		Claims:                      []ethicalreview.EvidenceClaim{},
		PublicNonSensitiveConfirmed: true,
		RemoteProcessingConsent:     true,
		HumanAuthorityConfirmed:     true,
	}
}
