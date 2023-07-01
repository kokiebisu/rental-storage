package rental_storage;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import rental_storage.resources.HelloResource;

public class SearchApplication extends Application<SearchApplicationConfiguration> {

    public static void main(final String[] args) throws Exception {
        new SearchApplication().run(args);
    }

    @Override
    public String getName() {
        return "service-search";
    }

    @Override
    public void initialize(final Bootstrap<SearchApplicationConfiguration> bootstrap) {
        // TODO: application initialization

    }

    @Override
    public void run(final SearchApplicationConfiguration configuration,
            final Environment environment) {
        // TODO: implement application
        environment.jersey().register(HelloResource.class);
    }

}
