package pets

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2 "github.com/otakakot/sample-go-openapi-gen/pkg/kiota/models"
)

// WithPet_ItemRequestBuilder builds and executes requests for operations under \pets\{pet_id}
type WithPet_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithPet_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithPet_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWithPet_ItemRequestBuilderInternal instantiates a new WithPet_ItemRequestBuilder and sets the default values.
func NewWithPet_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPet_ItemRequestBuilder) {
    m := &WithPet_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pets/{pet_id}", pathParameters),
    }
    return m
}
// NewWithPet_ItemRequestBuilder instantiates a new WithPet_ItemRequestBuilder and sets the default values.
func NewWithPet_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPet_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithPet_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get pet by id
func (m *WithPet_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithPet_ItemRequestBuilderGetRequestConfiguration)(i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "4XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "5XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreatePetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable), nil
}
// ToGetRequestInformation get pet by id
func (m *WithPet_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithPet_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WithPet_ItemRequestBuilder) WithUrl(rawUrl string)(*WithPet_ItemRequestBuilder) {
    return NewWithPet_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
