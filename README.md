# neurose1

## Como desenvolver?

1. Clone o repositório.
2. Rodar app pelo docker-compose

```console
git clone git@github.com:13team/neurose1.git
cd neurose1
docker-compose up -d
```


## Como criar o banco?

1. Rodar comando cqlsh
2. Criando banco no docker

```console
# para garantir que o schema foi copiado para o container
docker cp config/schema.cql $NOME_CONTAINER_CASSANDRA:/config/schema.cql

docker exec -it $NOME_CONTAINER_CASSANDRA cqlsh -f config/schema.cql
```

## Teste das chamadas REST

```
#Cadastrar order

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "number=10&reference=bluedream&status=DRAFT&notes=dinheiro" http://localhost:9090/api/v1/order

#Cadastra Items http://localhost:9090/api/v1/order/{id}/item

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "sku=A7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1&unit_price=1000&quantity=2" http://localhost:9090/api/v1/order/B7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1/item

#Cadastra Transacao

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "external_id=UUEF748A-EE40-48C1-8AF7-0E8A99D4D7A1&amount=1000&type=PAYMENT&authorization_code=100010&card_brand=VISA&card_bin=123465&card_last=1234" http://localhost:9090/api/v1/order/B7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1/payment
```

## schema
```cql
CREATE KEYSPACE IF NOT EXISTS "neurose1"
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 2 };

USE neurose1;

CREATE TYPE IF NOT EXISTS orderitem (
  sku text,
  unit_price int,
  quantity int
);

CREATE TYPE IF NOT EXISTS creditcard (
    brand text,
    bin int,
    last int
);

CREATE TYPE IF NOT EXISTS payment (
  external_id text,
  amount int,
  transaction_type text,
  auth_code text,
  creditcard frozen <creditcard>
);

CREATE TABLE IF NOT EXISTS neurorder (
  order_id uuid,
  number text,
  reference text,
  status int,
  notes text,

  items list<frozen <orderitem>>,
  payments list<frozen <payment>>,

  PRIMARY KEY (order_id)
);
```

## Notas
- Se usa o Atom para escrever os scripts CQL, você pode instalar o plugin language-cassandra-cql para deixar a sintaxe em highlight.
```console
apm install language-cassandra-cql
```
