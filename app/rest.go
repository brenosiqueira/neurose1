package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  // "runtime"
)

// Estrutura das ROTAS
type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

// Definicao das ROTAS
var routes = Routes {
  Route { "Order", "POST", "/api/v1/order", Order,  },
  Route { "OrderIten", "POST", "/api/v1/order/{id}/item",  OrderIten, },
  Route { "Payment", "POST", "/api/v1/order/{id}/payment", Payment, },
}

func InitRESTMap() {
  router := newRouter()
  log.Fatal(http.ListenAndServe(":9090", router))
}

func newRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
  return router
}

// Handlers para as Rotas, ou seja, quem trata as requisicoes HTTP

func Order(w http.ResponseWriter, r *http.Request) {
  //Número da Order. Geralmente esse número representa o ID da Order em um sistema externo através da integração com parceiros.
  number := r.FormValue("number")
  //Referência da Order. Usada para facilitar o acesso ou localização da mesma.
  reference := r.FormValue("reference")
  //Status da Order. DRAFT | ENTERED | CANCELED | PAID | APPROVED | REJECTED | RE-ENTERED | CLOSED
  status := r.FormValue("status")
  // Um texto livre usado pelo Merchant para comunicação.
  notes := r.FormValue("notes")

  fmt.Println("================ Order ===================")
  fmt.Println(number)
  fmt.Println(reference)
  fmt.Println(status)
  fmt.Println(notes)

  // Gravar no banco e retornar o UUID gerado

  // Retornar um JSON com o UUID (id da Order)
}

func OrderIten(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"] // Id da Order

  // UUID que identifica unicamente o  Produto que está sendo comprado.
  sku := r.FormValue("sku")
  // Preço do produto.	integer Valor em centavos, e.g R$ 10,00 serão representados como 1000, em centavos.
  unit_price := r.FormValue("unit_price")
  // Quantidade de produtos comprados. integer
  quantity := r.FormValue("quantity")

  fmt.Println("================ OrderIten ===================")
  fmt.Println(id)
  fmt.Println(sku)
  fmt.Println(unit_price)
  fmt.Println(quantity)
  // Gravar no banco o OrderIten vinculado ao  (id da Order)

  // Retornar um JSON com o UUID (id da Order)

  fmt.Fprintf(w, "id recebido:", id)
}

func Payment(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  // Id externa da transação
  external_id := r.FormValue("external_id")
  // Valor da transação	integer
  amount := r.FormValue("amount")
  // Tipo da transação	enum  PAYMENT   CANCEL
  transaction_type := r.FormValue("type")
  // Código de Autorização da Transação
  authorization_code := r.FormValue("authorization_code")
  // Bandeira de pagamento do cartão
  card_brand := r.FormValue("card_brand")
  // Bin do cartão	string - 6 primeiros digitos
  card_bin := r.FormValue("card_bin")
  // Last do cartão	string - 4 ultimos digitos
  card_last := r.FormValue("card_last")

  fmt.Println("================ Payment ===================")
  fmt.Println(id)
  fmt.Println(external_id)
  fmt.Println(amount)
  fmt.Println(transaction_type)

  fmt.Println(authorization_code)
  fmt.Println(card_brand)
  fmt.Println(card_bin)
  fmt.Println(card_last)

  fmt.Fprintf(w, "id recebido:", id)
}
