# API e-Commerce

Salah satu tugas kelompok untuk menyelesaikan course di Alterra Academy https://alterra.id

## Endpoint

- `[POST] /register`

  ```
  options role (0 == customers, 1 == merchants)
  example request :

  {
      "name" : "testing",
      "username" : "testing",
      "hp" : "081234567890",
      "email" : "testing@gmail.com",
      "password" : "password",
      "role" : 0
  }
  ```

- `[POST] /login `

  ```
  options (username, hp, email)
  example request :

  {
     "username" : "testing",
     "password" : "password"
  }
  ```

- `[PUT] /users/{username}`

  ```
  use autorization (bearer)
  example request :

  {
   "name" : "testing update"
  }
  ```

- `[DELETE] /users/{username}`
- `[GET] /users/{username}`
- `[POST] /admin/categories`

  ```
  use autorization (bearer)
  example request :

  {
     "name" : "produk fisik"
  }
  ```

- `[GET] /categories`
- `[POST] /merchants/products`

  ```
  use autorization (bearer)
  example request :

  {
     "name" : "Laptop 24 inch",
     "price" : 1000000,
     "stock" : 5,
     "description" : "ini laptop sangat multi talenta",
     "image" : "gambar/123.jpg",
     "category_id" : 1
  }
  ```

- `[PUT] /merchants/products/{slug}`

  ```
  use autorization (bearer)
  example request :

  {
     "price" : 1000000,
     "stock" : 5
  }
  ```

- `[DELETE] /merchants/products/{slug}`
- `[GET] /merchants/products`

  ```
  use autorization (bearer)
  ```

- `[GET] /products`
- `[GET] /products/{slug}`
- `[GET] /products/category/{id}`
- `[GET] /search`

  ```
  use query parameters
  example request :

  parameter=name    || value=laptop
  ```

## Contributing

- Mahmuda Karima (DAKA) - https://github.com/BE8-Daka
- Dwi Fajar Bachtiar - https://github.com/DwiBactiar12

## Copyrights

- Mei 2022
