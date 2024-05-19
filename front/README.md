# My finances

1. Build image
    ```bash
    docker build -t app-web .
    ```
2. Run container
    ```bash
   docker run -p 3000:3000 app-web
   ```
3. Go to http://localhost:3000/