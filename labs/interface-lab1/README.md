## Pricing Strategy

A bar wants to have two pricing strategy. Normal hours (no discount) and happy hours (50% discount).

You need to build the feature that allows the bar to choose/switch pricing strategy whenever they want.

Calculations on the total bill amount for the customer will be based on the time he ordered.

### Example:
```
A Customer comes in normal hours and orders a drink for 100. He stayed longer and ordered another drink in happy hours. His total bill should be calculated as:
Total = 100 + (50% of 100) = 150
```

<details>
  <summary>Hint</summary> 

Solution can be implemented using one of the two ways
1. Using interface applying strategy pattern
2. Using function as strategy inside CustomerBill struct

</details><br/>

