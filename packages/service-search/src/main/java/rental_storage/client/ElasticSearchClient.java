package rental_storage.client;

import org.apache.http.HttpHost;
// import org.elasticsearch.client.ElasticsearchClient;
import org.elasticsearch.client.RestClient;
import org.elasticsearch.client.RestClientBuilder;

import co.elastic.clients.elasticsearch.ElasticsearchClient;
import co.elastic.clients.json.jackson.JacksonJsonpMapper;
import co.elastic.clients.transport.ElasticsearchTransport;
import co.elastic.clients.transport.rest_client.RestClientTransport;

public class ElasticSearchClient {
    private ElasticsearchClient esClient;

    public ElasticSearchClient() {
        HttpHost httpHost = new HttpHost("example.com", 80, "http");

        RestClientBuilder builder = RestClient.builder(httpHost);
        RestClient restClient = builder.build();
        // Create the transport with a Jackson mapper
        ElasticsearchTransport transport = new RestClientTransport(
                restClient, new JacksonJsonpMapper());

        // And create the API client
        this.esClient = new ElasticsearchClient(transport);
    }
}
