gc 1        : First GC run since program started.
@0.009s     : Nine milliseconds since the program started.
1%          : One percent of the programs time has been spent in GC.

// wall-clock
0.059ms     : **STW** Sweep termination.
0.17ms      : Mark/Scan - Assist Time (GC performed in line with allocation).
0.005ms     : Mark/Scan - Background GC time.
0.24ms      : Mark/Scan - Idle GC time.
0.12ms      : **STW** Mark termination.

// CPU time
0.17ms      : **STW** Sweep termination.
0.17+0+0ms  : Mark/Scan - Assist Time (GC performed in line with allocation).
0.36ms      : Mark/Scan - Background GC time.
0.067ms     : Mark/Scan - Idle GC time.
0.38ms      : **STW** Mark termination.

5MB         : Heap size at GC start.
5MB         : Heap size at GC end.
3MB         : Live Heap.
4MB         : Goal heap size.
8P          : Number of logical processors.