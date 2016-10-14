host="192.168.100.27"
cli -h $host Aerospike -b name -v "Aerospike, Inc."

cli -h $host -n test -o set -k Aerospike -b address -v "Mountain View, CA 94043"

cli -h $host -n test -o set -k Aerospike -b email -v "info@aerospike.com"
