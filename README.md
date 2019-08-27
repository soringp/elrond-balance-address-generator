# Elrond Balance Address Generator

Find cool bech32 Elrond addresses with custom filters.

Uses the untouched elrond open source keygen code, but discards addresses not matching the filter.

Original keygenerator: https://github.com/ElrondNetwork/elrond-go/tree/master/cmd/keygenerator

To compile the keygen, download the https://github.com/ElrondNetwork/elrond-go repository using the tutorial for your platform, described below, and replace cmd/keygenerator/main.go with my version.

**Warning: Do not delete the original main.go, rename it to main.go.org, my version doesn't generate node identities, so you still need the original. I might introduce a separate folder in cmd as a better way to do this, but I'd have to test it first.**

Building on Linux:

Follow the tutorial at https://docs.elrond.com/start-a-validator-node/start-the-network/connecting-from-linux steps 1, 2, 3.

Between 2 and 3 replace main.go as described above.

Compiling on Windows:

Follow the tutorial at https://docs.elrond.com/start-a-validator-node/start-the-network/connecting-from-windows steps 1, 2, 3, 4.

Between 3 and 4 replace main.go as described above.

Read README found next to the binary or next to main.go on this repository for instructions on how to operate the keygenerator.

Have fun!

Contact me on Elrond's Riot channel for any issues or concerns you might have.

Feel free to download and modify the code, but it would be nice to mention me and the Elrond team in the description. Thanks!
