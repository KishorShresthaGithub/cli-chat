## SET URL
```sh
export HOST_URL=http://localhost:8080 
```

## Get all
```sh
curl $HOST_URL/items
```

## Get item

```sh
curl $HOST_URL/items/1
```

## Create Item


```sh
curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"test","price":20}' $HOST_URL/items
```

## Update Item


```sh
curl -X PUT -H "Content-Type: application/json" -d '{"name":"test2","price":30}' $HOST_URL/items/1
```

## Delete Item


```sh
curl -X DELETE $HOST_URL/items/1
```


