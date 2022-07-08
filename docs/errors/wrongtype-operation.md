# WRONGTYPE Operation 
WRONGTYPE Operation against a key holding the wrong kind of value

# how to reproduce bug?

misalkan kita sudah buat key article:11 untuk SET 
SADD(ctx, 'article:11', 'member')

lalu kita coba lakukan operasi HASH pada data type SET tadi.
HGETALL(ctx, 'article:11')

maka akan muncul error
WRONGTYPE Operation against a key holding the wrong kind of value

# kenapa bisa terjadi?
ini maksudnya: article:11 sudah diset dalam SET
maka jika kita sudah menetapkan dia sebagai SET, operasi selain utk SET tidak boleh dilakukan.
misalkan kita coba lakukan HGETALL untuk article:11, yang notabene operasi untuk data type HASH.


# penjelasan dari SF
This error means that the value indexed by the key "l_messages" is not of type hash, but rather something else. You've probably set it to that other value earlier in your code. 

Try various other value-getter commands, starting with GET, to see which one works and you'll know what type is actually here.

# penjelasan nomor 1

dijelaskan dulu data type yang ada pada redis.
kemudian cara untuk retrieve data untuk masing-masing data type tersebut.
kemudian cara untuk mengecek data type dari key yang coba kita retrieve atau create.

Redis supports 6 data types. You need to know what type of value that a key maps to, as for each data type, the command to retrieve it is different.

Here are the commands to retrieve key value:

if value is of type string -> GET <key>
if value is of type hash -> HGETALL <key>
if value is of type lists -> lrange <key> <start> <end>
if value is of type sets -> smembers <key>
if value is of type sorted sets -> ZRANGEBYSCORE <key> <min> <max>
if value is of type stream -> xread count <count> streams <key> <ID>. https://redis.io/commands/xread
Use the TYPE command to check the type of value a key is mapping to:

type <key>

# referensi
https://stackoverflow.com/questions/37953019/wrongtype-operation-against-a-key-holding-the-wrong-kind-of-value-php