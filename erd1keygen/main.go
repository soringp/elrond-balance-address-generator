package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/crypto"
	"github.com/ElrondNetwork/elrond-go/crypto/signing"
	"github.com/ElrondNetwork/elrond-go/crypto/signing/kyber"
	"github.com/ElrondNetwork/elrond-go/data/state/addressConverters"
	"github.com/urfave/cli"
)

var (
	fileGenHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}
VERSION:
   {{.Version}}
   {{end}}
`
	bechFilter = cli.StringFlag{
		Name:  "bech32-filter",
		Usage: "Bech32 prefix filter, finds addresses starting with erd1 followed by any of the given values",
		Value: "erd2|ddd2|eee2|fff2",
	}
)

func main() {
	app := cli.NewApp()
	cli.AppHelpTemplate = fileGenHelpTemplate
	app.Name = "Key generation Tool (tweaked by @soringp)"
	app.Version = "v0.0.1"
	app.Usage = "This binary will generate balance address pem files, containing one private key"
	app.Flags = []cli.Flag{bechFilter}
	app.Authors = []cli.Author{
		{
			Name:  "The Elrond Team + @soringp (find me on Elrond's Riot channel)",
			Email: "contact@elrond.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		return generateFiles(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func generateFiles(ctx *cli.Context) error {
	var balanceFile *os.File

	defer func() {
	}()

	bechFilter := ctx.GlobalString(bechFilter.Name)

	dirtyTerms := strings.Split(bechFilter, "|")
	prefixes := []string{}

	for i := 0; i < len(dirtyTerms); i++ {

		var term = strings.TrimSpace(dirtyTerms[i])

		if len(term) > 0 {
			prefixes = append(prefixes, term)
		}
	}

	genForBalanceSk := signing.NewKeyGenerator(getSuiteForBalanceSk())

	fmt.Println("Elrond KeyGen tweaked by @soringp")
	fmt.Println("Generating 10 balance addresses using bech32 filtering... This might take a while...")
	fmt.Println("")

	var count = 1
	var maxcount = 10
	var total = uint64(0)
	for count <= maxcount {

		pkHexBalance, skHex, err := getIdentifierAndPrivateKey(genForBalanceSk)
		if err != nil {
			return err
		}

		ac, err := addressConverters.NewPlainAddressConverter(32, "")

		adr, err := ac.CreateAddressFromHex(pkHexBalance)

		if err != nil {
			fmt.Println("For some peculiar reason I could not generate an addressConverter because ", err)
			return nil
		}

		bech32, err := ac.ConvertToBech32(adr)

		var bingo = false

		for i := 0; i < len(prefixes); i++ {
			if strings.HasPrefix(bech32, "erd1"+prefixes[i]) {
				bingo = true
				break
			}
		}

		if bingo {

			var balanceFileFileName = "./" + bech32 + ".pem"

			balanceFile, err = os.OpenFile(balanceFileFileName, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				return err
			}

			err = core.SaveSkToPemFile(balanceFile, pkHexBalance, skHex)

			fmt.Println(fmt.Sprintf("Key %d/%d: %s", count, maxcount, bech32))

			if balanceFile != nil {
				err := balanceFile.Close()
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			count = count + 1
		}

		total = total + 1

		if total >= 1000000 && total%1000000 == 0 {
			fmt.Println(fmt.Sprintf("%dm (%s)", total/1000000, time.Now().Local().Format("01-02 15:04:05")))
		}

	} // end loop

	return nil
}

func getSuiteForBalanceSk() crypto.Suite {
	return kyber.NewBlakeSHA256Ed25519()
}

func getIdentifierAndPrivateKey(keyGen crypto.KeyGenerator) (string, []byte, error) {
	sk, pk := keyGen.GeneratePair()
	skBytes, err := sk.ToByteArray()
	if err != nil {
		return "", nil, err
	}

	pkBytes, err := pk.ToByteArray()
	if err != nil {
		return "", nil, err
	}

	skHex := []byte(hex.EncodeToString(skBytes))
	pkHex := hex.EncodeToString(pkBytes)

	return pkHex, skHex, nil
}
