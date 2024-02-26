package pets

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2 "github.com/otakakot/sample-go-openapi-gen/pkg/kiota/models"
)

// PetsRequestBuilder builds and executes requests for operations under \pets
type PetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PetsRequestBuilderGetQueryParameters list pets
type PetsRequestBuilderGetQueryParameters struct {
    // How many items to return at one time (max 100)
    Limit *int32 `uriparametername:"limit"`
}
// PetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PetsRequestBuilderGetQueryParameters
}
// PetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPet_id gets an item from the github.com/otakakot/sample-go-openapi-gen/pkg/kiota.pets.item collection
func (m *PetsRequestBuilder) ByPet_id(pet_id string)(*WithPet_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if pet_id != "" {
        urlTplParams["pet_id"] = pet_id
    }
    return NewWithPet_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPetsRequestBuilderInternal instantiates a new PetsRequestBuilder and sets the default values.
func NewPetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PetsRequestBuilder) {
    m := &PetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pets{?limit*}", pathParameters),
    }
    return m
}
// NewPetsRequestBuilder instantiates a new PetsRequestBuilder and sets the default values.
func NewPetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list pets
func (m *PetsRequestBuilder) Get(ctx context.Context, requestConfiguration *PetsRequestBuilderGetRequestConfiguration)([]i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "5XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreatePetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable)
        }
    }
    return val, nil
}
// Post create pet
func (m *PetsRequestBuilder) Post(ctx context.Context, body i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable, requestConfiguration *PetsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "4XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
        "5XX": i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.CreateErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation list pets
func (m *PetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create pet
func (m *PetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i0320fe562ee30d5bef5a60656400b17a6b4eea2f8a7072c52c06297d81332fc2.Petable, requestConfiguration *PetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PetsRequestBuilder) WithUrl(rawUrl string)(*PetsRequestBuilder) {
    return NewPetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
