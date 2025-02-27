# Up-Down migrasi
muhammadrizkisetyanto@MacBook-Air-Muhammad arabiya-syari-fiber-1 % migrate -database "postgresql://postgres:qXdMRsMSGEgQvVrLuBjmUAGkytJwsaWk@trolley.proxy.rlwy.net:59123/railway" -path internals/database/migrations/user up

# Dirty migrasi
muhammadrizkisetyanto@MacBook-Air-Muhammad arabiya-syari-fiber-1 % migrate -database "postgresql://postgres:qXdMRsMSGEgQvVrLuBjmUAGkytJwsaWk@trolley.proxy.rlwy.net:59123/railway" -path internals/database/migrations force 20250221005048

# Masuk database
muhammadrizkisetyanto@MacBook-Air-Muhammad arabiya-syari-fiber-1 % PGPASSWORD="qXdMRsMSGEgQvVrLuBjmUAGkytJwsaWk" psql -h trolley.proxy.rlwy.net -p 59123 -U postgres -d railway

# Buat migrasi
muhammadrizkisetyanto@MacBook-Air-Muhammad arabiya-syari-fiber-1 % migrate create -ext sql -dir internals/database/migrations create-table-user-profile
/Users/muhammadrizkisetyanto/Documents/arabiya-syari-fiber-1/internals/database/migrations/20250226092703_create-table-user-profile.up.sql
/Users/muhammadrizkisetyanto/Documents/arabiya-syari-fiber-1/internals/database/migrations/20250226092703_create-table-user-profile.down.sql


# Refresh port
kill -9 $(lsof -t -i:8080)


