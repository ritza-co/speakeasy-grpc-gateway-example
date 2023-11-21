# SpeakeasyDrink

A drink served at the speakeasy


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `name`                                                                | *string*                                                              | :heavy_check_mark:                                                    | The name of the drink                                                 | Gin and Tonic                                                         |
| `price`                                                               | *number*                                                              | :heavy_minus_sign:                                                    | The price of the drink                                                | 5.99                                                                  |
| `productCode`                                                         | *string*                                                              | :heavy_minus_sign:                                                    | Unique drink identifier for server requests                           | 2438ac3c-37eb-4902-adef-ed16b4431030                                  |
| `stock`                                                               | *number*                                                              | :heavy_minus_sign:                                                    | The stock of the drink                                                | 10                                                                    |
| `type`                                                                | [shared.DrinkDrinkType](../../../sdk/models/shared/drinkdrinktype.md) | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |