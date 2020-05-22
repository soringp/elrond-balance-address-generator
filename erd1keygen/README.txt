When in doubt run: erd1keygen --help
Version info: erd1keygen --version

Running the erd1keygen directly, without any params, will search for matching bech32 addresses using the default filter values (check defaults with --help)

Running with custom filters:

A. Run with param and list of custom words: erd1keygen --bech32-filter "aaa;bbb;ccc"

';' is the terms separator, used to be '|' but that is not compatible with regex 

or

B. Run the included find_filter_addresses.bat file on Windows, or find_filter_addresses.sh on Linux (run chmod a+x on it first if needed). 
Edit bat/sh with notepad/nano to customize the words list.

Good luck!

Tip: lower your expectations, start with 3 or 4 letter words or combinations, even so some words will never be matched!

Important!!!  Letters 'b','i','o' and number '1' are never present, replace them in your words with '6' for 'b', 'l' for 'i' or '1', and '0' for 'o'

NEW: RegEx search, insted of a word, enter a regex with the "Rx:" prefix. Examples "Rx:^[0-9]{7};Rx:^[ac]{4}" finds addresses starting with 7 numbers, or 4 letters a or c

Obs: 

Keep your cool addresses safe, maybe in a password protected zip file, if key generation doesn't change, these could be used on mainnet.

At some point there should be a tool for generating wallets from private keys.

Share you cool address name but not the content of the pem file!

The pem files can be renamed to initialBalancesSk.pem and used with nodes.