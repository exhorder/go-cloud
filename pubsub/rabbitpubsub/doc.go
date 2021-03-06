// Copyright 2018 The Go Cloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rabbitpubsub provides a pubsub driver for RabbitMQ.
//
// RabbitMQ follows the AMQP specification, which uses different terminology
// than Go Cloud Pub/Sub.
//
// A Pub/Sub topic is an AMQP exchange. The exchange kind should be "fanout" to match
// the Pub/Sub model, although publishing will work with any kind of exchange.
//
// A Pub/Sub subscription is an AMQP queue. The queue should be bound to the exchange
// that is the topic of the subscription. See the package example for details.
//
// This package exposes the following types for As:
// Topic: *amqp.Connection
// Subscription: *amqp.Connection
package rabbitpubsub // import "gocloud.dev/pubsub/rabbitpubsub"
