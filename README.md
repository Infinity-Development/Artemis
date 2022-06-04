# Artemis
Short Link API for Infinity Bot List

--- 

## Available Endpoints

- Certified Bots can be found via their Vanity
- Uncertified Bots can be found via their Client ID

### Vanity Usage
- With Vanity: `/bot/infinity-bot-list`
- Without Vanity: `/bot/815553000470478850`

| Route                        | Description                    | Example                                   |  
| :---                         | :---                           | :---                                      |
| `/:botID`                    | Short Link to a Bot Page       | `/815553000470478850`                     |
| `/bot/:botID`                | Short Link to a Bot Page       | `/bot/815553000470478850`                 |
| `/bot/:botID/path`           | Debug a Bot Page Short Link    | `/bot/815553000470478850?debug=true`      |
| `/bots/:botID`               | Short Link to a Bot Page       | `/bots/815553000470478850?debug=true`     |
| `/bots/:botID/path`          | Debug a Bot Page Short Link    | `/bots/815553000470478850?debug=true`     |
| `/p/:packID`                 | Short Link to a Bot Pack       | `/p/my-personal-bots`                     |
| `/profile/:userID`           | Short Link to a User Profile   | `/profile/510065483693817867`             |
| `/profiles/:userID`          | Short Link to a User Profile   | `/profiles/510065483693817867`            |
| `/u/:userID`                 | Short Link to a User Profile   | `/u/510065483693817867`                   |
| `/user/:userID`              | Short Link to a User Profile   | `/user/510065483693817867`                |
| `/users/:userID`             | Short Link to a User Profile   | `/users/510065483693817867`               |
| `/:botID/i`                  | Short Link to a Bot Invite     | `/815553000470478850/i`                   |
| `/:botID/inv`                | Short Link to a Bot Invite     | `/815553000470478850/inv`                 |
| `/:botID/invite`             | Short Link to a Bot Invite     | `/815553000470478850/invite`              |
| `/:botID/packs`              | Short Link to a Bot Packs      | `/815553000470478850/packs`               |
| `/:botID/v`                  | Short Link to a Vote Page      | `/815553000470478850/v`                   |
| `/:botID/vote`               | Short Link to a Vote Page      | `/815553000470478850/vote`                |

---

## Contributors
<a href="https://github.com/InfinityBotList/Artemis/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=InfinityBotList/Artemis" />
</a>


