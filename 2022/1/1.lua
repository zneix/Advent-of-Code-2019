local function readAll(file)
    local f = assert(io.open(file, "rb"))
    local content = f:read("*a")
    f:close()
    return content
end

local function split(inputstr, sep)
	sep = sep or "%s"
    local t, sep = {}, sep or "%s"
    for str in string.gmatch(inputstr, "([^"..sep.."]+)") do
            table.insert(t, str)
    end
    return t
end

function add(array)
    local ret = 0
    for key, value in pairs(array) do
        ret = ret + value
    end
    return ret
end


local input = readAll("input")

local xd = input:gsub("\n\n", ":"):gsub("\n", " ")
local xd3 = split(xd, ":")

local integertable = {}

for key, value in pairs(xd3) do
    xddd = split(value, " ")
    table.insert(integertable, add(xddd))
end

print(math.max(table.unpack(integertable)))
