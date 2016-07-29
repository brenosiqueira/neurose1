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