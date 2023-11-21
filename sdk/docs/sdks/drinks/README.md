# Drinks
(*drinks*)

## Overview

Drinks API

### Available Operations

* [getDrink](#getdrink) - Get a drink
* [listAllDrinks](#listalldrinks) - Get a list of drinks

## getDrink

Get a drink

### Example Usage

```typescript
import { SDK } from "openapi";

(async() => {
  const sdk = new SDK();

  const res = await sdk.drinks.getDrink({
    productCode: "602a7da9-b8bb-46e6-b288-457b561029b8",
  });

  if (res.statusCode == 200) {
    // handle response
  }
})();
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetDrinkRequest](../../sdk/models/operations/getdrinkrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `retries`                                                                    | [utils.RetryConfig](../../internal/utils/retryconfig.md)                     | :heavy_minus_sign:                                                           | Configuration to override the default retry behavior of the client.          |
| `config`                                                                     | [AxiosRequestConfig](https://axios-http.com/docs/req_config)                 | :heavy_minus_sign:                                                           | Available config options for making requests.                                |


### Response

**Promise<[operations.GetDrinkResponse](../../sdk/models/operations/getdrinkresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 400-600         | */*             |

## listAllDrinks

Returns a list of all drinks available at the speakeasy

### Example Usage

```typescript
import { SDK } from "openapi";

(async() => {
  const sdk = new SDK();

  const res = await sdk.drinks.listAllDrinks({
    empty: {},
  });

  if (res.statusCode == 200) {
    // handle response
  }
})();
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.ListAllDrinksRequest](../../sdk/models/operations/listalldrinksrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `config`                                                                               | [AxiosRequestConfig](https://axios-http.com/docs/req_config)                           | :heavy_minus_sign:                                                                     | Available config options for making requests.                                          |


### Response

**Promise<[operations.ListAllDrinksResponse](../../sdk/models/operations/listalldrinksresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 400-600         | */*             |
