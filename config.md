## Configuration
Configuration file is in yaml format. Default path is `config.yaml`  
Example of configuration file [here](example_config.yaml)  
### Configuration fields:
- **port** - Port backend service. **Defaule value**: 8080
- **disable_auth** - Disable jwt auth. **Defaule value**: false
- **mongo_url** - MongoDB url. **Defaule value**: mongodb://mongo:27017
- **access_token_expired_hours** - Access token expired hours. **Defaule value**: 24
- **refresh_token_expired_hours** - Refresh token expired hours. **Defaule value**: 168
- **jwt_secret_key** - Jwt secret key. **Defaule value**: 
