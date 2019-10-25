# hashicorp
Repository to test out hashicorp's offerings as consul, serf, Raft and memberlist

The serf testing follows this [link.](https://jacobmartins.com/2017/01/29/practical-golang-building-a-simple-distributed-one-value-database-with-hashicorp-serf/)

Here we build a simple one value database with serf that is based on SWIM based gossip protocol to propagate the values along the cluster.