package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"

	"googleiapclient/pkg/client"
)

const CREDS_ENV_VAR_NAME = "GOOGLE_CREDS"

type cliOpts struct {
	CredsPath           string `long:"google-creds-file" value-name:"PATH" description:"Path to Google Cloud credentials file" required:"false"`
	Creds               string `long:"google-credentials" value-name:"BASE64_STRING" description:"Base64 encoded google credentials json" required:"false"`
	OAuthClientID       string `long:"oauth-client-id" value-name:"CLIENT_ID" description:"OAuth client ID provided by Google. For example \"823926513327-pr0714rqtdb223bahl0nq2jcd4ur79ec.apps.googleusercontent.com\"" required:"true"`
	RequestedExpiration string `long:"requested-expiration" value-name:"DURATION" description:"Time from now that you would like for IAP token to expire (e.g. 1h). Not guaranteed to be honored by Google OAuth service." required:"true"`
}

func parseCliOptions() (*cliOpts, []string, error) {
	opts := &cliOpts{}
	args, err := flags.ParseArgs(opts, os.Args)
	if opts.CredsPath != "" && opts.Creds == "" {
		return nil, nil, fmt.Errorf("Either google-creds-file or google-credentials has to be set")
	}
	return opts, args, err
}

func credsFromFileInBase64(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if len(b) == 0 {
		log.Panicf("Credentials file (\"%s\") must not be empty", path)
	}

	b64str := base64.StdEncoding.EncodeToString(b)

	return b64str
}

func main() {
	opts, _, err := parseCliOptions()
	if err != nil {
		return
	}

	// parse expiration
	requestedExpiration, err := time.ParseDuration(opts.RequestedExpiration)
	if err != nil {
		panic(err)
	}

	// take creds from file, base64 encode, and place in env (lib requires this)
	var credsB64 string
	if opts.CredsPath != "" {
		credsB64 = credsFromFileInBase64(opts.CredsPath)
	} else {
		credsB64 = opts.Creds
	}

	err = os.Setenv(CREDS_ENV_VAR_NAME, credsB64)
	if err != nil {
		panic(err)
	}

	iapClient := googleiapclient.NewIAPClient(CREDS_ENV_VAR_NAME, requestedExpiration)
	token, err := iapClient.JWTToken(opts.OAuthClientID)
	if err != nil {
		log.Panicf("Could not get JWT token: %+v", err)
	}

	fmt.Printf("Authorization: Bearer %s\n", token)
}
