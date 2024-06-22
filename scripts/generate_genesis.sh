DIR=`dirname "$0"`

rm -rf ~/.band ~/.band1 ~/.band2 ~/.band3

# initial new node
bandd init validator1 --chain-id bandchain --home ~/.band1
bandd init validator2 --chain-id bandchain --home ~/.band2 
bandd init validator3 --chain-id bandchain --home ~/.band3 

echo "lock nasty suffer dirt dream fine fall deal curtain plate husband sound tower mom crew crawl guard rack snake before fragile course bacon range" \
    | bandd keys add validator1 --recover --keyring-backend test --home ~/.band1 
echo "tonight hungry fine vapor original occur loud throw answer region wink alpha cannon dinosaur finger elevator crew degree weasel rack admit property oyster gloom"| bandd keys add validator2 --recover --keyring-backend test --home ~/.band1
echo "hollow seek domain mimic valid attract breeze celery movie blossom practice mention cool left asset sunny seven talent absorb scatter roof speak hard dignity" \ | bandd keys add validator3 --recover --keyring-backend test --home ~/.band1

echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" \
    | bandd keys add requester1 --recover --keyring-backend test --home ~/.band1
echo "swamp muffin feature setup pizza relief cruel raven panic bicycle alcohol clip sustain pilot session weapon attend coast slam olympic wild palm casual enough" \ | bandd keys add requester2 --recover --keyring-backend test --home ~/.band1
echo "evolve chat where height kitchen grunt erase protect replace miracle steak divert pact decline rose west wolf differ school swamp gate police chicken atom" \
    | bandd keys add requester3 --recover --keyring-backend test --home ~/.band1

# add accounts to genesis
bandd add-genesis-account validator1 10000000000000uband --keyring-backend test --home ~/.band1
bandd add-genesis-account validator2 10000000000000uband --keyring-backend test --home ~/.band1
bandd add-genesis-account validator3 10000000000000uband --keyring-backend test --home ~/.band1

bandd add-genesis-account requester1 10000000000000uband --keyring-backend test --home ~/.band1
bandd add-genesis-account requester2 10000000000000uband --keyring-backend test --home ~/.band1
bandd add-genesis-account requester3 10000000000000uband --keyring-backend test --home ~/.band1


# register initial validators
bandd gentx validator1 100000000uband \
    --chain-id bandchain \
    --keyring-backend test --home ~/.band1

# collect genesis transactions
bandd collect-gentxs --home ~/.band1


