package rental_storage;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;

public class SearchServiceApplication extends Application<SearchServiceConfiguration> {

    public static void main(final String[] args) throws Exception {
        new SearchServiceApplication().run(args);
    }

    @Override
    public String getName() {
        return "service-search";
    }

    @Override
    public void initialize(final Bootstrap<SearchServiceConfiguration> bootstrap) {
        // TODO: application initialization

    }

    @Override
    public void run(final SearchServiceConfiguration configuration,
            final Environment environment) {
        // TODO: implement application
        environment.jersey().register(Hello.class);
    }

}
