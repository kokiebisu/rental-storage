package rental_storage.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class HelloResource {
    @GET
    public Response sayHello() {
        String message = "Hello World";
        return Response.ok(message).build();
    }
    
}