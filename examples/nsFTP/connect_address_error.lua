local ftp = require("nsFTP")

function main()
    local _, err = ftp.connect("10.37.2.6:22")
    if err ~= nil then
        error(err)
    end
end