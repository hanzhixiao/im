---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by saeipi.
--- DateTime: 2023/1/4 9:50 AM
---

local KEYS = {"LK:USER_INFO:1",
              "LK:USER_INFO:2",
              "LK:USER_INFO:3",
              "LK:USER_INFO:4"}
local ARGV = {1000,1,2,3,4 }

for i=1,  #KEYS do
    if i==1 then
        print('SET', KEYS[i], ARGV[i+1])
        print('Expire', KEYS[i], ARGV[1])
    else
        print(KEYS[i], ARGV[i+1])
    end
end

--redis-cli -h 127.0.0.1 -p 63791 -a lark2022 script load "
--for i=1,  #KEYS do
--    if i==1 then
--        redis.call('SET', KEYS[i], ARGV[i+1], \"EX\", ARGV[1])
--    else
--        redis.call('SET', KEYS[i], ARGV[i+1])
--    end
--end
--return 0
--"

