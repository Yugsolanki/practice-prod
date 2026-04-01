# Nginx

## Starting, Stopping and Reloading

```sh
nginx # start
nginx -s stop # fast shutdown
nginx -s quit # graceful shutdown
nginx -s reload # reloading config file
nginx -s reopen # reopening the log files
```

## To create SSL Certificate (Self Signed)

```sh
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt
```
