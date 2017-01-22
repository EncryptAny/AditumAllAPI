# ApApi

All URIs are relative to *http://066727de.ngrok.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apCreate**](ApApi.md#apCreate) | **POST** /ap | create ap
[**apDownvote**](ApApi.md#apDownvote) | **POST** /ap/{apID}/upvote | downvote ap
[**apUpvote**](ApApi.md#apUpvote) | **POST** /ap/{apID}/downvote | upvote ap


<a name="apCreate"></a>
# **apCreate**
> Ap apCreate(payload)

create ap

Create an AP

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.ApApi;


ApApi apiInstance = new ApApi();
NewAP payload = new NewAP(); // NewAP | Type for creating a new AP
try {
    Ap result = apiInstance.apCreate(payload);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ApApi#apCreate");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**NewAP**](NewAP.md)| Type for creating a new AP |

### Return type

[**Ap**](Ap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="apDownvote"></a>
# **apDownvote**
> apDownvote(apID)

downvote ap

Downvote a particular AP

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.ApApi;


ApApi apiInstance = new ApApi();
Integer apID = 56; // Integer | AP ID
try {
    apiInstance.apDownvote(apID);
} catch (ApiException e) {
    System.err.println("Exception when calling ApApi#apDownvote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apID** | **Integer**| AP ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="apUpvote"></a>
# **apUpvote**
> apUpvote(apID)

upvote ap

Upvote a particular AP

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.ApApi;


ApApi apiInstance = new ApApi();
Integer apID = 56; // Integer | AP ID
try {
    apiInstance.apUpvote(apID);
} catch (ApiException e) {
    System.err.println("Exception when calling ApApi#apUpvote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apID** | **Integer**| AP ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

