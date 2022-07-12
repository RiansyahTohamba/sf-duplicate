

# user1 - chrome
MTY1NzYwNjYxMnxOd3dBTkZjeldWSldRMHhQVms4MFUwWkhURU5OUzBSRU1sQkZXRTFXVVVrMFNFazJSMDQwUWsxV1ZsaFpRa1ZHTWs1VlEwaFNRa0U9fCXN6FREPYu3WxbK9hMMLtvE-YfkWJ0kzgeguanDQtO-

# user2 - postman
MTY1NzYwNjU4N3xOd3dBTkVZM1RUVkZTMGhYVDBkRlN6VkJVazFFV2t4UVQwTkdSMHhFU2xFelVVWk1RalJaTWtsRVIxVlpSMUpPUWtGU1JGaFlXVkU9fIVDSqsyJn_zquQNSBdlBOyikqG1bguk6GHdkXkcaMPO


# at redis-server

1. 
127.0.0.1:6379> GET session_W3YRVCLOVO4SFGLCMKDD2PEXMVQI4HI6GN4BMVVXYBEF2NUCHRBA
"\x0e\xff\x81\x04\x01\x02\xff\x82\x00\x01\x10\x01\x10\x00\x00\x1e\xff\x82\x00\x01\x06string\x0c\t\x00\acounter\x03int\x04\x02\x00\x02"


2. 
127.0.0.1:6379> GET session_F7M5EKHWOGEK5ARMDZLPOCFGLDJQ3QFLB4Y2IDGUYGRNBARDXXYQ
"\x0e\xff\x81\x04\x01\x02\xff\x82\x00\x01\x10\x01\x10\x00\x00\x1e\xff\x82\x00\x01\x06string\x0c\t\x00\acounter\x03int\x04\x02\x00\x10"

kenapa antara data yang disimpan di redis server beda dengan yang ada di client? (postman maupun chrome)

yang sama hanya ke nya, value nya beda?