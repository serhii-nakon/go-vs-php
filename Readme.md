# Small description

Performance differences between GO and PHP in server side scenario and regular REST API with JSON response

# How to run

Just run `docker compose up`

# How get results

1. GO server side rendering HTML page - `curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/users"`
2. GO JSON REST API - `curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/json_users"`
3. PHP server side rendering HTML page - `curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8081/users"`
4. PHP JSON REST API - `curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8081/json_users"`

# My results

1. GO server side rendering HTML page - `0.750918s`
2. GO JSON REST API - `0.113451s`
3. PHP server side rendering HTML page - `0.195282s`
4. PHP JSON REST API - `0.166908s`
