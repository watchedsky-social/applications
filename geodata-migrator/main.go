//go:build migrations && fts5

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/watchedsky-social/libwatchedsky/geodata"
	"github.com/watchedsky-social/libwatchedsky/geodata/migrations"

	_ "github.com/watchedsky-social/go-spatialite"
)

type command struct{}

type cliArgs struct {
	Up       *command `cmd:""`
	UpByOne  *command `cmd:""`
	Down     *command `cmd:""`
	Redo     *command `cmd:""`
	Reset    *command `cmd:""`
	Status   *command `cmd:""`
	Version  *command `cmd:""`
	Fix      *command `cmd:""`
	Validate *command `cmd:""`
	DBFile   string   `short:"f" default:"geodata.db" help:"The DB file name"`
	DataDir  string   `short:"d" type:"existingdir" default:"../../geodata-sources/zones" help:"The directory with source data"`
	S3Bucket string   `default:"watchedsky-social" hidden:""`
	S3Key    string   `default:"geodata.db" hidden:""`
	S3Region string   `default:"us-east-2" hidden:""`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var cli cliArgs
	k := kong.Parse(&cli)

	if cli.S3Bucket != "" && cli.S3Key != "" {
		awsCfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if err = geodata.CopyFromS3(ctx, &geodata.S3Config{
			Region: cli.S3Region,
			Bucket: cli.S3Bucket,
			Key:    cli.S3Key,
		}, awsCfg.Credentials, cli.DBFile); err != nil {
			log.Fatal(err)
		}
	}

	var cmd migrations.GooseCommand
	switch k.Command() {
	case "up":
		cmd = migrations.Up
	case "up-by-one":
		cmd = migrations.UpByOne
	case "down":
		cmd = migrations.Down
	case "redo":
		cmd = migrations.Redo
	case "reset":
		cmd = migrations.Reset
	case "status":
		cmd = migrations.Status
	case "version":
		cmd = migrations.Version
	case "fix":
		cmd = migrations.Fix
	case "validate":
		cmd = migrations.Validate
	default:
		log.Fatalf("Invalid command %s", k.Command())
	}

	if err := cmd(ctx, cli.DBFile, cli.DataDir); err != nil {
		log.Fatal(err)
	}

	if cli.S3Bucket != "" && cli.S3Key != "" {
		awsCfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if err = geodata.SaveToS3(ctx, &geodata.S3Config{
			Region: cli.S3Region,
			Bucket: cli.S3Bucket,
			Key:    cli.S3Key,
		}, awsCfg.Credentials, cli.DBFile); err != nil {
			log.Fatal(err)
		}
	}
}
