package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/crypto"
	"github.com/ElrondNetwork/elrond-go/crypto/signing"
	"github.com/ElrondNetwork/elrond-go/crypto/signing/mcl"
	"github.com/ElrondNetwork/elrond-go/data/state/factory"
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
		Name:  "filter",
		Usage: "Prefix filter, finds addresses starting with the given values",
		Value: "abc|def",
	}
)

const txSignPubkeyLen = 32
const blsPubkeyLen = 96

func main() {
	app := cli.NewApp()
	cli.AppHelpTemplate = fileGenHelpTemplate
	app.Name = "Validator Key generation Tool (tweaked by @soringp)"
	app.Version = "v0.0.3"
	app.Usage = "This binary will generate validator pem files, containing one private key"
	app.Flags = []cli.Flag{bechFilter}
	app.Authors = []cli.Author{
		{
			Name:  "The Elrond Team + @soringp (find me on Elrond's TG channel)",
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

	genForBlockSigningSk := signing.NewKeyGenerator(mcl.NewSuiteBLS12())

	pubkeyConverter, err := factory.NewPubkeyConverter(
		config.PubkeyConfig{
			Length: blsPubkeyLen,
			Type:   factory.HexFormat,
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Elrond KeyGen tweaked by @soringp")
	fmt.Println("Generating 100 validator keys using filtering... This might take a while...")
	fmt.Println("")

	var count = 1
	var maxcount = 100
	var total = uint64(0)
	for count <= maxcount {

		sk, pk, err := generateKeys(genForBlockSigningSk)
		if err != nil {
			return err
		}

		pkString := pubkeyConverter.Encode(pk)

		var bingo = false

		for i := 0; i < len(prefixes); i++ {
			if strings.HasPrefix(pkString, prefixes[i]) {
				bingo = true
				break
			}
		}

		if bingo {

			var balanceFileFileName = "./" + pkString[0:10] + ".pem"

			balanceFile, err = os.OpenFile(balanceFileFileName, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				return err
			}

			err = core.SaveSkToPemFile(balanceFile, pkString, []byte(hex.EncodeToString(sk)))

			fmt.Println(fmt.Sprintf("Key %d/%d: %s", count, maxcount, pkString[0:10]))

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

func generateKeys(keyGen crypto.KeyGenerator) ([]byte, []byte, error) {
	sk, pk := keyGen.GeneratePair()
	skBytes, err := sk.ToByteArray()
	if err != nil {
		return nil, nil, err
	}

	pkBytes, err := pk.ToByteArray()
	if err != nil {
		return nil, nil, err
	}

	return skBytes, pkBytes, nil
}
