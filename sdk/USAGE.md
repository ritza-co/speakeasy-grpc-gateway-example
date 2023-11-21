<!-- Start SDK Example Usage -->
```typescript
import { SDK } from "openapi";

(async () => {
    const sdk = new SDK();

    const res = await sdk.drinks.getDrink({
        productCode: "602a7da9-b8bb-46e6-b288-457b561029b8",
    });

    if (res.statusCode == 200) {
        // handle response
    }
})();

```
<!-- End SDK Example Usage -->