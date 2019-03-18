# How to run
  - cd wattx-backend
  - ./build.sh
  - http://127.0.0.1:8888/data will return top 200 assets

# How to test
  - cd wattx-backend/pricing && go test
  - cd wattx-backend/model && go test

# Challenges
  - Please update the Coinmarketcap API key since it has a monthly rate limite.
  - Many Cryptocompare API asset name are not in sync with Coinmarketcap API. Temporary solution is to create a whitelist for top assets.
