import fastify, { FastifyReply, FastifyRequest } from "fastify";
import awsLambdaFastify from "@fastify/aws-lambda";
import { ListingServiceImpl } from "../../../app/service";

exports.handler = async (event: any, context: any) => {
  const app = fastify();
  const service = await ListingServiceImpl.create();

  const proxy = awsLambdaFastify(app);

  app.get(
    "/listings/:listingId",
    async (
      request: FastifyRequest<{ Params: { listingId: string } }>,
      reply: FastifyReply
    ) => {
      const { listingId } = request.params;
      const data = await service.findListingById(listingId);
      reply.send(data);
    }
  );

  app.get(
    "/listings",
    async (
      request: FastifyRequest<{
        Querystring: { range: number; latitude: number; longitude: number };
      }>,
      reply: FastifyReply
    ) => {
      const { range, latitude, longitude } = request.query;
      const data = await service.findListingsWithinLatLng(
        latitude,
        longitude,
        range
      );
      reply.send(data);
    }
  );

  await app.ready();
  return proxy(event, context);
};
