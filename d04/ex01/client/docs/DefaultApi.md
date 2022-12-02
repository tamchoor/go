# \DefaultApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyCandy**](DefaultApi.md#BuyCandy) | **Post** /buy_candy | 


# **BuyCandy**
> InlineResponse201 BuyCandy(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiBuyCandyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBuyCandyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**optional.Interface of Order**](Order.md)| summary of the candy order | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

