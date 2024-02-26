package redirect

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2 "github.com/otakakot/sample-go-openapi-gen/pkg/kiota/models"
)

// RedirectRequestBuilder builds and executes requests for operations under \redirect
type RedirectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RedirectRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RedirectRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRedirectRequestBuilderInternal instantiates a new RedirectRequestBuilder and sets the default values.
func NewRedirectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RedirectRequestBuilder) {
    m := &RedirectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/redirect", pathParameters),
    }
    return m
}
// NewRedirectRequestBuilder instantiates a new RedirectRequestBuilder and sets the default values.
func NewRedirectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RedirectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRedirectRequestBuilderInternal(urlParams, requestAdapter)
}
// Get redirect
func (m *RedirectRequestBuilder) Get(ctx context.Context, requestConfiguration *RedirectRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "4XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "5XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation redirect
func (m *RedirectRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RedirectRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *RedirectRequestBuilder) WithUrl(rawUrl string)(*RedirectRequestBuilder) {
    return NewRedirectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
