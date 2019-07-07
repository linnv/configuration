./configure
make
su
make install
adduser postgres
mkdir /usr/local/pgsql/data
chown postgres /usr/local/pgsql/data
su - postgres
/usr/local/pgsql/bin/initdb -D /usr/local/pgsql/data
/usr/local/pgsql/bin/postgres -D /usr/local/pgsql/data >logfile 2>&1 &
/usr/local/pgsql/bin/createdb test
/usr/local/pgsql/bin/psql test


Please note the following commands:

\list or \l: list all databases
\dt: list all tables in the current database
You will never see tables in other databases, these tables aren't visible. You have to connect to the correct database to see its tables (and other objects).

To switch databases:

\connect database_name or \c database_name



[
CREATE TABLE weather (
    city            varchar(80),
    temp_lo         int,           -- low temperature
    temp_hi         int,           -- high temperature
    prcp            real,          -- precipitation
    date            date
);
]


 psql -U postgres -c 'SHOW config_file'
