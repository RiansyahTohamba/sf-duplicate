# apa saja dependensi yang dibutuhkan?
Session Based authentication dengan Gin dan Redis

Session Based authentication menggunakan sesi, yang merupakan file kecil yang menyimpan informasi tentang pengguna termasuk ID sesi unik, waktu login dan login expired time, dan banyak lagi, yang dibuat oleh (go) server dan disimpan dalam database setelah kita berhasil login.

Pada kali ini kita akan menggunakan bantuan pustaka Gin yaitu: gin-contrib/sessions.

# bagaimana gin-contrib/sessions menggunakan redis?
Session pada gin-contrib/sessions dapat disimpan dalam bentuk memcached,mongo, postgres maupun redis. 

pada gin.Engine yang sudah dibuat, kita dapat menggunakan package session

```go
router := gin.Default()
router.Use(sessions.Sessions("counter", getRedisStore()))
```
Sessions menawarkan pilihan pada tipe store yang akan kita gunakan. Pada tulisan ini, tipe store yang akan kita gunakan adalah redis store.

```go
redisStore, _ := redis.NewStore(size, "tcp", "localhost:6379", "", pwd)
```

syntax untuk menyimpan session pada Go adalah sebagai berikut 

```go
session.Set("counter", counter)
session.Save()
```

# bagaimana tampilan session pada cookie header?
Pada cookie header kita dapat melihat bagaimana session disimpan.
Name: Counter
Value: MTY1NzU5NTg4NnxOd3dBTkVZM1RUVkZTMGhYVDBkRlN6VkJVazFFV2t4UVQwTkdSMHhFU2xFelVVWk1RalJaTWtsRVIxVlpSMUpPUWtGU1JGaFlXVkU9fNsNnynHNhIuaO_b83Kzk_ehObcUG_Dym6j_O11V-7IH


## pada redis-server, session disimpan seperti apa?
Sementara itu, Redis-server juga menyimpan session yang sudah disimpan dengan package session. Session disimpan dengan redis data type String. Nama key nya diawali dengan `session_`. Berikut contoh dari session counter yang sudah disimpan pada redis server.

127.0.0.1:6379> GET session_F7M5EKHWOGEK5ARMDZLPOCFGLDJQ3QFLB4Y2IDGUYGRNBARDXXYQ

"\x0e\xff\x81\x04\x01\x02\xff\x82\x00\x01\x10\x01\x10\x00\x00\x1e\xff\x82\x00\x01\x06string\x0c\t\x00\acounter\x03int\x04\x02\x00\n"

kita bisa lihat nama dari sesi `counter` yang sudah kita set pada kode Golang sebelumnya.

`x00\acounter\x03int`
