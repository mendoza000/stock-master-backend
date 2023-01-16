# Stock master - BACKEND ⚡

StockMaster es un software para manejar y administrar inventarios de algun negocio, sus principales caracteriscias son:

- Permite tener varios usuarios con sus respectivos rangos como lo son ADMIN, CLIENT y SELLER
- Permite crear productos con la siguiente estructura

  ```go
  type Product struct {
  ID      int     `json:"id"`
  Title   string  `json:"title"`
  Brand   string  `json:"brand"`
  Stock   bool    `json:"stock"`
  Price   float32 `json:"price"`
  Details string  `json:"details"`
  Amount  int     `json:"amount"`
  }
  ```

---

## Rutas de API 🐙

### Rutas de Products

- [GET] `/api/products`

  Obetenemos todos los productos

  Parametros: `null`

- [GET] `/api/products`

  Obtenemos un producto por la ID

  Parametros: `ID: int`

- [POST] `/api/products`

  Creamos un producto

  Parametros:

  ```go
  type Product struct {
  Title   string  `json:"title"`
  Brand   string  `json:"brand"`
  Stock   bool    `json:"stock"`
  Price   float32 `json:"price"`
  Details string  `json:"details"`
  Amount  int     `json:"amount"`
  }
  ```

- [DELETE] `/api/products`

  Eliminamos un producto por la ID

  Parametros: `ID: int`

- [PUT] `/api/products`

  Actualizamos un producto por la ID

  Parametros: `ID: int`
