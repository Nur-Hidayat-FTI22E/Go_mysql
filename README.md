# Go_mysql

# Database pooling
(Method)                                            (Information) <br>
(DB) SetMaxIdleConns(number)            -->         setting the minimum number of connections <br>
(DB) SetMaxOpenConns(number)            -->         setting the maximum number of connections <br>
(DB) SetConnMaxIdleTime(duration)       -->         setting how long a connection will be deleted when it is not used <br>
(DB) SetConnMaxLifetime(duration)       -->         setting how long a connection can be used <br>