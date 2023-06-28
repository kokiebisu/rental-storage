package com.rental-storage;

import io.dropwizard.Application;
import io.dropwizard.Configuration;
import io.dropwizard.setup.Environment;
import java.ws.rs.GET;
import java.ws.rs.Path;
import java.ws.rs.Produces;
import java.ws.core.MediaType;


@Produces(MediaType.APPLICATION_JSON)
public class Server extends Application<Configuration> {
    public static void main(String[] args) throws Exception {
        new Server().run(args);
    }

    @Override
    public void run(Configuration configuration, Environment environment) {
        final Resources resources = new Resources();
        environment.jersey().register(resource);
    }

    public class Resources {
        @GET
        public String Greeting() {
            return "Hello World!";
        }
    }
}