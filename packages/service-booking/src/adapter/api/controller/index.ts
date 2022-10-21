import fastify, { FastifyReply, FastifyRequest } from "fastify";
import awsLambdaFastify from "@fastify/aws-lambda";
import { BookingServiceImpl } from "../../../apps/service";

exports.handler = async (event: any, context: any) => {
  const app = fastify();
  const service = await BookingServiceImpl.create();

  const proxy = awsLambdaFastify(app);

  app.get(
    "/bookings",
    async (request: FastifyRequest<{}>, reply: FastifyReply) => {
      const data = await service.findAllCreatedBookings();
      reply.send({ data });
    }
  );

  await app.ready();
  return proxy(event, context);
};
