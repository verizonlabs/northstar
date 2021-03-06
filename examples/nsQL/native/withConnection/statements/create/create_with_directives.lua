local nsQL = require("nsQL")

function main()
	local query = [[
		CREATE TABLE IF NOT EXISTS nssim.test5 (col1 text, col2 int, col3 float, col4 double, col5 text, PRIMARY KEY((col1, col2), col3, col4))
		WITH CLUSTERING ORDER BY (col3 asc, col4 desc) AND COMPACT STORAGE;
	]]
	local source = {
		Protocol = "cassandra",
		Host = "10.32.49.6",
		Port = "31838",
		Backend = "native"
	}
	local options = {}
	local connection = createConnection(source)
	processQuery(connection, query, options)
	teardownConnection(connection)
end

function createConnection(source)
	local connection, err = nsQL.connect(source)
	if(err ~= nil) then
		error(err)
	end
	return connection
end

function teardownConnection(connection)
	local err = connection:disconnect()
	if(err ~= nil) then
		error(err)
	end
end

function processQuery(connection, query, options)
	local resp, err = connection:query(query, options)
	if(err ~= nil) then
		error(err)
	end
	return resp
end