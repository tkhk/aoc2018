ngx.req.read_body()
local d = ngx.req.get_body_data()

local result = 0
for v in string.gmatch(d, "(.-)\n") do
	result = result + tonumber(v)
end

ngx.say(result)
