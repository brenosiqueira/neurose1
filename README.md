# neurose1

## Como desenvolver?

1. Clone o repositório.
2. Rodar app pelo docker-compose

```console
git clone git@github.com:13team/neurose1.git
cd neurose1
docker-compose up -d
```


## Teste das chamadas REST

```
#cadastrar order

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "number=10&reference=bluedream&status=DRAFT&notes=dinheiro" http://localhost:9090/neurorder/order

# Cadastra Items http://localhost:9090/neurorder/order/{id}/item

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "sku=A7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1&unit_price=1000&quantity=2" http://localhost:9090/neurorder/order/B7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1/item

# Cadastra Transacao

curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "external_id=UUEF748A-EE40-48C1-8AF7-0E8A99D4D7A1&amount=1000&type=PAYMENT&authorization_code=100010&card_brand=VISA&card_bin=123465&card_last=1234" http://localhost:9090/neurorder/order/B7EF748A-EE40-48C1-8AF7-0E8A99D4D7A1/payment
```


## Esboço Scripts de banco (Obs.: a aplicação ainda não está acessando a base)

```
CREATE KEYSPACE neurose WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

USE neurose;
create table neurorder (
  id uuid primary key,
  number varchar,
  reference varchar,
  status varchar,
  created_at timestamp,
  updated_at timestamp,
  notes varchar,
  price int
);

create table neurorder_item (
  id uuid primary key,
  sku uuid,
  unit_price int,
  quantity int,
  neurorder_id uuid
);

create table transaction(
   id uuid primary key,
   external_id varchar,
   amount int,
   type varchar,
   authorization_code varchar,
   card_brand varchar
   card_bin varchar
   card_last varchar
   neurorder_id uuid
  )

```
