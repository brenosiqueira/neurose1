package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/gocql/gocql"
)

// Estrutura das ROTAS
type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRESTMap(session *gocql.Session) {
  router := newRouter(session)
  log.Fatal(http.ListenAndServe(":9090", router))
}

func newRouter(session *gocql.Session) *mux.Router {
  // Definicao das ROTAS
  var routes = Routes {
    Route { "Order", "POST", "/api/v1/order", func(w http.ResponseWriter, r *http.Request) { Order(w, r, session) } ,  },
    Route { "OrderIten", "POST", "/api/v1/order/{id}/item",  func(w http.ResponseWriter, r *http.Request) { OrderIten(w, r, session) }, },
    Route { "Payment", "POST", "/api/v1/order/{id}/payment", func(w http.ResponseWriter, r *http.Request) { Payment(w, r, session) }, },
  }

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

type OrderResponse struct {
  Uuid string `json:"uuid"`
}

// Handlers para as Rotas, ou seja, quem trata as requisicoes HTTP

func Order(w http.ResponseWriter, r *http.Request, session *gocql.Session) {
  //Número da Order. Geralmente esse número representa o ID da Order em um sistema externo através da integração com parceiros.
  number := r.FormValue("number")
  //Referência da Order. Usada para facilitar o acesso ou localização da mesma.
  reference := r.FormValue("reference")
  //Status da Order. DRAFT | ENTERED | CANCELED | PAID | APPROVED | REJECTED | RE-ENTERED | CLOSED
  status := r.FormValue("status")
  // Um texto livre usado pelo Merchant para comunicação.
  notes := r.FormValue("notes")
  fmt.Printf("Chegou uma requisicoes de order: number %s, reference %s, status %s, notes %s \n", number, reference, status, notes)

  uuid := gocql.TimeUUID()
  statusInt := translateStatus(status)
  if statusInt == 99 {
    http.Error(w, "Parametro status invalido", http.StatusPreconditionFailed)
    return
  }

  // Gravar no banco e retornar o UUID gerado
  if err := session.Query("INSERT INTO neurorder (order_id, number, reference, status, notes) VALUES (?,?,?,?,?)", uuid, number, reference, statusInt, notes).Exec(); err != nil {
    fmt.Println(err)
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  } else {
    // Retornar um JSON com o UUID (id da Order)
    w.WriteHeader(http.StatusCreated)
    orderResponse := OrderResponse { Uuid: uuid.String() }
    json.NewEncoder(w).Encode(orderResponse)
  }
}

 func translateStatus(status string) int {
   var statusInt int
   switch status {
     case "DRAFT":
      statusInt = 0
     case "ENTERED":
      statusInt = 1
     case "CANCELED":
      statusInt = 2
     case "PAID":
      statusInt = 3
     case "APPROVED":
      statusInt = 4
     case "REJECTED":
      statusInt = 5
     case "RE-ENTERED":
      statusInt = 6
     case "CLOSED":
      statusInt = 7
     default:
       // DESCONHECIDO
      statusInt = 99
    }

   return statusInt
 }

func OrderIten(w http.ResponseWriter, r *http.Request, session *gocql.Session) {
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

func Payment(w http.ResponseWriter, r *http.Request, session *gocql.Session) {
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
