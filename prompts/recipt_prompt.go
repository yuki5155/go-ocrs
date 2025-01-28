package prompts

const StoreSchemaPrompt = `
Please generate JSON data in the following format based on the provided information. Include the store name and date if available. For each item, include the name, price (including tax), and price (excluding tax).

Format:
{
  "store_name": "Store Name (if available, otherwise omit this field)",
  "date": "Date in yyyy-mm-dd format (if available, otherwise omit this field)",
  "items": [
    {
      "name": "Item Name",
      "price": Price,
      "price_excl_tax": Price (excluding tax)
    },
    {
      "name": "Item Name",
      "price": Price,
    }
  ]
}

Rules:
1. The store name should only be included if provided. If not, exclude it from the JSON.
2. The date should only be included if provided. If not, exclude it from the JSON.
3. Each item must have a name, price.


Example:
Input:
Store Name: Tech Hub
Date: 2025-01-25
Items:
- Name: Monitor, Price: 30000
- Name: Keyboard, Price: 4000

Output:
{
  "store_name": "Tech Hub",
  "date": "2025-01-25",
  "items": [
    {
      "name": "Monitor",
      "price": 33000,
    },
    {
      "name": "Keyboard",
      "price": 4400,
    }
  ]
}

Input:
Store Name: (unknown)
Date: (unknown)
Items:
- Name: Laptop, Price: 100000
- Name: Headphones, Price: 5000

Output:
{
  "items": [
    {
      "name": "Laptop",
      "price": 110000,
      "price": 100000
    },
    {
      "name": "Headphones",
      "price": 5500,
      "price": 5000
    }
  ]
}
`
