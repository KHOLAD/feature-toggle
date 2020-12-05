# Feature toggle API
This project uses MongoDB as database, GoLang with echo library (https://echo.labstack.com)
and docker as an environment.

### Project setup
- For this project docker needs to be installed (https://docs.docker.com/docker-for-windows/install/)
- Git clone project
- In root folder of project run this command (will setup API with port:8080)
    ```
    docker-compose up --build
    ```
- Database user (admin/password) if needed
### API endpoints

| API |  |
| ------ | ------ |
| GET | api/features |
| POST | api/feature |
| PUT | api/feature/:id |
| GET | api/customers |
| GET | api/customer/:id |
| PUT | api/toggle/:customerId/:featureName |
