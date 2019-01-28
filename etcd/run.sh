if [ -z ${CLIENT_URLS+x} ]; then
  CLIENT_URLS="http://0.0.0.0:4001,http://0.0.0.0:2379"
  echo "Using default CLIENT_URLS ($CLIENT_URLS)"
else
  echo "Detected new CLIENT_URLS value of $CLIENT_URLS"
fi

if [ -z ${ADV_URLS+x} ]; then
  ADV_URLS="http://0.0.0.0:4001,http://0.0.0.0:2379"
  echo "Using default PEER_URLS ($ADV_URLS)"
else
  echo "Detected new PEER_URLS value of $ADV_URLS"
fi
ETCD_CMD="/bin/etcd --data-dir=/data --listen-client-urls=${CLIENT_URLS} --advertise-client-urls=${ADV_URLS} $*"
echo -e "Running '$ETCD_CMD'\nBEGIN ETCD OUTPUT\n"
exec $ETCD_CMD
