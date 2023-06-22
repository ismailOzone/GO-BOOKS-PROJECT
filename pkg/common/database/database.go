package database

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/config"
)

type elasticDB struct {
	Client *elasticsearch.Client
}

var eDB *elasticDB

type ElasticDB interface {
	Index(ctx *context.Context, buf bytes.Buffer, indexname string, id string) error
	Delete(ctx *context.Context, id string, indexname string) error
	Search(ctx *context.Context, buf bytes.Buffer, indexname string) (io.ReadCloser, error)
}

func NewElasticDB(Client *elasticsearch.Client) ElasticDB{
	return &elasticDB{
		Client: Client,
	}
}

func Get() ElasticDB{
	var err error
	if eDB==nil{
		var esClient *elasticsearch.Client
		esClient, err=ConnectElasticsearch()
		if err!=nil{
			log.Fatal("failed to connect to ES ", err)
		}
		eDB=&elasticDB{
			Client: esClient,
		}
	}
	return eDB
}

// ConnectElasticsearch creates and returns a new Elasticsearch client
func ConnectElasticsearch() (client *elasticsearch.Client, err error) {
	log.Println("Connecting to ES")
	defer log.Println("Done COnnecting to ES")
	cfg:=config.Get()

	// Giving the Elasticsearch client configuration
	escfg := elasticsearch.Config{
		// Addresses: []string{ElasticsearchURL}, // Specify the Elasticsearch server address(es)

		Addresses: []string{
			cfg.ElasticURL,
		},
		Username: cfg.ElasticUsername,
		Password: cfg.ElasticPassword,
	}
	fmt.Println("========================================",escfg)
	client, err = elasticsearch.NewClient(escfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	// Ping the Elasticsearch cluster to verify the connection
	_ , err = client.Ping(client.Ping.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch cluster: %s", err)
	}

	// defer func() {
	// 	err = res.Body.Close()
	// 	if err != nil {
	// 		log.Fatalf("Error closing response body: %s", err)
	// 	}
	// }()

	fmt.Println("Connected to Elasticsearch!")
	return client, nil
}

func (t *elasticDB) Index(ctx *context.Context, buf bytes.Buffer, indexname string, id string) error {
	res, err := t.Client.Index(indexname, &buf, t.Client.Index.WithDocumentID(id))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to index document: %s", res.Status())
	}

	io.Copy(io.Discard, res.Body)

	return nil
}

func (t *elasticDB) Delete(ctx *context.Context, id string, indexname string) error {
	res, err := t.Client.Delete(indexname, id)
	if err != nil {
		return err
	}
	// defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to delete document: %s", res.Status())
	}

	io.Copy(io.Discard, res.Body)

	return nil
}

func (t *elasticDB) Search(ctx *context.Context, buf bytes.Buffer, indexname string) (io.ReadCloser, error) {
	res, err := t.Client.Search(
		t.Client.Search.WithIndex(indexname),
		t.Client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("search error: %s", res.Status())
	}

	return res.Body, nil
}