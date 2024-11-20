# What is this?
With 2 days left for my assignment, I had the bright idea to rewrite my previous backend project in Go and submit it as the assignment. 

For the rewrite, instead of calling out to Redis and MySQL I am storing data on memory. I am using a (unbalanced) binary search tree to replace user MySQL table and 2 priority queues to replace the Redis and hall of fame MySQL table.

Find the original repository here: [Mind-Space](https://github.com/icecreampoop/Mind-Space)

# Where it can be improved
1. balance the binary search tree

2. binary heap for the priority queue

3. the daily priority queue should "expire" upon 0000 Singapore time

# How to use/play it
Prerequisites:
- GoLang installed
- Node.js installed (v18 works)
- Angular CLI installed (v17 works)
<br>
<br>

1) Clone the repository

2) Run commands in terminal (runs the project in dev mode)<br>
\> npm i (run this in the frontend folder)<br>
\> ng serve --proxy-config proxy.config.json (run this in the frontend folder)<br>
\> go run main.go apicontroller.go fakedb.go (run this in the backend folder)<br>

3) Open your browser and navigate to http://localhost:4200/
