// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// subscriberCmd represents the subscriber command
var subscriberCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		connStr := viper.GetString("CONNECTION_STRING")
		ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
		if err != nil {
			log.Fatal(err)
		}

		topic, err := ns.NewTopic(context.Background(), viper.GetString("TOPIC_NAME"))
		if err != nil {
			log.Fatal(err)
		}

		subscription, err := topic.NewSubscription(context.Background(), viper.GetString("SUBSCRIPTION_NAME"))
		if err != nil {
			log.Fatal(err)
		}

		listenHandle, err := subscription.Receive(context.Background(),
			func(ctx context.Context, msg *servicebus.Message) servicebus.DispositionAction {
				fmt.Println(string(msg.Data))
				return msg.Complete()
			})

		if err != nil {
			log.Fatal(err)
		}

		// Close the listener handle for the Service Bus Queue
		defer listenHandle.Close(context.Background())

		// Wait for a signal to quit:
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, os.Kill)
		<-signalChan

	},
}

func init() {
	rootCmd.AddCommand(subscriberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscriberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscriberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
