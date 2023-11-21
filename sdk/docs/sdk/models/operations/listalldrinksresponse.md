# ListAllDrinksResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `contentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | HTTP response content type for this operation                           |
| `statusCode`                                                            | *number*                                                                | :heavy_check_mark:                                                      | HTTP response status code for this operation                            |
| `rawResponse`                                                           | [AxiosResponse](https://axios-http.com/docs/res_schema)                 | :heavy_check_mark:                                                      | Raw HTTP response; suitable for custom response parsing                 |
| `classes`                                                               | [shared.SpeakeasyDrink](../../../sdk/models/shared/speakeasydrink.md)[] | :heavy_minus_sign:                                                      | N/A                                                                     |
| `rpcStatus`                                                             | [shared.RpcStatus](../../../sdk/models/shared/rpcstatus.md)             | :heavy_minus_sign:                                                      | An unexpected error response.                                           |