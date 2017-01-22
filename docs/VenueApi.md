# VenueApi

All URIs are relative to *http://066727de.ngrok.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**venueAll**](VenueApi.md#venueAll) | **GET** /venue | all venue
[**venueDownvoteAI**](VenueApi.md#venueDownvoteAI) | **POST** /venue/{venueID}/{aiType}/downvote | downvoteAI venue
[**venueNew**](VenueApi.md#venueNew) | **POST** /venue | new venue
[**venueNewAI**](VenueApi.md#venueNewAI) | **POST** /venue/{venueID}/{aiType} | newAI venue
[**venueShow**](VenueApi.md#venueShow) | **GET** /venue/{venueID} | show venue
[**venueUpvoteAI**](VenueApi.md#venueUpvoteAI) | **POST** /venue/{venueID}/{aiType}/upvote | upvoteAI venue


<a name="venueAll"></a>
# **venueAll**
> Venues venueAll()

all venue

Get all venues

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
try {
    Venues result = apiInstance.venueAll();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueAll");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Venues**](Venues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="venueDownvoteAI"></a>
# **venueDownvoteAI**
> venueDownvoteAI(aiType, venueID)

downvoteAI venue

Downvote an existing AI for a venue

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
String aiType = "aiType_example"; // String | Type of AI
String venueID = "venueID_example"; // String | Venue ID
try {
    apiInstance.venueDownvoteAI(aiType, venueID);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueDownvoteAI");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiType** | **String**| Type of AI | [enum: epiPen, bathroom, aed]
 **venueID** | **String**| Venue ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="venueNew"></a>
# **venueNew**
> venueNew(payload)

new venue

Create a new venue

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
NewVenue payload = new NewVenue(); // NewVenue | Type for creating a new venue
try {
    apiInstance.venueNew(payload);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueNew");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**NewVenue**](NewVenue.md)| Type for creating a new venue |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="venueNewAI"></a>
# **venueNewAI**
> venueNewAI(aiType, venueID)

newAI venue

Create a new AI for a venue

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
String aiType = "aiType_example"; // String | Type of AI to Create
String venueID = "venueID_example"; // String | Venue ID
try {
    apiInstance.venueNewAI(aiType, venueID);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueNewAI");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiType** | **String**| Type of AI to Create | [enum: epiPen, bathroom, aed]
 **venueID** | **String**| Venue ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="venueShow"></a>
# **venueShow**
> venueShow(venueID)

show venue

Get venue by ID

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
String venueID = "venueID_example"; // String | Venue ID
try {
    apiInstance.venueShow(venueID);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueShow");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **venueID** | **String**| Venue ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

<a name="venueUpvoteAI"></a>
# **venueUpvoteAI**
> venueUpvoteAI(aiType, venueID)

upvoteAI venue

Upvote an existing AI for a venue

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.VenueApi;


VenueApi apiInstance = new VenueApi();
String aiType = "aiType_example"; // String | Type of AI
String venueID = "venueID_example"; // String | Venue ID
try {
    apiInstance.venueUpvoteAI(aiType, venueID);
} catch (ApiException e) {
    System.err.println("Exception when calling VenueApi#venueUpvoteAI");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiType** | **String**| Type of AI | [enum: epiPen, bathroom, aed]
 **venueID** | **String**| Venue ID |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob, application/x-gob
 - **Accept**: application/json, application/xml, application/gob, application/x-gob

