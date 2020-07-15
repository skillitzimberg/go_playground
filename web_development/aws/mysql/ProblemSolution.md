## HTTP 503: Service unavailable

The target groups for the load balancer have no registered targets.
I had forgotten to register my running instance as a target.

#### Solution

Make sure you have registered your targets with the load balancer.

## HTTP 502: Bad gateway

### First occurance

Home page ('/') route renders as expected.
Logged in Users page ('/loggedinusers') route renders as expected. (No users are logged in, so no users are displayed.)

'/register', '/login', and '/users' all report "Bad gateway".

## HTTP 502: Bad gateway

### Second occurance

After moving to a more simple app to test connection to the database, I was still getting this error. The reason turned out to be due to my having forgotten to open a connection to the database. This is not the problem in the previous error.

#### Solution

Make sure you are opening (and closing) a database connection properly.
