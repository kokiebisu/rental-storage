import fastify, { FastifyReply, FastifyRequest } from "fastify";
import awsLambdaFastify from "@fastify/aws-lambda";
import { UserServiceImpl } from "../../../application/Service";

exports.handler = async (event: any, context: any) => {
  const app = fastify();

  const userService = await UserServiceImpl.create();
  const proxy = awsLambdaFastify(app);

  app.get(
    "/users/:userId",
    async (
      request: FastifyRequest<{ Params: { userId: string } }>,
      reply: FastifyReply
    ) => {
      const { userId } = request.params;
      const data = await userService.findById(userId);
      reply.send(data);
    }
  );

  app.get(
    "/users/find-by-email",
    async (
      request: FastifyRequest<{ Querystring: { emailAddress: string } }>,
      reply: FastifyReply
    ) => {
      const { emailAddress } = request.query;
      const data = await userService.findByEmail(emailAddress);
      reply.send(data);
    }
  );

  app.post(
    "/users",
    async (
      request: FastifyRequest<{
        Body: {
          emailAddress: string;
          firstName: string;
          lastName: string;
          password: string;
        };
      }>,
      reply: FastifyReply
    ) => {
      const { emailAddress, firstName, lastName, password } = request.body;
      try {
        const data = await userService.createUser({
          emailAddress,
          firstName,
          lastName,
          password,
        });
        reply.send(data);
      } catch (err) {
        reply.statusCode = 500;
        reply.send({ message: "Email already exists" });
      }
    }
  );

  await app.ready();
  return proxy(event, context);
};
