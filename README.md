# Setup

  

This repository contains the necessary files to set up and run Automatic Guacamole using Docker.

  

## Prerequisites

  

- Docker installed on your machine.

- The `db.sql` file for initializing the database.

  

## Instructions



1. Execute Database Setup

2. Pull Container and Run
`docker run -d \
--name guacamole-container \
-p 8000:8000 \
-e APP_PORT="" \
-e APP_SECRET="" \
-e DB_HOST="" \
-e DB_USER="" \
-e DB_PASS="" \
-e DB_NAME="" \
-e DB_PORT="" \
-e DB_SSL_MODE="" \
-e DB_TIMEZONE="" \
sandipradana94/automatic-guacamole:1.0.0`