ngx.req.read_body()
local d = ngx.req.get_body_data()

local tmpResult = 0
local finalResult = nil
local resultTable = {}

while true do
	for v in string.gmatch(d, "(.-)\n") do
		tmpResult = tmpResult + tonumber(v)
		if resultTable[tmpResult] ~= nil then
			finalResult = tmpResult
			break
		else 
			resultTable[tmpResult] = true
		end
	end

	if finalResult ~= nil then
		break
	end
end

ngx.say(finalResult)
