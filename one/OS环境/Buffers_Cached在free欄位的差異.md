# What is the major difference between the buffer cache and the page cache? Why were they separate entities in older kernels? Why were they merged later on?
The page cache caches pages of files to optimize file I/O. The buffer cache caches disk blocks to optimize block I/O.

Prior to Linux kernel version 2.4, the two caches were distinct: Files were in the page cache, disk blocks were in the buffer cache. Given that most files are represented by a filesystem on a disk, data was represented twice, once in each of the caches. Many Unix systems follow a similar pattern.

This is simple to implement, but with an obvious inelegance and inefficiency. Starting with Linux kernel version 2.4, the contents of the two caches were unified. The VM subsystem now drives I/O and it does so out of the page cache. If cached data has both a file and a block representation—as most data does—the buffer cache will simply point into the page cache; thus only one instance of the data is cached in memory. The page cache is what you picture when you think of a disk cache: It caches file data from a disk to make subsequent I/O faster.

The buffer cache remains, however, as the kernel still needs to perform block I/O in terms of blocks, not pages. As most blocks represent file data, most of the buffer cache is represented by the page cache. But a small amount of block data isn't file backed—metadata and raw block I/O for example—and thus is solely represented by the buffer cache.


# What is the difference between Buffers and Cached columns in /proc/meminfo output?
https://www.quora.com/What-is-the-difference-between-Buffers-and-Cached-columns-in-proc-meminfo-output
Short answer: Cached is the size of the page cache. Buffers is the size of in-memory block I/O buffers. Cached matters; Buffers is largely irrelevant.

Long answer: Cached is the size of the Linux page cache, minus the memory in the swap cache, which is represented by SwapCached (thus the total page cache size is Cached + SwapCached). Linux performs all file I/O through the page cache. Writes are implemented as simply marking as dirty the corresponding pages in the page cache; the flusher threads then periodically write back to disk any dirty pages. Reads are implemented by returning the data from the page cache; if the data is not yet in the cache, it is first populated. On a modern Linux system, Cached can easily be several gigabytes. It will shrink only in response to memory pressure. The system will purge the page cache along with swapping data out to disk to make available more memory as needed.

`Buffers` are in-memory block I/O buffers. They are relatively short-lived. Prior to Linux kernel version 2.4, Linux had separate page and buffer caches. Since 2.4, the page and buffer cache are unified and Buffers is raw disk blocks not represented in the page cache—i.e., not file data. The Buffers metric is thus of minimal importance. On most systems, Buffers is often only tens of megabytes.


# refer
- https://shuheikagawa.com/blog/2017/05/27/memory-usage/
- https://fabiokung.com/2014/03/13/memory-inside-linux-containers/
- https://docs.docker.com/config/containers/runmetrics/