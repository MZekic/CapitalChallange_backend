CREATE TABLE IF NOT EXISTS users(
   id text PRIMARY KEY,
   username text UNIQUE NOT NULL,
   password text NOT NULL,
   email text UNIQUE NOT NULL,
   current_game_number integer,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_balance(
   id text PRIMARY KEY,
   user_id text NOT NULL,
   starting_balance numeric(20,10),
   current_balance numeric(20,10),
   game_number integer,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   CONSTRAINT fk_user_balance
    FOREIGN KEY (id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS companies(
   id text PRIMARY KEY,
   name text UNIQUE NOT NULL,
   ticker text UNIQUE NOT NULL,
   market text NOT NULL,
   primary_exchange text NOT NULL,
   type text NOT NULL,
   currency_name text NOT NULL,
   locale text,
   homepage_url text,
   description text,
   total_employees integer,
   logo_url text,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS company_stock(
   id text PRIMARY KEY,
   ticker text NOT NULL,
   close_price numeric(20,10),
   highest_price numeric(20,10),
   lowest_price numeric(20,10),
   number_of_transactions integer,
   open_price numeric(20,10),
   otc boolean,
   trading_volume integer,
   volume_weighted_average_price numeric(20,10),
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_assets(
   id text PRIMARY KEY,
   user_id text NOT NULL,
   company_stock_id text NOT NULL,
   quantity integer,
   game_number integer,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   CONSTRAINT fk_user_assets_user
    FOREIGN KEY (id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user_assets_company_stock
    FOREIGN KEY (id)
        REFERENCES company_stock(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_transactions(
   id text PRIMARY KEY,
   user_id text NOT NULL,
   company_stock_id text NOT NULL,
   buy_or_sell boolean,
   quantity integer,
   game_number integer,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   CONSTRAINT fk_user_transactions_user
    FOREIGN KEY (id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user_assets_company_stock
    FOREIGN KEY (id)
        REFERENCES company_stock(id)
        ON DELETE CASCADE
)