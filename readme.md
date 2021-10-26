# spunky
**spunky** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).


### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/singhp1069/spunky@latest! | sudo bash
```
`singhp1069/spunky` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

### Functionality

```
  create-scores Create a new scores (store the highscore)
  update-scores Update a scores (update the score if higher than previous)
  create-nft    Create a new NFT (create a new NFT)
  update-nft    Update a NFT (update the NFT metainfo only by the creator, If the ownership changed then creator can't update NFT)
  transfer-nft  Broadcast message transferNFT (transferring the ownership of the NFT)
  create-rewards Create a new Rewards (lock a balance to module with a milestone and first one to achieve the milestone can claim the reward)
  delete-rewards Delete a Rewards by id (creator can delete the reward only if it was not claimed and the balance will revert back to the creator)
  claim-reward   Broadcast message ClaimReward (User can claim for reward if they are the first one to achieve the milestone)
```


## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
