# API Endpoints

## Authentication

### POST `/register`
Register a new user.
- **Request Body**: `UserSignUp` DTO
- **Public**: Yes
- **Response**:
  - `200 OK`: `{ "message": "register success", "token": "jwt_token" }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Registration failed

### POST `/login`
Authenticate user and return JWT token.
- **Request Body**: `UserLogin` DTO (`email`, `password`)
- **Public**: Yes
- **Response**:
  - `200 OK`: `{ "message": "login success", "token": "jwt_token" }`
  - `401 Unauthorized`: Invalid credentials
  - `400 Bad Request`: Invalid request body

---

## Catalog (Products & Categories)

### Public Endpoints

#### GET `/products`
Fetch all products.
- **Public**: Yes
- **Response**:
  - `200 OK`: `{ "message": "get products success", "data": [...] }`
  - `404 Not Found`: No products or error

#### GET `/products/:id`
Fetch a specific product by ID.
- **Public**: Yes
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "get product by id success", "data": {...} }`
  - `404 Not Found`: Product not found

#### GET `/categories`
Fetch all categories.
- **Public**: Yes
- **Response**:
  - `200 OK`: `{ "message": "get categories success", "data": [...] }`
  - `404 Not Found`: Categories not found

#### GET `/categories/:id`
Fetch a specific category by ID.
- **Public**: Yes
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "get category by id success", "data": {...} }`
  - `404 Not Found`: Category not found

### Seller Endpoints (`/seller`)

#### Authorization
All seller endpoints require the `AuthorizeSeller` middleware (valid JWT + seller role).

#### POST `/seller/categories`
Create a new category.
- **Request Body**: `CreateCategoryRequest` DTO
- **Response**:
  - `200 OK`: `{ "message": "category created successfully" }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Creation failed

#### PATCH `/seller/categories/:id`
Edit an existing category.
- **Path Param**: `id` (uint)
- **Request Body**: `CreateCategoryRequest` DTO
- **Response**:
  - `200 OK`: `{ "message": "edit category success", "data": {...} }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Edit failed

#### DELETE `/seller/categories/:id`
Delete a category.
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "delete category success" }`
  - `500 Internal Server Error`: Deletion failed

#### POST `/seller/products`
Create a new product.
- **Request Body**: `CreateProduct` DTO
- **Response**:
  - `200 OK`: `{ "message": "create product success" }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Creation failed

#### GET `/seller/products`
Fetch seller’s own products.
- **Response**:
  - `200 OK`: `{ "message": "get products success", "data": [...] }`
  - `404 Not Found`: Products not found

#### GET `/seller/products/:id`
Fetch seller’s own product by ID.
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "get product by id success", "data": {...} }`
  - `404 Not Found`: Product not found

#### PUT `/seller/products/:id`
Edit product details.
- **Path Param**: `id` (uint)
- **Request Body**: `CreateProduct` DTO
- **Response**:
  - `200 OK`: `{ "message": "edit product success", "data": {...} }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Edit failed

#### PATCH `/seller/products/:id`
Update product stock.
- **Path Param**: `id` (uint)
- **Request Body**: `UpdateStock` DTO
- **Response**:
  - `200 OK`: `{ "message": "update stock success", "data": {...} }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Stock update failed

#### DELETE `/seller/products/:id`
Delete a product.
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "delete product success" }`
  - `500 Internal Server Error`: Deletion failed

---

## User Profile & Cart

### Authorization
All endpoints under `/users` require `Authorize` middleware (valid JWT).

#### POST `/users/verify`
Submit verification code.
- **Request Body**: `VerificationCode` DTO
- **Response**:
  - `200 OK`: `{ "message": "verified successfully" }`
  - `400 Bad Request`: Invalid code format
  - `500 Internal Server Error`: Verification failed

#### GET `/users/verify`
Generate and return a new verification code.
- **Response**:
  - `200 OK`: `{ "message": "get verification code success", "code": 123456 }`
  - `500 Internal Server Error`: Code generation failed

#### POST `/users/profile`
Create user profile.
- **Request Body**: `ProfileInput` DTO
- **Response**:
  - `200 OK`: `{ "message": "create profile success" }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Profile creation failed

#### GET `/users/profile`
Fetch user profile.
- **Response**:
  - `200 OK`: `{ "message": "get profile success", "user": {...} }`
  - `500 Internal Server Error`: Fetch failed

#### PATCH `/users/profile`
Update user profile.
- **Request Body**: `ProfileInput` DTO
- **Response**:
  - `200 OK`: `{ "message": "update profile success", "user": {...} }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Update failed

#### POST `/users/cart`
Add item(s) to cart.
- **Request Body**: `CartRequest` DTO
- **Response**:
  - `200 OK`: `{ "message": "add to cart success", "data": [...] }`
  - `400 Bad Request`: Invalid request body
  - `500 Internal Server Error`: Cart update failed

#### GET `/users/cart`
Fetch user’s cart.
- **Response**:
  - `200 OK`: `{ "message": "get cart success", "data": [...] }`
  - `404 Not Found`: Cart not found

#### GET `/users/order`
Fetch user’s orders.
- **Response**:
  - `200 OK`: `{ "message": "get orders success", "orders": [...] }`
  - `404 Not Found`: Orders not found

#### GET `/users/order/:id`
Fetch specific order by ID.
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "get order by id success", "order": {...} }`
  - `404 Not Found`: Order not found

#### POST `/users/become-seller`
Request to become a seller.
- **Request Body**: `SellerInput` DTO
- **Response**:
  - `200 OK`: `{ "message": "become seller success", "token": "...." }`
  - `400 Bad Request`: Invalid input
  - `500 Internal Server Error`: Request failed

---

## Transactions & Payments

### Buyer Endpoints (`/buyer`)

#### Authorization
All buyer endpoints require `Authorize` middleware (valid JWT).

#### POST `/buyer/payment`
Initiate payment (creates or returns existing Stripe PaymentIntent).
- **Response**:
  - `200 OK`: `{ "message": "create payment", "stripe_pub_key": "...", "client_secret": "..." }`
  - `404 Not Found`: Cart empty
  - `500 Internal Server Error`: Payment initiation failed

#### GET `/buyer/verify`
Verify payment status with Stripe and finalize order if succeeded.
- **Response**:
  - `200 OK`: `{ "message": "payment verified successfully", "data": PaymentIntent }`
  - `404 Not Found`: No active payment
  - `500 Internal Server Error`: Verification or order creation failed

### Seller Endpoints (`/seller`)

#### Authorization
All seller transaction endpoints require `AuthorizeSeller` middleware.

#### GET `/seller/order`
Fetch seller’s orders (placeholder — returns stub).
- **Response**:
  - `200 OK`: `{ "message": "get orders success" }`

#### GET `/seller/order/:id`
Fetch seller’s specific order by ID (placeholder — returns stub).
- **Path Param**: `id` (uint)
- **Response**:
  - `200 OK`: `{ "message": "get order details success" }`