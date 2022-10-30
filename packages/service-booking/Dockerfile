FROM amazon/aws-lambda-nodejs:14

COPY . ${LAMBDA_TASK_ROOT}/

RUN npm install -g pnpm
RUN pnpm install

CMD [ "app.handler" ]