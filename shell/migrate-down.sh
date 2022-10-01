if [ -z $MYSQL_HOST ]
then
    # localhost machine IP
    MYSQL_HOST=$(echo -e `/sbin/ip route | awk '/default/ { print $3 }'`)
fi

./migrate \
	-source file://./migrations \
	-database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DATABASE" down 1
