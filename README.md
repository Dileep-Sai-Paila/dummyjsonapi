# Endpoints
The API exposes the following 4 endpoints:

### 1. Get All Products (Paginated)
- Method: GET
- Endpoint: /products/all
- Description: Fetches a paginated list of all products from the source API.

Example Response:

{\
    "products": [ { "id": 1, ... }, { "id": 2, ... } ],\
    "total": 194,\
    "skip": 0,\
    "limit": 30\
}

### 2. Get Products by Category
- Method: GET
- Endpoint: /products/category/{categoryName}
- Description: Fetches all products for a single, specified category.

Example Request: GET /products/category/laptops

### 3. Get All Products Grouped by Category
- Method: GET
- Endpoint: /products/categories
- Description: Fetches all available categories and then returns all products, grouped by their category name.

Example Response:

{\
    "laptops": [ { "id": 6, ... } ],\
    "smartphones": [ { "id": 1, ... } ],\
    "fragrances": [ ... ]\
}

#### 4. Get Products from a List of Categories
- Method: POST
- Endpoint: /products/by-categories
- Description: Accepts a list of category names and returns a single, merged list of all products from those categories.

Example Request Body:

{\
  "categories": ["laptops", "smartphones"]\
}