If your message size is 600 bytes with 17k msg/s, then your throughput would be ~10MB/s [17000*600/(1024*1024)]. If you're partitioning the topic and using the 5 brokers, with 3 replicas that would be 10/5*3 = 6MB/s per broker you'd need for buffering which shouldn't be a problem on any normal hardware. Buffering 30s would mean 180MB of memory.

In the case that you meant a message size of 600kB, then you'd need to look at adding plenty of very fast storage to reach 6GB/s and actually it would be better to increase the number of nodes of the cluster instead.
