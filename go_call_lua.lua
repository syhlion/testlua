cjson = require("json")

function process(a)
    local obj = cjson.decode(a)
    obj.id = 88888
    local str = cjson.encode(obj)
    return str
end
