Hi kawan2 Buyer Growth Platform Tech Tribe!

Tugas kalian adalah membuat dua buah API set dan get ke dalam redis dan db docker

API "set_data" method *POST*
```
Request:
{
    "name":"sebuah_nama",
    "status":1
}

Response:
{
    "id":1
}
```

API "get_data" method *GET*
```
Request:
curl localhost:8181/get_data?id=1

Response:
{
    "id":1,
    "name":"sebuah_nama",
    "status":1
}
```

NOTE:
1. Cek initRedis dan initDB, pastikan benar untuk urlnya
2. Data yg disimpan ke table `table_1` (lihat di schema.sql)