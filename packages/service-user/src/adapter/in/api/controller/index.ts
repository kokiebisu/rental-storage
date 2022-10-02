import fastify, { FastifyReply, FastifyRequest } from "fastify";
import awsLambdaFastify from "@fastify/aws-lambda";
import { UserServiceImpl } from "../../../../application/service/UserService";

exports.handler = async (event: any, context: any) => {
  const app = fastify();
  const service = await UserServiceImpl.create();

  const proxy = awsLambdaFastify(app);

  app.get(
    "/users/guest/:guestId",
    async (
      request: FastifyRequest<{ Params: { guestId: number } }>,
      reply: FastifyReply
    ) => {
      const { guestId } = request.params;
      const data = await service.findGuestById(guestId);
      reply.send(data);
    }
  );

  app.get(
    "/users/host/:hostId",
    async (
      request: FastifyRequest<{ Params: { hostId: number } }>,
      reply: FastifyReply
    ) => {
      const { hostId } = request.params;
      const data = await service.findHostById(hostId);
      reply.send(data);
    }
  );

  await app.ready();
  return proxy(event, context);
};
