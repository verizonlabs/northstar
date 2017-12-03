local nsQL = require("nsQL")
local nsOutput = require("nsOutput")

function main()
    local query = [[
        SELECT  JSON_FETCH(stats, 'module') as module,
                JSON_FETCH(stats, 'etime') as etime,
                JSON_FETCH(stats, 'interval') as interval
        FROM    logging.stats
        WHERE   msgdate = '2017-06-20' AND `group` = 'northstar' AND appname = 'data' AND interval = 3
                AND msgtime < '11:59:58' AND msgtime > '11:59:50'
        LIMIT   100;
    ]]

    local source = {
        Protocol = "cassandra",
        Host = "10.44.15.13",
        --Host = "cassandra1-log-dakota.mon-marathon-service.mesos",
        Port = "9042",
        Backend = "spark"
    }

    local options = {}
    local result = processQuery(query, source, options)
    return generateTable(result)
end

function processQuery(query, source, options)
    local resp, err = nsQL.query(query, source, options)
    if(err ~= nil) then
        error(err)
    end
    return resp
end

function generateTable(table)
    local out, err = nsOutput.table(table)
    if(err ~= nil) then
        error(err)
    end
    return out
end


