# API e-Commerce

Salah satu tugas kelompok untuk menyelesaikan course di Alterra Academy https://alterra.id

## Endpoint

- > `[POST] /register` endpoint ini digunakan untuk mendaftar akun.

  untuk pengisian field `role` (0 = customers (pembeli), 1 = merchants (penjual))

  ```
  {
      "name" : "testing",
      "username" : "testing",
      "hp" : "081234567890",
      "email" : "testing@gmail.com",
      "password" : "password",
      "role" : 0
  }
  ```

- > `[POST] /login ` endpoint ini digunakan untuk login

  untuk field `username` bisa diganti menjadi `email` atau `hp`

  ```
  {
     "username" : "testing",
     "password" : "password"
  }
  ```

- > `[PUT] /users/{username}` endpoint ini digunakan untuk mengedit user yang telah login

  ```
  {
   "name" : "testing update"
  }
  ```

- > `[DELETE] /users/{username}` endpoint ini digunakan untuk menghapus user
- > `[GET] /users/{username}` endpoint ini digunakan untuk mengambil data profil user
- > `[POST] /admin/categories` endpoint ini digunakan oleh admin untuk membuat kategori baru

  ```
  {
     "name" : "produk fisik"
  }
  ```

- > `[GET] /categories` endpoint ini digunakan untuk mengambil data kategori
- > `[POST] /merchants/products` endpoint ini digunakan oleh merchant untuk membuat produk baru

  ```
  {
     "name" : "Laptop 24 inch",
     "price" : 1000000,
     "stock" : 5,
     "description" : "ini laptop sangat multi talenta",
     "image" : "gambar/123.jpg",
     "category_id" : 1
  }
  ```

- > `[PUT] /merchants/products/{slug}` endpoint ini digunakan oleh merchant untuk mengedit data productnya

  ```
  {
     "price" : 1000000,
     "stock" : 5
  }
  ```

- > `[DELETE] /merchants/products/{slug}` endpoint ini digunakan oleh merchant untuk menghapus data produk
- > `[GET] /merchants/products`

  ```
  use autorization (bearer)
  ```

- > `[GET] /products`
- > `[GET] /products/{slug}`
- > `[GET] /products/category/{id}`
- > `[GET] /search`

  ```
  parameter=name    || value=laptop
  ```

## Contributing

- Mahmuda Karima (DAKA) - https://github.com/BE8-Daka
- Dwi Fajar Bachtiar - https://github.com/DwiBactiar12

## Copyrights

- Mei 2022
