When in doubt run: erd1keygen --help
Version info: erd1keygen --version

Running the erd1keygen directly, without any params, will search for matching bech32 addresses using the default filter values (check defaults with --help)

Running with custom filters:

A. Run with param and list of custom words: erd1keygen --bech32-filter "aaa|bbb|ccc"

or

B. Run the included find_filter_addresses.bat file on Windows, or find_filter_addresses.sh on Linux (run chmod a+x on it first if needed). 
Edit bat/sh with notepad/nano to customize the words list.

Good luck!

Tip: lower your expectations, start with 3 or 4 letter words or combinations, even so some words will never be matched!

Obs: 

Keep your cool addresses safe, maybe in a password protected zip file, if key generation doesn't change, these could be used on mainnet.

At some point there should be a tool for generating wallets from private keys.

Share you cool address name but not the content of the pem file!

The pem files can be renamed to initialBalancesSk.pem and used with nodes.