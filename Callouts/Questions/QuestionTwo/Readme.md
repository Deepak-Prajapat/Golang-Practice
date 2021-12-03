In the 2nd question when defining gorouting they were not giving perameter to it, because goroutine runs concurrently 
and loop executes so fast so It was holding the last value of slice only for all threads that is 80 and returning 160 for all process

What I did
I just give current value of slice while iteration over it to goroutine so each time goroutine have seperate value and the issue is fixed
