# Env variable

create .env from .env.dist & set correct values

# Create new migration

connect to container via
```
docker exec -it migrate /bin/sh
```

run in current working dir (/usr/app)
```
shell/migrate-create.sh add_new_table
```
if you have troubles fix permissions then run outside of docker
```
make fix-permissions
```

# Migrate up script
```
shell/migrate-up.sh
```

# Migrate down script
```
shell/migrate-down.sh
```

For MacOS local usage use this env setting (this will forward to hostmachine):
```
MYSQL_HOST=host.docker.internal
```
