CREATE TABLE "cost_calculation" (
  "id" SERIAL PRIMARY KEY,
  "parkcost" int,
  "assetcost" int,
  "loancost" int,
  "tax_cost" int,
  "asset_value_cost" int,
  "revenue" int,
  "total_cost" int,
  "result" int,
  "result_perc" int,
  "saving_revenue_amount" int,
  "saving_revenue_perc" int,
  "yield_diff" int,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);