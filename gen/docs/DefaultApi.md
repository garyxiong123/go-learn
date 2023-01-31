# DefaultApi

All URIs are relative to *http://petstore.swagger.io:8080/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**testapiGetStringByIntSomeIdGet**](DefaultApi.md#testapiGetStringByIntSomeIdGet) | **GET** /testapi/get-string-by-int/{some_id} | Add a new pet to the store
[**testapiGetStructArrayByStringSomeIdGet**](DefaultApi.md#testapiGetStructArrayByStringSomeIdGet) | **GET** /testapi/get-struct-array-by-string/{some_id} | 


<a name="testapiGetStringByIntSomeIdGet"></a>
# **testapiGetStringByIntSomeIdGet**
> String testapiGetStringByIntSomeIdGet(someId)

Add a new pet to the store

get string by ID

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io:8080/v2");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    Integer someId = 56; // Integer | Some ID
    try {
      String result = apiInstance.testapiGetStringByIntSomeIdGet(someId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#testapiGetStringByIntSomeIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **someId** | **Integer**| Some ID |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | ok |  -  |
**400** | We need ID!! |  -  |
**404** | Can not find ID |  -  |

<a name="testapiGetStructArrayByStringSomeIdGet"></a>
# **testapiGetStructArrayByStringSomeIdGet**
> String testapiGetStructArrayByStringSomeIdGet(someId, offset, limit)



get struct array by ID

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io:8080/v2");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String someId = "someId_example"; // String | Some ID
    Integer offset = 56; // Integer | Offset
    Integer limit = 56; // Integer | Offset
    try {
      String result = apiInstance.testapiGetStructArrayByStringSomeIdGet(someId, offset, limit);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#testapiGetStructArrayByStringSomeIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **someId** | **String**| Some ID |
 **offset** | **Integer**| Offset |
 **limit** | **Integer**| Offset |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | ok |  -  |
**400** | We need ID!! |  -  |
**404** | Can not find ID |  -  |

