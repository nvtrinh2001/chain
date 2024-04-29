DIR=`dirname "$0"`

rm -rf ~/.band

# initial new node
bandd init validator --chain-id bandchain
echo "lock nasty suffer dirt dream fine fall deal curtain plate husband sound tower mom crew crawl guard rack snake before fragile course bacon range" \
    | bandd keys add validator --recover --keyring-backend test
echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" \
    | bandd keys add requester1 --recover --keyring-backend test
echo "swamp muffin feature setup pizza relief cruel raven panic bicycle alcohol clip sustain pilot session weapon attend coast slam olympic wild palm casual enough" \
    | bandd keys add requester2 --recover --keyring-backend test

echo "evolve chat where height kitchen grunt erase protect replace miracle steak divert pact decline rose west wolf differ school swamp gate police chicken atom" \
    | bandd keys add requester3 --recover --keyring-backend test

# add accounts to genesis
bandd add-genesis-account validator 10000000000000uband --keyring-backend test
bandd add-genesis-account requester1 10000000000000uband --keyring-backend test
bandd add-genesis-account requester2 10000000000000uband --keyring-backend test
bandd add-genesis-account requester3 10000000000000uband --keyring-backend test


# register initial validators
bandd gentx validator 100000000uband \
    --chain-id bandchain \
    --keyring-backend test

# collect genesis transactions
bandd collect-gentxs


