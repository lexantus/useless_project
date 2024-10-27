# \DefaultAPI

All URIs are relative to *https://192.168.x.x/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivitiesMillisecondsGet**](DefaultAPI.md#ActivitiesMillisecondsGet) | **Get** /activities/{milliseconds} | Получение всех активностей, которые можно было бы сделать за присланное время
[**ActivityMillisecondsGet**](DefaultAPI.md#ActivityMillisecondsGet) | **Get** /activity/{milliseconds} | Получение случайной активности, которую можно успеть сделать за присланное время



## ActivitiesMillisecondsGet

> []Activity ActivitiesMillisecondsGet(ctx, milliseconds).Execute()

Получение всех активностей, которые можно было бы сделать за присланное время

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	milliseconds := int32(56) // int32 | Потраченное время в миллисекундах

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ActivitiesMillisecondsGet(context.Background(), milliseconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActivitiesMillisecondsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivitiesMillisecondsGet`: []Activity
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ActivitiesMillisecondsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milliseconds** | **int32** | Потраченное время в миллисекундах | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivitiesMillisecondsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Activity**](Activity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityMillisecondsGet

> Activity ActivityMillisecondsGet(ctx, milliseconds).Sex(sex).Execute()

Получение случайной активности, которую можно успеть сделать за присланное время

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	milliseconds := int32(56) // int32 | Присланное время в миллисекундах
	sex := "sex_example" // string | Пол клиента чтобы понимать как склонять слова в ответе (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ActivityMillisecondsGet(context.Background(), milliseconds).Sex(sex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActivityMillisecondsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivityMillisecondsGet`: Activity
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ActivityMillisecondsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milliseconds** | **int32** | Присланное время в миллисекундах | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityMillisecondsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sex** | **string** | Пол клиента чтобы понимать как склонять слова в ответе | 

### Return type

[**Activity**](Activity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

