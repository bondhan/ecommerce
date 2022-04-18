# ecommerce


## create migration
```bash
migrate create -ext sql -dir migrations/scripts/ -seq -digits 2 create_table_order_details
```

## migrate up
```bash
migrate -database mysql://root@/ecommerce -path ./migrations/scripts up
```

## migrate down
```bash
migrate -database mysql://root@/ecommerce -path ./migrations/scripts down
```