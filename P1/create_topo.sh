ip netns add node1
ip netns add node2
ip netns add node3
ip netns add node4
ip netns add router

ip link add br1 type bridge
ip link set dev br1 up
ip link add br2 type bridge
ip link set dev br2 up


ip link add veth-1 type veth peer name veth-br11 ## node 1
ip link set veth-1 netns node1
ip link set veth-br11 master br1
ip netns exec node1 ip addr add 172.0.0.2/24 dev veth-1
ip netns exec node1 ip link set veth-1 up
ip link set veth-br11 up

ip link add veth-2 type veth peer name veth-br12 ## node 2
ip link set veth-2 netns node2
ip link set veth-br12 master br1
ip netns exec node2 ip addr add 172.0.0.3/24 dev veth-2
ip netns exec node2 ip link set veth-2 up
ip link set veth-br12 up


ip link add veth-3 type veth peer name veth-br23 ## node 3
ip link set veth-3 netns node3
ip link set veth-br23 master br2
ip netns exec node3 ip addr add 10.10.0.2/24 dev veth-3
ip netns exec node3 ip link set veth-3 up
ip link set veth-br23 up

ip link add veth-4 type veth peer name veth-br24 ## node 4
ip link set veth-4 netns node4
ip link set veth-br24 master br2
ip netns exec node4 ip addr add 10.10.0.3/24 dev veth-4
ip netns exec node4 ip link set veth-4 up
ip link set veth-br24 up


ip link add veth-r1 type veth peer name veth-br1 ## router-br1
ip link set veth-r1 netns router
ip link set veth-br1 master br1
ip netns exec router ip addr add 172.0.0.1/24 dev veth-r1
ip netns exec router ip link set veth-r1 up
ip link set veth-br1 up

ip link add veth-r2 type veth peer name veth-br2 ## router-br2
ip link set veth-r2 netns router
ip link set veth-br2 master br2
ip netns exec router ip addr add 10.10.0.1/24 dev veth-r2
ip netns exec router ip link set veth-r2 up
ip link set veth-br2 up


# default gateway
ip netns exec node1 ip route add default via 172.0.0.1
ip netns exec node2 ip route add default via 172.0.0.1
ip netns exec node3 ip route add default via 10.10.0.1
ip netns exec node4 ip route add default via 10.10.0.1
